package cache

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/ronymmoura/spending-sage-api/internal/db/sqlc"
)

func (cache RedisCache) SetOrigins(ctx *gin.Context, origins []db.Origin) error {
	for _, origin := range origins {
		json, err := json.Marshal(origin)
		if err != nil {
			return err
		}

		err = cache.client.Set(ctx, "origin:"+strconv.FormatInt(origin.ID, 10), string(json), 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (cache RedisCache) GetOrigin(ctx *gin.Context, originID int64) (origin db.Origin, err error) {
	cat := cache.client.Get(ctx, "origin:"+strconv.FormatInt(originID, 10)).Val()
	err = json.Unmarshal([]byte(cat), &origin)

	return
}
