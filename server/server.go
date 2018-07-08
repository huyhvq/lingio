package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/huyhvq/lingio/db"
	"github.com/huyhvq/lingio/server/handler"
)

type Server struct {
	DB db.DB
}

func NewServer() (Server) {
	d := db.NewDB()
	return Server{DB: d}
}

func (s *Server) Start() {
	r := mux.NewRouter()
	h := handler.NewHandler(s.DB)
	r.HandleFunc("/users/{user_id}/games/{game_id}", h.UserGame).Methods("GET")
	r.HandleFunc("/users/{user_id}/scores", h.UserScores).Methods("GET")
	http.ListenAndServe(":8099", r)
}
