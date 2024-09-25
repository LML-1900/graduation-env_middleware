package main

import (
	"context"
	"env_middleware/data"
	pb "env_middleware/grpc_env_service"
	"env_middleware/service"
	"flag"
	"fmt"
	"log"
	"unsafe"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
import "C"

//var (
//	addr = flag.String("addr", "localhost:50052", "the address to connect to")
//	//addr = flag.String("addr", "10.134.92.104:50052", "the address to connect to")
//	//addr = flag.String("addr", "10.134.114.97:50052", "the address to connect to")
//)

var (
	rdb        *redis.Client
	mq         *service.RabbitMq
	s          *service.RPCService
	grpcClient pb.EnvironmentDataClient
	conn       *grpc.ClientConn
	stopSignal chan struct{}
)

//export InitMiddleware
func InitMiddleware(serverUrl, redisUrl, rabbitmqUrl *C.char) {
	// set viper
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// set Redis
	var redisUrlString string
	if redisUrl == nil {
		redisUrlString = viper.GetString("redis.url")
	} else {
		redisUrlString = C.GoString(redisUrl)
	}
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisUrlString, // Redis服务器地址
		Password: "",             // 没有设置密码
		DB:       0,              // 使用默认数据库
	})
	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// set RabbitMQ
	var rabbitmqUrlString string
	if rabbitmqUrl == nil {
		rabbitmqUrlString = viper.GetString("rabbitmq.url")
	} else {
		rabbitmqUrlString = C.GoString(rabbitmqUrl)
	}
	mq, err = service.RunRabbitMqConsumer("dynamic_data_topic", rdb, rabbitmqUrlString)
	if err != nil {
		log.Fatalf("Failed to declare an rabbitmq exchange, err:%v", err)
		return
	}

	// init service
	s = service.NewRPCService(rdb)

	// connect to server
	var addr string
	if serverUrl == nil {
		addr = viper.GetString("grpc.addr")
	} else {
		addr = C.GoString(serverUrl)
	}
	flag.Parse()
	// Set up a connection to the server.
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to  connect to server %s, error: %v", addr, err)
	}
	grpcClient = pb.NewEnvironmentDataClient(conn)

	// start consume
	go func() {
		for {
			select {
			case <-stopSignal:
				fmt.Println("Stopping message consumption...")
				return
			default:
				mq.ConsumeMsgs()
			}
		}
	}()
}

//export CloseMiddleware
func CloseMiddleware() {
	mq.Conn.Close()
	mq.Ch.Close()
	conn.Close()
	close(stopSignal)
}

//export GetRawData
func GetRawData(minLat, minLon, maxLat, maxLon C.double, level C.int) {
	request := service.MakeStaticDataRequest(float64(minLat), float64(minLon), float64(maxLat), float64(maxLon), int(level), data.DEM_DATA_TYPE)
	s.CallGetStaticDataRequestRPC(grpcClient, request)
}

// 导出为 C 语言兼容的结构体
type C_LonLatPosition struct {
	Longitude C.double
	Latitude  C.double
}

// // export GetRoutePoints
// func GetRoutePoints(startLatitude, startLongitude, stopLatitude, stopLongitude C.double) (*C_LonLatPosition, C.int) {
// 	startStopPoints := service.MakeStartStopPoints(startLongitude, startLatitude, stopLongitude, stopLatitude)
// 	message := service.CallGetRoutePoints(grpcClient, startStopPoints)
// 	// 分配 C 语言兼容的数组
// 	cPositions := C.malloc(C.size_t(len(message.Pos)) * C.size_t(unsafe.Sizeof(C_LonLatPosition{})))
// 	for i, pos := range message.Pos {
// 		cPositions[i] = C_LonLatPosition{
// 			Longitude: C.double(pos.Longitude),
// 			Latitude:  C.double(pos.Latitude),
// 		}
// 	}
// 	// 返回指向数组第一个元素的指针，以及数组的长度
// 	return (*C_LonLatPosition)(unsafe.Pointer(&cPositions[0])), C.int(len(cPositions))
// }

//export GetRoutePoints
func GetRoutePoints(startLongitude, startLatitude, stopLongitude, stopLatitude C.double) (C.int, *C.double, *C.double) {
	startStopPoints := service.MakeStartStopPoints(float64(startLongitude), float64(startLatitude), float64(stopLongitude), float64(stopLatitude))
	fmt.Printf("start: (%f, %f) stop: (%f, %f)\n", startStopPoints.Start.Longitude, startStopPoints.Start.Latitude, startStopPoints.End.Longitude, startStopPoints.End.Latitude)
	message := service.CallGetRoutePoints(grpcClient, startStopPoints)
	if message == nil {
		return 0, nil, nil
	}
	length := len(message.Pos)
	longitude_arr := (*[1 << 30]C.double)(C.malloc(C.size_t(length) * C.size_t(unsafe.Sizeof(C.double(0)))))
	latitude_arr := (*[1 << 30]C.double)(C.malloc(C.size_t(length) * C.size_t(unsafe.Sizeof(C.double(0)))))

	for i, pos := range message.Pos {
		longitude_arr[i] = C.double(pos.Longitude)
		latitude_arr[i] = C.double(pos.Latitude)
	}
	return C.int(length), &longitude_arr[0], &latitude_arr[0]
}

