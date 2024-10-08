package test

import (
	"encoding/json"
	pb "env_middleware/grpc_env_service"
	"env_middleware/service"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// range of the map
const minLon, minLat = 113.07256, 21.8791665
const maxLon, maxLat = 113.7121265, 22.4699971

type tester struct {
	log       *os.File
	obstacles *os.File
	generator *BarrierGenerator
}

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

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func (g *BarrierGenerator) ReadRoutes(dirPath string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("fail to read dirPath %v", dirPath))
	}
	for _, file := range files {
		if !file.IsDir() {
			fileData, err := os.ReadFile(dirPath + "/" + file.Name())
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
}

func (g *BarrierGenerator) GetRandomCoordinate(inRoute bool) Coordinates {
	// 设置随机种子
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	if inRoute {
		// 随机抽取一个坐标
		randomIndex := r.Intn(len(g.coordinates))
		randomCoordinate := g.coordinates[randomIndex]
		return Coordinates{
			Longitude: randomCoordinate[0],
			Latitude:  randomCoordinate[1],
		}
	} else {
		lon := minLon + rand.Float64()*(maxLon-minLon)
		lat := minLat + rand.Float64()*(maxLat-minLat)
		return Coordinates{
			Longitude: lon,
			Latitude:  lat,
		}
	}
}

func (t *tester) getRouteAndRecord(groupID int, carID int, client pb.EnvironmentDataClient, start, end Coordinates) {
	for {
		startStopPoints := service.MakeStartStopPoints(start.Longitude, start.Latitude, end.Longitude, end.Latitude)
		position, err := service.CallGetRoutePointsAndRecord(client, startStopPoints)
		if err != nil {
			t.log.WriteString(fmt.Sprintf("GroupID %d, carID %d get route failed, err: %v", groupID, carID, err))
			log.Printf("GroupID %d, carID %d get route failed, err: %v", groupID, carID, err)
			return
		}
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%d_%d.json", groupID, carID, timestamp)
		writePoints(position, fileName)
		_, err = t.log.WriteString(fmt.Sprintf("%d: successfully write %d points in file %s\n", timestamp, len(position), fileName))
		if err != nil {
			fmt.Printf("fail to write conclusion to file %s, error: %s\n", fileName, err)
		}
		time.Sleep(time.Duration((5 + rand.Intn(5)) * int(time.Second)))
	}
}

func (t *tester) addBarrierInRoutes(client pb.EnvironmentDataClient) {
	for {
		coordinates := t.generator.GetRandomCoordinate(true)
		timestamp := time.Now().Unix()
		t.obstacles.WriteString(fmt.Sprintf("%d, %f, %f\n", timestamp, coordinates.Longitude, coordinates.Latitude))
		t.log.WriteString(fmt.Sprintf("%d: add obstacle in route: (%f-%f)\n", timestamp, coordinates.Longitude, coordinates.Latitude))
		obstacle := service.MakeObstacle(coordinates.Longitude, coordinates.Latitude, "road attack")
		service.CallUpdateObstacles(client, obstacle)
		time.Sleep(time.Duration((3 + rand.Float64()*3.0) * float64(time.Second)))
	}
}

func (t *tester) addBarrierInMap(client pb.EnvironmentDataClient) {
	for {
		coordinates := t.generator.GetRandomCoordinate(false)
		timestamp := time.Now().Unix()
		t.obstacles.WriteString(fmt.Sprintf("%d, %f, %f\n", timestamp, coordinates.Longitude, coordinates.Latitude))
		t.log.WriteString(fmt.Sprintf("%d: add obstacle in map: (%f-%f)\n", timestamp, coordinates.Longitude, coordinates.Latitude))
		obstacle := service.MakeObstacle(coordinates.Longitude, coordinates.Latitude, "road attack")
		service.CallUpdateObstacles(client, obstacle)
		time.Sleep(time.Duration((0.5 + rand.Float64()*1.0) * float64(time.Second)))
	}
}

func TestOSRM(client pb.EnvironmentDataClient) {
	log_file, err := os.Create("log.txt")
	if err != nil {
		fmt.Printf("fail to create log_file, err: %s\n", err)
		return
	}
	obstacles_file, err := os.Create("obstacles.txt")
	if err != nil {
		fmt.Printf("fail to create obstacle_file, err: %s\n", err)
		return
	}
	t := tester{
		log:       log_file,
		obstacles: obstacles_file,
		generator: &BarrierGenerator{
			coordinates: make([][]float64, 0),
		},
	}
	// read routes
	t.generator.ReadRoutes("/home/lml/graduation/env_middleware/routes")

	// generate barriers
	for i := 0; i < 100; i++ {
		go t.addBarrierInRoutes(client)
		go t.addBarrierInMap(client)
	}

	// get routes
	pairs := [][]Coordinates{
		// 1.珠海站-金湾机场
		{
			{Longitude: 113.5439406, Latitude: 22.2180959},        // 起点
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 终点
		},
		// 2.珠海站-横琴站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 终点
		},
		// 3.珠海站-十字门站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5163374, Latitude: 22.1756236}, // 终点
		},
		// 4.珠海站-前山站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748}, // 终点
		},
		// 5.珠海站-明珠站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963}, // 终点
		},
		// 6.金湾机场-横琴站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5396497, Latitude: 22.1410194},        // 终点
		},
		// 7.金湾机场-十字门站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5163374, Latitude: 22.1756236},        // 终点
		},
		// 8.金湾机场-前山站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748},        // 终点
		},
		// 9.金湾机场-明珠站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963},        // 终点
		},
		// 10.横琴站-十字门站
		{
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 起点
			{Longitude: 113.5163374, Latitude: .1756236},   // 终点
		},
		// 11.横琴站-前山站
		{
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748}, // 终点
		},
		// 12.横琴站-明珠站
		{
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963}, // 终点
		},
		// 13.十字门站-前山站
		{
			{Longitude: 113.5163374, Latitude: 22.1756236}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748}, // 终点
		},
		// 14.十字门站-明珠站
		{
			{Longitude: 113.5163374, Latitude: 22.1756236}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963}, // 终点
		},
		// 15.前山站-明珠站
		{
			{Longitude: 113.5205738, Latitude: 22.2373748},  // 起点
			{Longitude: 113.5103774, Latitude: 22.27139634}, // 终点
		},
	}
	for groupID := 0; groupID < len(pairs); groupID++ {
		for carID := 0; carID < 10; carID++ {
			go t.getRouteAndRecord(groupID+1, carID+1, client, pairs[groupID][0], pairs[groupID][1])
		}
	}
}

