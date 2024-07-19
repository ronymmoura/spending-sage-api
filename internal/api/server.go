package api

import (
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

type Server struct {
	router *gin.Engine
	config util.Config
}

func NewServer(config util.Config) (server *Server, err error) {
	clerk.SetKey(config.ClerkKey)

	server = &Server{
		config: config,
	}

	server.setupRouter()

	return
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	if server.config.Environment == "development" {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.POST("auth/signIn", server.signIn)

	protectedRoutes := router.Group("/").Use(authMiddleware())

	protectedRoutes.GET("/user", server.getUser)

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
