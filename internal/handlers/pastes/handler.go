package pastes

import (
	"context"
	"cpypst/internal/database"
	"cpypst/internal/models/domain"
	"cpypst/internal/models/generated"
	"cpypst/pkg/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PasteHandler interface {
	// GetPasteByID retrieves a paste by its ID.
	GetPastesByUser(w http.ResponseWriter, r *http.Request)

	// CreatePaste creates a new paste.
	CreatePaste(w http.ResponseWriter, r *http.Request)

	// DeletePaste deletes a paste by its ID.
	DeletePaste(w http.ResponseWriter, r *http.Request)

	// GetPasteBySlug retrieves a paste by its slug.
	GetPasteBySlug(w http.ResponseWriter, r *http.Request)
}

type PasteHandlerImpl struct{}

func (p PasteHandlerImpl) GetPastesByUser(w http.ResponseWriter, r *http.Request) {
	// implementation

	vars := mux.Vars(r)

	param := vars["userid"]

	id, err := strconv.Atoi(param)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid userID")
	}

	if id <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid userID")
	}

	query := generated.New(database.New())

	pastes, err := query.GetPastesByUser(context.Background(), sql.NullInt32{
		Int32: int32(id), Valid: true})

	output := make([]domain.Paste, 0, len(pastes))

	for key := range pastes {

		output = append(output, domain.ToDomainPaste(pastes[key]))
	}

	utils.RespondWithJSON(w, http.StatusOK, output)

}

func (p PasteHandlerImpl) GetPasteBySlug(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	slug := vars["slug"]

	if slug == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid slug")
	}

	query := generated.New(database.New())

	paste, err := query.GetPasteBySlug(context.Background(), slug)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Paste not found")
	}

	//response with domain paste
	output := domain.Paste{
		Id:         int(paste.ID),
		Title:      paste.Title.String,
		Content:    paste.Content,
		UserId:     int(paste.UserID.Int32),
		Slug:       paste.Slug,
		Syntax:     paste.Syntax.String,
		IsEditable: paste.Editable.Bool,
		ExpiresAt:  paste.ExpirationTime.Time,
		CreatedAt:  paste.CreatedAt.Time,
	}

	utils.RespondWithJSON(w, http.StatusOK, output)
}

func (p PasteHandlerImpl) CreatePaste(w http.ResponseWriter, r *http.Request) {

	//get user id
	// id := r.Context().Value("user_id")
	id := 1

	var paste domain.Paste
	//parse the request body
	if err := json.NewDecoder(r.Body).Decode(&paste); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paste.Slug = "" + utils.GenerateSlug(int64(id))

	//validate the paste
	err := paste.Validate()

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error: "+err.Error())
		return
	}

	//save the paste to the database

	// paste.UserId = int(id)

	dbs := database.New()

	query := generated.New(dbs)

	queryParams := generated.CreatePastesParams{
		Title: sql.NullString{
			String: paste.Title,
			Valid:  true,
		},
		Content: paste.Content,
		Syntax: sql.NullString{
			String: paste.Syntax,
			Valid:  true,
		},
		Editable: sql.NullBool{
			Bool:  paste.IsEditable,
			Valid: true,
		},
		ExpirationTime: sql.NullTime{
			Time:  paste.ExpiresAt,
			Valid: true,
		},
		UserID: sql.NullInt32{
			Int32: int32(id),
			Valid: true,
		},
		Slug: paste.Slug,
	}
	err = query.CreatePastes(r.Context(), queryParams)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating paste: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, paste)

}