func writePoints(points []*pb.Position, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("fail to create file %s, err: %s\n", fileName, err)
		return
	}
	defer file.Close()
	var pointsArray [][]float64
	for _, point := range points {
		pointsArray = append(pointsArray, []float64{point.Longitude, point.Latitude})
	}
	// 构造 GeoJSON 数据
	geoJSON := FeatureCollection{
		Type: "FeatureCollection",
		Features: []Feature{
			{
				Type:       "Feature",
				Properties: struct{}{}, // 空 properties
				Geometry: Geometry{
					Type:        "LineString",
					Coordinates: pointsArray,
				},
			},
		},
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // 设置缩进格式
	err = encoder.Encode(geoJSON)
	if err != nil {
		fmt.Println("fail to encode, err:", err)
	}
}

func RunOriginalRoute(client pb.EnvironmentDataClient) {
	pairs := [][]Coordinates{
		// 1.珠海站-金湾机场
		{
			{Longitude: 113.5439406, Latitude: 22.2180959},        // 起点
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 终点
		},
		// 2.珠海站-横琴站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 终点
		},
		// 3.珠海站-十字门站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5163374, Latitude: 22.1756236}, // 终点
		},
		// 4.珠海站-前山站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748}, // 终点
		},
		// 5.珠海站-明珠站
		{
			{Longitude: 113.5439372, Latitude: 22.2180642}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963}, // 终点
		},
		// 6.金湾机场-横琴站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5396497, Latitude: 22.1410194},        // 终点
		},
		// 7.金湾机场-十字门站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5163374, Latitude: 22.1756236},        // 终点
		},
		// 8.金湾机场-前山站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748},        // 终点
		},
		// 9.金湾机场-明珠站
		{
			{Longitude: 113.37242495504282, Latitude: 22.0084784}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963},        // 终点
		},
		// 10.横琴站-十字门站
		{
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 起点
			{Longitude: 113.5163374, Latitude: .1756236},   // 终点
		},
		// 11.横琴站-前山站
		{
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748}, // 终点
		},
		// 12.横琴站-明珠站
		{
			{Longitude: 113.5396497, Latitude: 22.1410194}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963}, // 终点
		},
		// 13.十字门站-前山站
		{
			{Longitude: 113.5163374, Latitude: 22.1756236}, // 起点
			{Longitude: 113.5205738, Latitude: 22.2373748}, // 终点
		},
		// 14.十字门站-明珠站
		{
			{Longitude: 113.5163374, Latitude: 22.1756236}, // 起点
			{Longitude: 113.5103774, Latitude: 22.2713963}, // 终点
		},
		// 15.前山站-明珠站
		{
			{Longitude: 113.5205738, Latitude: 22.2373748},  // 起点
			{Longitude: 113.5103774, Latitude: 22.27139634}, // 终点
		},
	}
	for i, pair := range pairs {
		startStopPoints := service.MakeStartStopPoints(pair[0].Longitude, pair[0].Latitude, pair[1].Longitude, pair[1].Latitude)
		points, err := service.CallGetRoutePointsAndRecord(client, startStopPoints)
		if err != nil {
			log.Printf("point %d get route failed, err: %v", i, err)
		} else {
			writePoints(points, fmt.Sprintf("route-%d.json", i+1))
		}
	}
}
