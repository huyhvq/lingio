package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/huyhvq/lingio/db"
	"github.com/huyhvq/lingio/server/handler"
	"fmt"
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
	r.HandleFunc("/", h.Index).Methods("GET")
	r.HandleFunc("/users/{user_id}/games/{game_id}", h.UserGame).Methods("POST")
	r.HandleFunc("/users/{user_id}/scores", h.UserScores).Methods("GET")
	r.HandleFunc("/users/{user_id}/medal", h.UserMedals).Methods("GET")
	var dir string
	r.PathPrefix("/docs/").Handler(http.FileServer(http.Dir(dir)))

	fmt.Println("Server started at :8099")
	http.ListenAndServe(":8099", r)
}
