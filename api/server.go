package api

import (
	"fmt"
	db "github.com/Quaestiox/JustTalk_backend/db/sqlc"
	"github.com/Quaestiox/JustTalk_backend/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *db.Store
	router     *gin.Engine
	tokenMaker util.Maker
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := util.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.POST("/user", server.createUser)
	router.GET("/user/:id", server.getUser)
	router.GET("/user", server.listUser)
	router.POST("/user/login", server.loginUser)

	authRoutes.POST("/friendship", server.createFriendship)
	authRoutes.GET("/friendship/:id", server.getFriendship)
	authRoutes.GET("/friendship", server.listFriendship)

	authRoutes.POST("/message", server.CreateMessage)
	authRoutes.GET("/message/:id", server.getMessage)
	authRoutes.GET("/message", server.listMessage)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
