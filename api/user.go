package api

import (
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserRequest struct {
	Name      string  `json:"name" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Nickname  string  `json:"nickname" binding:"required"`
	AvatarURL string  `json:"avatar_url" binding:"required"`
	Friends   []int64 `json:"friends" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:      util.RandomAccount(),
		Password:  util.RandomPassword(),
		Nickname:  util.RandomAccount(),
		AvatarURL: util.RandomPassword(),
		Friends:   []int64{},
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
