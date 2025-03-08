package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	pb "env_middleware/grpc_env_service"
	"env_middleware/service"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	rdb        *redis.Client
	mq         *service.RabbitMq
	s          *service.RPCService
	grpcClient pb.EnvironmentDataClient
	conn       *grpc.ClientConn
	stopSignal chan struct{}
)

// 消费消息的上下文
var consumerCtx context.Context
var cancelConsumer context.CancelFunc

//export InitMiddleware
func InitMiddleware(serverUrl, redisUrl, rabbitmqUrl string) {
	// 初始化 stopSignal 和 context
	stopSignal = make(chan struct{})
	consumerCtx, cancelConsumer = context.WithCancel(context.Background())

	// set viper
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// set Redis
	var redisUrlString string
	if redisUrl == "" {
		redisUrlString = viper.GetString("redis.url")
	} else {
		redisUrlString = redisUrl
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
	if rabbitmqUrl == "" {
		rabbitmqUrlString = viper.GetString("rabbitmq.url")
	} else {
		rabbitmqUrlString = rabbitmqUrl
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
	if serverUrl == "" {
		addr = viper.GetString("grpc.addr")
	} else {
		addr = serverUrl
	}
	flag.Parse()
	// Set up a connection to the server.
	conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to  connect to server %s, error: %v", addr, err)
	}
	grpcClient = pb.NewEnvironmentDataClient(conn)

	// // start consume
	// go func() {
	// 	for {
	// 		select {
	// 		case <-stopSignal:
	// 			fmt.Println("Stopping message consumption...")
	// 			return
	// 		default:
	// 			mq.ConsumeMsgs(consumerCtx)
	// 		}
	// 	}
	// }()
}

func main() {
	serverUrl := "10.134.114.218:50052"
	rabbitmqUrl := "amqp://guest:guest@10.134.114.218:5672"
	InitMiddleware(serverUrl, "", rabbitmqUrl)

	// 监听 TCP 端口
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on 127.0.0.1:8080")

	// startStopPoints := service.MakeStartStopPoints(113.5439372, 22.2180642, 113.5425177, 22.2252363)
	// fmt.Printf("start: (%f, %f) stop: (%f, %f)\n", startStopPoints.Start.Longitude, startStopPoints.Start.Latitude, startStopPoints.End.Longitude, startStopPoints.End.Latitude)
	// message := service.CallGetRoutePoints(grpcClient, startStopPoints)
	// if message == nil {
	// 	fmt.Println("no route!")
	// } else {
	// 	fmt.Printf("got %d points\n", len(message.Pos))
	// }

	for {
		// 接受客户端连接
		socket_conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(socket_conn) // 处理客户端连接
	}
}

type Request struct {
	Function string          `json:"function"`
	Data     json.RawMessage `json:"data"` // Store function-specific parameters
}

type Response struct {
	Length int          `json:"length"`
	Path   [][2]float64 `json:"path"`
}

func handleConnection(socket_conn net.Conn) {
	defer socket_conn.Close()
	fmt.Println("Client connected:", socket_conn.RemoteAddr())

	buffer := make([]byte, 1024)
	for {
		// 设置读取超时时间，例如 5 分钟
		socket_conn.SetReadDeadline(time.Now().Add(5 * time.Minute))

		n, err := socket_conn.Read(buffer)
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Println("Read timeout, closing connection:", socket_conn.RemoteAddr())
			} else if err.Error() == "EOF" {
				fmt.Println("Client disconnected:", socket_conn.RemoteAddr())
			} else {
				fmt.Println("Error reading data:", err)
			}
			return
		}

		// 解析请求和处理逻辑
		var req Request
		err = json.Unmarshal(buffer[:n], &req)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}

		var resp Response
		needResponse := true
		switch req.Function {
		case "routePlanning":
			var routeParams struct {
				Start [2]float64 `json:"start"`
				End   [2]float64 `json:"end"`
			}
			err := json.Unmarshal(req.Data, &routeParams)
			if err != nil {
				fmt.Println("Error decoding routePlanning data:", err)
				continue
			}
			startStopPoints := service.MakeStartStopPoints(routeParams.Start[0], routeParams.Start[1], routeParams.End[0], routeParams.End[1])
			message := service.CallGetRoutePoints(grpcClient, startStopPoints)
			path := make([][2]float64, 0)
			if message != nil {
				path = make([][2]float64, len(message.Pos))
				for i, pos := range message.Pos {
					path[i][0] = pos.Longitude
					path[i][1] = pos.Latitude
					// fmt.Printf("lon: %v, lat: %v\n", path[i][0], path[i][1])
				}
			}
			resp = Response{
				Length: len(path),
				Path:   path,
			}

		case "updateObstacle":
			var obstacleParams struct {
				Location [2]float64 `json:"location"`
				Cause    string
			}
			err := json.Unmarshal(req.Data, &obstacleParams)
			if err != nil {
				fmt.Println("Error decoding updateObstacle data:", err)
				continue
			}
			fmt.Printf("received obstacle: pos(%v, %v), cause(%v)\n", obstacleParams.Location[0], obstacleParams.Location[1], obstacleParams.Cause)
			obstacle := service.MakeObstacle(obstacleParams.Location[0], obstacleParams.Location[1], obstacleParams.Cause)
			service.CallUpdateObstacles(grpcClient, obstacle)
			needResponse = false

		default:
			fmt.Println("Unknown function:", req.Function)
		}
		// 发送响应
		if !needResponse {
			continue
		}
		respData, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			continue
		}

		err = sendMessage(socket_conn, respData)
		if err != nil {
			fmt.Println("Error sending response:", err)
			return
		}
	}
}

func sendMessage(conn net.Conn, data []byte) error {
	// Step 1: 构造头部
	header := make([]byte, 4) // 4字节头部
	messageSize := len(data)
	binary.BigEndian.PutUint32(header, uint32(messageSize)) // 网络字节序

	// Step 2: 发送头部
	_, err := conn.Write(header)
	if err != nil {
		return fmt.Errorf("failed to send header: %w", err)
	}

	// Step 3: 发送消息体
	_, err = conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to send message body: %w", err)
	}

	return nil
}
