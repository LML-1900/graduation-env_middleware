package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetLonLat(X, Y, level int) (lon, lat float64) {
	n := math.Pow(2, float64(level+1))
	lon = float64(X)/n*360 - 180
	lat = float64(Y)/n*360 - 90
	return lon, lat
}

func ParseTileID(tileID string) (level, X, Y int, err error) {
	// 以'-'为分隔符拆分字符串
	parts := strings.Split(tileID, "-")
	if len(parts) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid tileID format")
	}
	// 转换level部分
	level, err = strconv.Atoi(parts[0][5:])
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
