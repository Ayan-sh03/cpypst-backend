package server

import (
	"cpypst/internal/auth/handler"
	"cpypst/internal/handlers/pastes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.HelloWorldHandler)

	//--- Pastes ---//
	pasteHandler := pastes.PasteHandlerImpl{}
	r.HandleFunc("/api/v1/paste", func(w http.ResponseWriter, r *http.Request) {
		pasteHandler.CreatePaste(w, r)
	}).Methods("POST")
	r.HandleFunc("/api/v1/paste/{slug:[A-Z]+}", func(w http.ResponseWriter, r *http.Request) {
		pasteHandler.GetPasteBySlug(w, r)
	}).Methods("GET")
	r.HandleFunc("/api/v1/paste/{userid:[0-9]+}/pastes", func(w http.ResponseWriter, r *http.Request) {
		pasteHandler.GetPastesByUser(w, r)
	})


	//-- User --//
	userHandler := handler.AuthImpl{}

	r.HandleFunc("/api/v1/user", func(w http.ResponseWriter, r *http.Request) {
		userHandler.RegisterUserHandler(w, r)
	}).Methods("POST")
	r.HandleFunc("/api/v1/user/login", func(w http.ResponseWriter, r *http.Request) {
		userHandler.LoginUserHandler(w, r)
	}).Methods("POST")
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
