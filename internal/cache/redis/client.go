package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
	"github.com/ronymmoura/spending-sage-api/internal/usecases"
)

type RedisCache struct {
	client *redis.Client
}

func NewCache(url string, password string, db int) RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       db,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return RedisCache{
		client: client,
	}
}

func (cache RedisCache) FillCache(ctx *gin.Context, store db.Store) error {
	categories, err := usecases.ListCategoriesUseCase(ctx, store)
	if err != nil {
		return err
	}

	if err := cache.SetCategories(ctx, categories); err != nil {
		return err
	}

	origins, err := usecases.ListOriginsUseCase(ctx, store)
	if err != nil {
		return err
	}

	if err := cache.SetOrigins(ctx, origins); err != nil {
		return err
	}

	return nil
}
