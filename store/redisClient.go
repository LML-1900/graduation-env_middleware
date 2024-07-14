package store

import (
	"context"
	"encoding/binary"
	"env_middleware/data"
	"env_middleware/util"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math"
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

func (redisClient *RedisClient) GetAltitude(ctx context.Context, position data.LonLatPosition, level int) (float64, error) {
	staticAltitude, err := redisClient.GetStaticAltitude(ctx, position, level)
	if err != nil {
		return 0, err
	}
	dynamicAltitude, err := redisClient.GetDynamicAltitude(ctx, position)
	if err != nil {
		return 0, err
	}
	return staticAltitude - dynamicAltitude, nil
}

func (redisClient *RedisClient) GetStaticAltitude(ctx context.Context, position data.LonLatPosition, level int) (float64, error) {
	neighbors := make([]float64, 0)
	GetNeignbor(position, level, &neighbors)
	minlon, maxlon := neighbors[0], neighbors[1]
	minlat, maxlat := neighbors[2], neighbors[3]

	h1, err := redisClient.GetAltitudeValFromRedis(ctx, minlon, maxlat, level)
	if err != nil {
		return 0, err
	}
	h2, err := redisClient.GetAltitudeValFromRedis(ctx, maxlon, maxlat, level)
	if err != nil {
		return 0, err
	}
	h3, err := redisClient.GetAltitudeValFromRedis(ctx, maxlon, minlat, level)
	if err != nil {
		return 0, err
	}
	h4, err := redisClient.GetAltitudeValFromRedis(ctx, minlon, minlat, level)
	if err != nil {
		return 0, err
	}

	fxy1 := (maxlon-position.Longitude)*h4/(maxlon-minlon) + (position.Longitude-minlon)*h3/(maxlon-minlon)
	fxy2 := (maxlon-position.Longitude)*h1/(maxlon-minlon) + (position.Longitude-minlon)*h2/(maxlon-minlon)
	originalDem := (maxlat-position.Latitude)*fxy1/(maxlat-minlat) + (position.Latitude-minlat)*fxy2/(maxlat-minlat)
	fmt.Printf("原高度：%f\n", originalDem)

	return originalDem, nil
}

func (redisClient *RedisClient) GetDynamicAltitude(ctx context.Context, position data.LonLatPosition) (float64, error) {
	craters, distance, err := redisClient.GetCratersNearby(ctx, position, 10)
	if err != nil {
		return 0, err
	}
	depth := 0.0
	for idx, crater := range craters {
		r := crater.Width / 2
		depth += (r - distance[idx]) / r * crater.Depth
	}
	return depth, nil
}

// 获取某点所在瓦片的 X-Y 坐标
func GetTileXY(lon, lat float64, level int) (int, int) {
	n := math.Pow(2, float64(level+1))
	x := (lon + 180) / 360 * n
	y := (lat + 90) / 360 * n

	X := int(math.Floor(x))
	Y := int(math.Floor(y))

	return X, Y
}

// 将 X-Y 坐标转换为经纬度
func GetLonLat(X, Y, level int) (float64, float64) {
	n := math.Pow(2, float64(level+1))
	lon := float64(X)/n*360 - 180
	lat := float64(Y)/n*360 - 90

	return lon, lat
}

// 获取某点所在网格的四个顶点经纬度坐标
func GetNeignbor(position data.LonLatPosition, level int, data *[]float64) {
	tileX, tileY := GetTileXY(position.Longitude, position.Latitude, level)
	tileLon, tileLat := GetLonLat(tileX, tileY, level)
	pixels := 65
	LonPerPixel := 360 / (math.Pow(2, float64(level+1)) * float64(pixels))
	LatPerPixel := LonPerPixel

	pixelLocatX := math.Floor((position.Longitude - tileLon) / LonPerPixel)
	pixelLocatY := math.Floor((position.Latitude - tileLat) / LatPerPixel)
	minLon := pixelLocatX*LonPerPixel + tileLon
	maxLon := minLon + LonPerPixel
	minLat := pixelLocatY*LatPerPixel + tileLat
	maxLat := minLat + LatPerPixel

	*data = append(*data, minLon, maxLon, minLat, maxLat)
}

func (redisClient *RedisClient) GetAltitudeValFromRedis(ctx context.Context, lon, lat float64, level int) (float64, error) {
	key := fmt.Sprintf("%d-%f-%f", level, lon, lat)
	val, err := redisClient.client.HGet(ctx, key, "TERRAIN").Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(val, 64)
}

func (redisClient *RedisClient) InsertCrater(ctx context.Context, crater data.Crater) error {
	craterId := "crater-" + crater.CraterID
	err := redisClient.client.GeoAdd(ctx, "craters", &redis.GeoLocation{
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

func (redisClient *RedisClient) GetCratersNearby(ctx context.Context, position data.LonLatPosition, radius float64) ([]data.Crater, []float64, error) {
	result, err := redisClient.client.GeoRadius(ctx, "craters", position.Longitude, position.Latitude, &redis.GeoRadiusQuery{
		Radius:      radius,
		Unit:        "m",   // 单位：米
		WithCoord:   false, // 不返回成员的坐标
		WithDist:    false, // 不返回成员与中心的距离
		WithGeoHash: false, // 不返回成员的 GeoHash
		Count:       100,   // 限制返回的成员数量
		Sort:        "ASC", // 按距离升序排序
	}).Result()
	if err != nil {
		fmt.Println("Error querying geo radius:", err)
		return nil, nil, nil
	}

	craters := make([]data.Crater, 0)
	distance := make([]float64, 0)
	for _, members := range result {
		crater, err := redisClient.GetCraterParamsByCraterID(ctx, members.Name)
		if err != nil {
			return nil, nil, err
		}
		craters = append(craters, crater)
		distance = append(distance, members.Dist)
	}
	return craters, distance, nil
}

func (redisClient *RedisClient) GetCraterParamsByCraterID(ctx context.Context, craterID string) (data.Crater, error) {
	var crater data.Crater
	crater.CraterID = craterID
	result, err := redisClient.client.HGetAll(ctx, craterID).Result()
	if err != nil {
		return crater, err
	}
	for key, value := range result {
		if key == "width" {
			crater.Width, err = strconv.ParseFloat(value, 64)
		} else if key == "depth" {
			crater.Depth, err = strconv.ParseFloat(value, 64)
		} else if key == "longitude" {
			crater.Position.Longitude, err = strconv.ParseFloat(value, 64)
		} else if key == "latitude" {
			crater.Position.Latitude, err = strconv.ParseFloat(value, 64)
		}
		if err != nil {
			return crater, err
		}
	}
	return crater, nil
}
