package store

import (
	"context"
	"encoding/binary"
	"env_middleware/data"
	"env_middleware/util"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redis *redis.Client) *RedisClient {
	redisClient := RedisClient{
		client: redis,
	}
	return &redisClient
}

func (redisClient *RedisClient) ParseCesHeightmap(ctx context.Context, content, tileID string) error {
	level, X, Y, err := util.ParseTileID(tileID)
	if err != nil {
		return err
	}
	minLon, minLat := util.GetLonLat(level, X, Y)
	maxLon, maxLat := util.GetLonLat(level, X+1, Y+1)
	terrainData := []byte(content)
	offset := 0

	for row := 0; row < 65; row++ {
		for col := 0; col < 65; col++ {
			tempDataTile := binary.LittleEndian.Uint16(terrainData[offset:])
			offset += 2
			tempDataTile = uint16(0.2*float64(tempDataTile) - 1000)

			minLatCell := float64(65-row-1)/65.0*(maxLat-minLat) + minLat
			minLonCell := float64(col)/65.0*(maxLon-minLon) + minLon
			//maxLatCell := float64(65-row)/65.0*(maxLat-minLat) + minLat
			//maxLonCell := float64(col+1)/65.0*(maxLon-minLon) + minLon

			key := fmt.Sprintf("%d-%f-%f", level, minLonCell, minLatCell)
			err := redisClient.client.HSet(ctx, key, "TERRAIN", strconv.FormatFloat(float64(tempDataTile), 'f', -1, 64)).Err()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (redisClient *RedisClient) InsertCrater(ctx context.Context, crater data.Crater) error {
	craterId := "crater-" + strconv.FormatInt(crater.CraterID, 10)
	err := redisClient.client.GeoAdd(ctx, "geo_locations", &redis.GeoLocation{
		Longitude: crater.Position.Longitude,
		Latitude:  crater.Position.Latitude,
		Name:      craterId,
	}).Err()
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"width":     crater.Width,
		"depth":     crater.Depth,
		"longitude": crater.Position.Longitude,
		"latitude":  crater.Position.Latitude,
	}
	err = redisClient.client.HSet(ctx, craterId, params).Err()
	return err
}
