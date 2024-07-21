package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetLonLat(X, Y, level int) (lon, lat float64) {
	n := math.Pow(2, float64(level+1))
	// GPT considers code follows is right
	// but I think here must be level+1, cause there are 2 tiles under level0
	//n := math.Pow(2, float64(level))

	lon = float64(X)/n*360.0 - 180.0
	//lat = float64(Y)/n*360.0 - 90.0
	// Calculate latitude, gpt says lat must consider Web Mercator 投影
	latRad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(Y)/n)))
	lat = latRad * 180.0 / math.Pi
	return lon, lat
}

func ParseTileID(tileID string) (level, X, Y int, err error) {
	// 以'-'为分隔符拆分字符串
	parts := strings.Split(tileID, "-")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid tileID format")
	}
	// 转换level部分
	level, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid level value: %v", err)
	}

	// 转换X部分
	X, err = strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid X value: %v", err)
	}

	// 转换Y部分
	Y, err = strconv.Atoi(parts[2])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("invalid Y value: %v", err)
	}

	return level, X, Y, nil
}
