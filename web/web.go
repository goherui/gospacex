package web

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

// WebService 定义web服务接口
type WebService interface {
	Start() error
	Stop() error
}

// GrpcService gRPC服务实现
type GrpcService struct {
	port   int
	server *grpc.Server
	lis    net.Listener
}

// NewGrpcService 创建gRPC服务实例
func NewGrpcService(port int) *GrpcService {
	return &GrpcService{
		port: port,
	}
}

// Start 启动gRPC服务
func (s *GrpcService) Start() error {
	var err error
	s.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	s.server = grpc.NewServer()
	log.Printf("gRPC server listening at %v", s.lis.Addr())
	go func() {
		if err := s.server.Serve(s.lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return nil
}
func (s *GrpcService) Stop() error {
	log.Println("正在关闭gRPC服务...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.server.GracefulStop()
	log.Println("gRPC服务已关闭")
	return nil
}

// HttpService HTTP服务实现
type HttpService struct {
	port   int
	server *http.Server
}

// NewHttpService 创建HTTP服务实例
func NewHttpService(port int) *HttpService {
	return &HttpService{
		port: port,
	}
}

// Start 启动HTTP服务
func (s *HttpService) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}
	log.Printf("HTTP server listening at %s", s.server.Addr)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return nil
}

// Stop 停止HTTP服务
func (s *HttpService) Stop() error {
	log.Println("正在关闭HTTP服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}
	log.Println("HTTP服务已关闭")
	return nil
}

// HttpsService HTTPS服务实现
type HttpsService struct {
	port   int
	server *http.Server
}

// NewHttpsService 创建HTTPS服务实例
func NewHttpsService(port int) *HttpsService {
	return &HttpsService{
		port: port,
	}
}

// Start 启动HTTPS服务
func (s *HttpsService) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}
	log.Printf("HTTPS server listening at %s", s.server.Addr)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return nil
}

// Stop 停止HTTPS服务
func (s *HttpsService) Stop() error {
	log.Println("正在关闭HTTPS服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}
	log.Println("HTTPS服务已关闭")
	return nil
}

// WebSocketService WebSocket服务实现
type WebSocketService struct {
	port   int
	server *http.Server
}

// NewWebSocketService 创建WebSocket服务实例
func NewWebSocketService(port int) *WebSocketService {
	return &WebSocketService{
		port: port,
	}
}

// Start 启动WebSocket服务
func (s *WebSocketService) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("WebSocket service"))
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}
	log.Printf("WebSocket server listening at %s", s.server.Addr)
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return nil
}

// Stop 停止WebSocket服务
func (s *WebSocketService) Stop() error {
	log.Println("正在关闭WebSocket服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %v", err)
	}
	log.Println("WebSocket服务已关闭")
	return nil
}

// WebServiceFactory web服务工厂
type WebServiceFactory struct{}

// NewWebServiceFactory 创建web服务工厂实例
func NewWebServiceFactory() *WebServiceFactory {
	return &WebServiceFactory{}
}

// CreateService 根据类型创建web服务
func (f *WebServiceFactory) CreateService(serviceType string, port int) (WebService, error) {
	switch serviceType {
	case "grpc":
		return NewGrpcService(port), nil
	case "http":
		return NewHttpService(port), nil
	case "https":
		return NewHttpsService(port), nil
	case "websocket":
		return NewWebSocketService(port), nil
	default:
		return nil, fmt.Errorf("unsupported service type: %s", serviceType)
	}
}
