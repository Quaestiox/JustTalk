package api

import (
	"database/sql"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type createUserRequest struct {
	Name      string  `json:"name" binding:"required,alphanum"`
	Password  string  `json:"password" binding:"required,min=6"`
	Nickname  string  `json:"nickname" binding:"required,alphanum"`
	AvatarURL string  `json:"avatar_url" binding:"required"`
	Friends   []int64 `json:"friends" binding:"required"`
}

type userResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Nickname    string    `json:"nickname"`
	AvatarURL   string    `json:"avatarURL"`
	FriendCount int32     `json:"friendCount"`
	Friends     []int64   `json:"friends"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:          user.ID,
		Name:        user.Name,
		Nickname:    user.Nickname,
		AvatarURL:   user.AvatarUrl,
		FriendCount: user.FriendCount,
		Friends:     user.Friends,
		CreateAt:    user.CreateAt,
		UpdateAt:    user.UpdateAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPd, _ := util.HashPassword(req.Password)

	arg := db.CreateUserParams{
		Name:      req.Name,
		Password:  hashedPd,
		Nickname:  req.Nickname,
		AvatarUrl: req.AvatarURL,
		Friends:   req.Friends,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rp := newUserResponse(user)

	ctx.JSON(http.StatusOK, rp)
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUserParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type loginUserRequest struct {
	Name     string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByName(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		user.ID,
		user.Name,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rp)
}
