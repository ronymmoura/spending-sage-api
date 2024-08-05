package api

import (
	"context"
	"log"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	cache "github.com/ronymmoura/spending-sage-api/internal/cache/redis"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

type Server struct {
	Router *gin.Engine
	Config util.Config
	Store  db.Store
	Cache  cache.RedisCache
}

func NewServer(config util.Config) (server *Server, err error) {
	clerk.SetKey(config.ClerkKey)

	server = &Server{
		Config: config,
	}

	server.setupDB()
	server.setupCache()
	server.setupRouter()

	return
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func (server *Server) setupDB() {
	connPool, err := pgxpool.New(context.Background(), server.Config.DatabaseUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server.Store = store
}

func (server *Server) setupCache() {
	cache := cache.NewCache(server.Config.CacheUrl, server.Config.CachePassword, server.Config.CacheDatabase)
	server.Cache = cache
}

// func cacheMiddleware(server *Server) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		server.Cache.FillCache(ctx, server.Store)
// 	}
// }

func (server *Server) setupRouter() {
	if server.Config.Environment == "development" {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//router.Use(cacheMiddleware(server))

	router.POST("auth/signIn", server.SignInRoute)

	protectedRoutes := router.Group("/").Use(authMiddleware())

	protectedRoutes.GET("/user", server.GetUserRoute)

	protectedRoutes.GET("/lists", server.GetListsRoute)

	protectedRoutes.GET("/months", server.ListMonthsRoute)
	protectedRoutes.POST("/months", server.CreateMonthRoute)

	protectedRoutes.GET("/months/:month_id/entries", server.SearchMonthEntriesRoute)
	protectedRoutes.POST("/months/:month_id/entries", server.CreateMonthEntryRoute)

	protectedRoutes.GET("/fixedEntries", server.SearchFixedEntriesRoute)
	protectedRoutes.POST("/fixedEntries", server.CreateFixedEntryRoute)

	server.Router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
