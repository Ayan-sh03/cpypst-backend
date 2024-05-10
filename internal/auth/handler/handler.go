package handler

import (
	"context"
	"cpypst/internal/auth"
	"cpypst/internal/database"
	"cpypst/internal/models/domain"
	"cpypst/internal/models/generated"
	"cpypst/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Auth interface {
	// func to register the user
	RegisterUserHandler(w http.ResponseWriter, r *http.Request)

	//func to Login the user and return token in josn
	LoginUserHandler(w http.ResponseWriter, r *http.Request)
}

type AuthImpl struct{}

func (a AuthImpl) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	// user.Password = hashedPassword

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	query := generated.New(database.New())

	genUserParams := generated.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err = query.CreateUser(context.Background(), genUserParams)

	if err != nil {
		log.Println("Error occured creating user", err)
		if strings.Contains(err.Error(), "\"users_email_key\"") {
			utils.RespondWithError(w, http.StatusConflict, "User already exists")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
	}
	utils.RespondWithJSON(w, http.StatusCreated, response)

}

func (a AuthImpl) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	query := generated.New(database.New())

	dbUser, userErr := query.GetUserByUsername(context.Background(), user.Username)

	if userErr != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, userErr.Error())
		return
	}

	// check password
	securityErr := auth.CheckPassword(user.Password, dbUser.Password)
	if securityErr != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid Credentials : "+securityErr.Error())
		return
	}

	token, err := auth.GenerateJWT(dbUser.Username, int64(dbUser.ID)) 

	if !err {
		utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error : Error While generating Token")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": "Bearer " + token}) // Only token
}
