package service

import (
	"context"
	pb "env_middleware/grpc_env_service"
	"fmt"
	"io"
	"log"
	"time"
)

func MakeStaticDataRequest(minLatitude, minLongitude, maxLatitude, maxLongitude float64, level int, dataType int) *pb.GetStaticDataRequest {
	bottomLeft := pb.Position{
		Latitude:  minLatitude,
		Longitude: minLongitude,
	}
	topRight := pb.Position{
		Latitude:  maxLatitude,
		Longitude: maxLongitude,
	}
	area := pb.Area{
		Bottomleft: &bottomLeft,
		Topright:   &topRight,
	}
	request := pb.GetStaticDataRequest{
		Area:     &area,
		Level:    int32(level),
		DataType: pb.DataType(dataType),
	}
	return &request
}

func MakeCrater(longitude, latitude, width, depth float64) *pb.Crater {
	crater := pb.Crater{
		Pos: &pb.Position{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Width: width,
		Depth: depth,
	}
	return &crater
}

func MakeObstacle(longitude, latitude float64, cause string) *pb.Obstacle {
	obstacle := pb.Obstacle{
		Pos: &pb.Position{
			Latitude:  latitude,
			Longitude: longitude,
		},
		Cause: cause,
	}
	return &obstacle
}

func MakeStartStopPoints(startLongitude, startLatitude, stopLongitude, stopLatitude float64) *pb.StartStopPoints {
	startPoint := pb.Position{
		Latitude:  startLatitude,
		Longitude: startLongitude,
	}
	endPoint := pb.Position{
		Latitude:  stopLatitude,
		Longitude: stopLongitude,
	}
	startStopPoints := pb.StartStopPoints{
		Start: &startPoint,
		End:   &endPoint,
	}
	return &startStopPoints
}

func CallGetRoutePoints(client pb.EnvironmentDataClient, startStopPoints *pb.StartStopPoints) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	message, err := client.GetRoutePoints(ctx, startStopPoints)
	if err != nil {
		log.Fatalf("client.GetRoutePoints failed: %v", err)
	}
	fmt.Printf("total points: %v\n", len(message.Pos))
	for _, point := range message.Pos {
		fmt.Printf("lon:%v-lat:%v\n", point.Longitude, point.Latitude)
	}
}

func CallUpdateObstacles(client pb.EnvironmentDataClient, obstacle *pb.Obstacle) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	message, err := client.UpdateObstacle(ctx, obstacle)
	if err != nil {
		log.Fatalf("client.UpdateObstacle failed: %v", err)
	}
	fmt.Println(message)
}

func CallUpdateCrater(client pb.EnvironmentDataClient, crater *pb.Crater) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	message, err := client.UpdateCrater(ctx, crater)
	if err != nil {
		log.Fatalf("client.UpdateCrater failed: %v", err)
	}
	fmt.Println(message)
}

func CallGetStaticDataRequestRPC(client pb.EnvironmentDataClient, request *pb.GetStaticDataRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetStaticData(ctx, request)
	if err != nil {
		log.Fatalf("client.GetStaticData failed: %v", err)
	}
	count := 0
	for {
		tile, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.GetStaticData failed: %v", err)
		}
		fmt.Println(tile.TileID)
		count++
	}
	fmt.Println(count)
}
