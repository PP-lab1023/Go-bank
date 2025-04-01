package api

import (
	db "github.com/PP-lab1023/Go-bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// Add routes to router
	router.POST("/accounts", server.createAccount)
	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address to start listning for API requests.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
