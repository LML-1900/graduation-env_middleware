package main

import (
	pb "env_middleware/grpc_env_service"
	"env_middleware/service"
	"flag"
	"fmt"
	"github.com/spf13/viper"

	//"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	// set viper
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	mq, err := service.RunRabbitMqConsumer("dynamic_data_topic")

	if err != nil {
		log.Fatalf("Failed to declare an exchange, err:%v", err)
	}
	defer mq.Conn.Close()
	defer mq.Ch.Close()

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEnvironmentDataClient(conn)

	//request := service.MakeStaticDataRequest(34.46, 78.41, 34.79, 78.74, 14, data.DEM_DATA_TYPE)
	//service.CallGetStaticDataRequestRPC(c, request)

	crater := service.MakeCrater(78.45, 34.49, 5.33, 4.32)
	service.CallUpdateCrater(c, crater)
	crater = service.MakeCrater(78.41, 34.43, 2.78, 10.78)
	service.CallUpdateCrater(c, crater)

	mq.ConsumeMsgs()
	var forever chan struct{}
	<-forever
}
