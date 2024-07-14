package main

import (
	"context"
	"env_middleware/data"
	pb "env_middleware/grpc_env_service"
	"env_middleware/service"
	"flag"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

//var (
//	addr = flag.String("addr", "localhost:50052", "the address to connect to")
//	//addr = flag.String("addr", "10.134.92.104:50052", "the address to connect to")
//	//addr = flag.String("addr", "10.134.114.97:50052", "the address to connect to")
//)

func main() {
	// set viper
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// set Redis
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"), // Redis服务器地址
		Password: "",                           // 没有设置密码
		DB:       0,                            // 使用默认数据库
	})
	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// set RabbitMQ
	mq, err := service.RunRabbitMqConsumer("dynamic_data_topic", rdb)

	if err != nil {
		log.Fatalf("Failed to declare an exchange, err:%v", err)
	}
	defer mq.Conn.Close()
	defer mq.Ch.Close()

	// connect to server
	addr := os.Getenv("SERVER_ADDRESS")
	if addr == "" {
		addr = "localhost:50051"
	}
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEnvironmentDataClient(conn)

	// request := service.MakeStaticDataRequest(34.46, 78.41, 34.79, 78.74, 14, data.DEM_DATA_TYPE)
	request := service.MakeStaticDataRequest(34.46, 78.41, 34.491, 78.448, 14, data.DEM_DATA_TYPE)
	service.CallGetStaticDataRequestRPC(c, request)

	//crater := service.MakeCrater(78.45, 34.49, 5.33, 4.32)
	//service.CallUpdateCrater(c, crater)

	//startStopPoints := service.MakeStartStopPoints(113.5439372, 22.2180642, 113.5425177, 22.2252363)
	//service.CallGetRoutePoints(c, startStopPoints)

	obstacle := service.MakeObstacle(113.416793, 22.158472, "road attack")
	service.CallUpdateObstacles(c, obstacle)

	mq.ConsumeMsgs()
	var forever chan struct{}
	<-forever
}
