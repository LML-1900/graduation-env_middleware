// only for test
package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type BarrierGenerator struct {
	coordinates [][]float64
}

type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string   `json:"type"`
	Properties struct{} `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"` // 注意这里是二维切片
}

func (g *BarrierGenerator) ReadRoutes(filePath []string) {
	for _, fileName := range filePath {
		fileData, err := os.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		var featureCollection FeatureCollection
		err = json.Unmarshal(fileData, &featureCollection)
		if err != nil {
			panic(err)
		}

		// 获取坐标列表
		for _, feature := range featureCollection.Features {
			g.coordinates = append(g.coordinates, feature.Geometry.Coordinates...)
		}
	}
}

func (g *BarrierGenerator) GetRandomCoordinate(inRoute bool) {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	if inRoute {
		// 随机抽取一个坐标
		randomIndex := rand.Intn(len(g.coordinates))
		randomCoordinate := g.coordinates[randomIndex]
		fmt.Printf("随机抽取的坐标: %v\n", randomCoordinate)
	} else {

	}
}
