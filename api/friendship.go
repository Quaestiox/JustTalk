package api

import (
	"database/sql"
	"errors"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateFriendshipRequest struct {
	ToID int64 `json:"to_id" binding:"required,min=1"`
}

func (server *Server) createFriendship(ctx *gin.Context) {
	var req CreateFriendshipRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)

	arg := db.CreateFriendShipParams{
		FromID: authPayload.UserID,
		ToID:   req.ToID,
	}

	friendship, err := server.store.CreateFriendShip(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, friendship)
}

type getFriendshipRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getFriendship(ctx *gin.Context) {
	var req getFriendshipRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	friendship, err := server.store.GetFriendShip(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)
	if friendship.FromID != authPayload.UserID && friendship.ToID != authPayload.UserID {
		err := errors.New("friendship doesn't relate to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, friendship)
}

type listFriendshipRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFriendship(ctx *gin.Context) {
	var req listFriendshipRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)

	arg := db.ListFriendShipParams{
		FromID: authPayload.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListFriendShip(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
