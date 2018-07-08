package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/huyhvq/lingio/db"
	"github.com/huyhvq/lingio/util"
)

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) (Handler) {
	return Handler{db: db}
}

func (h *Handler) UserGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	gameId := vars["game_id"]

	h.db.Set(db.Field{UserId: userId, GameId: gameId, Level: "123", Score: 100})

	f, b := h.db.GetByUserIdAndGameId(userId, gameId)
	if b {
		util.ResponseOk(w, f)
		return
	}
	util.ResponseError(w, http.StatusConflict, "can't update your scored")
	return
}

func (h *Handler) UserScores(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["user_id"]
	scores := h.db.GetUserScoresByUserId(userId)
	util.ResponseOk(w, struct {
		UserId      string
		TotalScores int
	}{UserId: userId, TotalScores: scores})
	return
}
