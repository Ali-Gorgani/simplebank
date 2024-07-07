package api

import (
	"log"

	db "github.com/Ali-Gorgani/simplebank/db/sqlc"
	"github.com/Ali-Gorgani/simplebank/util"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/Ali-Gorgani/simplebank/token"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:username", server.getUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.PATCH("/accounts", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse returns error details in JSON format.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}