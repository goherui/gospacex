package main

import (
	"context"
	"flag"
	"fmt"
	"gospacex/order-service/basic/initializer"
	_ "gospacex/order-service/basic/initializer"
	"gospacex/order-service/handler/service"
	__ "gospacex/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	if err := initializer.ConsulInit(); err != nil {
		log.Fatalf("Consul初始化失败: %v", err)
	}
	log.Println("Consul初始化成功")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	services, err := initializer.GetServiceWithLoadBalancer("user-service")
	if err != nil {
		log.Printf("获取用户服务失败: %v", err)
	} else {
		log.Printf("获取到用户服务: %s, 地址: %s:%d", services.Service, services.Address, services.Port)
	}

	s := grpc.NewServer()
	__.RegisterStreamGreeterServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务...")
	if err := initializer.ConsulShutdown(); err != nil {
		log.Printf("Consul注销失败: %v", err)
	}
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.GracefulStop()
	log.Println("服务已关闭")
}
