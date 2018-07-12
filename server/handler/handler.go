package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/huyhvq/lingio/db"
	"github.com/huyhvq/lingio/util"
	"encoding/json"
)

type Handler struct {
	db db.DB
}

type UserGameRequestBody struct {
	Level string `json:"level"`
	Score int    `json:"score"`
}

func NewHandler(db db.DB) (Handler) {
	return Handler{db: db}
}

func (h *Handler) UserGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	gameId := vars["game_id"]
	var requestBody UserGameRequestBody
	json.NewDecoder(r.Body).Decode(&requestBody)

	h.db.Set(db.Field{UserId: userId, GameId: gameId, Level: requestBody.Level, Score: requestBody.Score})

	f, _ := h.db.GetByUserIdAndGameId(userId, gameId)
	util.ResponseOk(w, f)
	return
}

func (h *Handler) UserScores(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	scores := h.db.GetUserScoresByUserId(userId)
	util.ResponseOk(w, struct {
		UserId      string `json:"user_id"`
		TotalScores int    `json:"total_scores"`
	}{UserId: userId, TotalScores: scores})
	return
}

func (h *Handler) UserMedals(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	scores := h.db.GetUserMedalByUserId(userId)
	util.ResponseOk(w, struct {
		UserId     string `json:"user_id"`
		TotalMedal int    `json:"total_medal"`
	}{UserId: userId, TotalMedal: scores})
	return
}
