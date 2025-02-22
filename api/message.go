package api

import (
	"database/sql"
	"errors"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createMessageRequest struct {
	ReceiverID int64  `json:"receiver_id" binding:"required,min=1"`
	Content    string `json:"content" binding:"required"`
}

func (server *Server) CreateMessage(ctx *gin.Context) {
	var req createMessageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)

	arg := db.CreateMessageParams{
		SenderID:   authPayload.UserID,
		ReceiverID: req.ReceiverID,
		Content:    req.Content,
	}

	message, err := server.store.CreateMessage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, message)
}

type getMessageRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getMessage(ctx *gin.Context) {
	var req getMessageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	message, err := server.store.GetMessage(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)
	if message.SenderID != authPayload.UserID && message.ReceiverID != authPayload.UserID {
		err := errors.New("message doesn't relate to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, message)
}

type listMessageRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMessage(ctx *gin.Context) {
	var req listMessageRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*util.Payload)

	arg := db.ListMessageParams{
		SenderID: authPayload.UserID,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	message, err := server.store.ListMessage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, message)
}
