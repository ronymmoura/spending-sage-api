package cache

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func (cache RedisCache) SetCategories(ctx *gin.Context, categories []db.Category) error {
	for _, category := range categories {
		json, err := json.Marshal(category)
		if err != nil {
			return err
		}

		err = cache.client.Set(ctx, "category:"+strconv.FormatInt(category.ID, 10), string(json), 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (cache RedisCache) GetCategory(ctx *gin.Context, categoryID int64) (category db.Category, err error) {
	cat := cache.client.Get(ctx, "category:"+strconv.FormatInt(categoryID, 10)).Val()
	err = json.Unmarshal([]byte(cat), &category)

	return
}