//export FreePositionsPointer
func FreePositionsPointer(longitude_arr *C.double, latitude_arr *C.double) {
	C.free(unsafe.Pointer(longitude_arr))
	C.free(unsafe.Pointer(latitude_arr))
}

//export UpdateObstacle
func UpdateObstacle(latitude, longitude C.double, cause *C.char) {
	obstacle := service.MakeObstacle(float64(longitude), float64(latitude), C.GoString(cause))
	service.CallUpdateObstacles(grpcClient, obstacle)
}

//export UpdateCrater
func UpdateCrater(longitude, latitude, width, depth C.double) {
	crater := service.MakeCrater(float64(longitude), float64(latitude), float64(width), float64(depth))
	service.CallUpdateCrater(grpcClient, crater)
}

//export GetAltitude
func GetAltitude(longitude, latitude C.double) C.double {
	altitude, err := s.GetAltitude(float64(longitude), float64(latitude))
	if err != nil {
		fmt.Printf("fail to GetAltitude, err: %s\n", err)
		return 0
	}
	return C.double(altitude)
}

func main() {

	// // test InitMiddleware
	// var serverUrl *C.char = C.CString("10.134.114.97:50052")
	// var rabbitmqUrl *C.char = C.CString("amqp://guest:guest@10.134.114.97:5672")
	// // var redisUrl *C.char = C.CString("10.134.114.97:6379")
	// defer C.free(unsafe.Pointer(serverUrl))
	// defer C.free(unsafe.Pointer(rabbitmqUrl))
	// InitMiddleware(serverUrl, nil, rabbitmqUrl)

	// // test GetRawData
	// // GetRawData(34.46, 78.41, 34.47, 78.42, 14)

	// // test GetRoutePoints
	// len, latitude_arr, longitude_arr := GetRoutePoints(113.5439372, 22.2180642, 113.5425177, 22.2252363)
	// c_latitude_arr := (*[1 << 30]C.double)(unsafe.Pointer(latitude_arr))
	// c_longitude_arr := (*[1 << 30]C.double)(unsafe.Pointer(longitude_arr))
	// for i := 0; i < int(len); i++ {
	// 	fmt.Printf("lon:%f - lat:%f\n", float64(c_longitude_arr[i]), float64(c_latitude_arr[i]))
	// }
	// FreePositionsPointer(longitude_arr, latitude_arr)

	// // test UpdateObstacle
	// var cause *C.char = C.CString("road attack")
	// defer C.free(unsafe.Pointer(cause))
	// UpdateObstacle(113.416793, 22.158472, cause)

	// // test UpdateCrater
	// UpdateCrater(78.45, 34.49, 5.33, 4.32)

	// // test GetAltitude
	// altitude := GetAltitude(78.45, 34.39)
	// fmt.Printf("altitude at lonlat(%f, %f) is %f", 78.45, 34.39, float64(altitude))

	// CloseMiddleware()

	// request := service.MakeStaticDataRequest(34.46, 78.41, 34.79, 78.74, 14, data.DEM_DATA_TYPE)
	//request := service.MakeStaticDataRequest(34.46, 78.41, 34.491, 78.448, 14, data.DEM_DATA_TYPE)
	// request := service.MakeStaticDataRequest(34.46, 78.41, 34.47, 78.42, 14, data.DEM_DATA_TYPE)
	// s.CallGetStaticDataRequestRPC(c, request)

	// crater := service.MakeCrater(78.45, 34.49, 5.33, 4.32)
	// service.CallUpdateCrater(c, crater)

	// startStopPoints := service.MakeStartStopPoints(113.5439372, 22.2180642, 113.5425177, 22.2252363)
	// service.CallGetRoutePoints(c, startStopPoints)

	// obstacle := service.MakeObstacle(113.416793, 22.158472, "road attack")
	// service.CallUpdateObstacles(c, obstacle)

	// mq.ConsumeMsgs()
	// var forever chan struct{}
	// <-forever
	// altitude, err := s.GetAltitude(78.45, 34.39)
	// fmt.Printf("altitude: %v", altitude)
}
