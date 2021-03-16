package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/constants"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/db"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/middleware"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/token"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/util"
)

// Server will server HTTP requests
type Server struct {
	store   db.Store
	router  *gin.Engine
	tokener token.Tokener
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) (*Server, error) {
	tokener, err := token.NewJWTToken("12345678901234567890123456789012")
	if err != nil {
		return nil, fmt.Errorf("failed to create tokener: %w", err)
	}

	router := gin.New()

	if util.GetEnv() == constants.EnvRelease {
		log.Println("RELEASE MODE ENABLED")
		gin.SetMode(gin.ReleaseMode)
	}

	server := &Server{
		store:   store,
		tokener: tokener,
		router:  router,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupAuthRouter() {
	r := server.router

	authGroup := r.Group("/")

	authGroup.Use(middleware.Auth(server.tokener))
	authGroup.GET("/users", server.getUsers)
	authGroup.POST("/users", server.createUser)
	authGroup.PUT("/users", server.updateUser)

}

func (server *Server) setupRouter() {
	r := server.router

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", server.loginUser)

	server.setupAuthRouter()
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
