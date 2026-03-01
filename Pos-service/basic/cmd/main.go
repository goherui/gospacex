package main

import (
	"context"
	"flag"
	"fmt"
	"gospaacex/Pos-service/basic/config"
	"gospaacex/Pos-service/basic/initializer"
	"gospaacex/Pos-service/handler/service"
	__ "gospaacex/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	_ "gospaacex/Pos-service/basic/initializer"
	"log"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func loggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		md, _ := metadata.FromIncomingContext(ctx)
		log.Printf("收到请求: %s, 入参: %v, 元数据: %v", info.FullMethod, req, md)
		resp, err := handler(ctx, req)
		if err != nil {
			log.Printf("请求处理失败: %s, 错误: %v, 耗时: %v", info.FullMethod, err, time.Since(start))
		} else {
			log.Printf("请求处理成功: %s, 响应: %v, 耗时: %v", info.FullMethod, resp, time.Since(start))
		}
		return resp, err
	}
}
func main() {
	flag.Parse()
	if err := initializer.InitConsul(); err != nil {
		log.Fatalf("初始化 Consul 失败: %v", err)
	}
	log.Println("Consul 初始化成功")
	services, err := initializer.GetService(config.GlobalConfig.Consul.ServiceName)
	if err != nil {
		log.Printf("获取服务失败: %v", err)
	} else {
		log.Printf("找到 %d 个服务实例", len(services))
		for _, srv := range services {
			log.Printf("  - 实例: ID=%s, Address=%s, Port=%d", srv.ID, srv.Address, srv.Port)
		}
	}
	log.Println("\n2. 使用负载均衡获取服务实例:")
	srv, err := initializer.GetServiceWithLoadBalancer(config.GlobalConfig.Consul.ServiceName)
	if err != nil {
		log.Printf("获取服务失败: %v", err)
	} else {
		address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
		log.Printf("负载均衡选择的服务: %s", address)
	}
	log.Println("\n3. 获取所有注册的服务:")
	allServices, err := initializer.GetAllServices()
	if err != nil {
		log.Printf("获取所有服务失败: %v", err)
	} else {
		log.Printf("找到 %d 个服务类型", len(allServices))
		for serviceName, instances := range allServices {
			log.Printf("  - 服务 %s: %d 个实例", serviceName, len(instances))
		}
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() {
		err = initializer.DeregisterService()
		if err != nil {
			log.Printf("服务注销失败:%v", err)
		} else {
			log.Println("服务注销成功")
		}
	}()
	s := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor()),
	)
	__.RegisterStreamGreeterServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//func main() {
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer(
//		grpc.UnaryInterceptor(loggingInterceptor()),
//	)
//	__.RegisterStreamGreeterServer(s, &service.Server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
