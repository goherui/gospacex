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

type WebService interface {
	Start() error
	Stop() error
}
type GrpcService struct {
	port   int
	server *grpc.Server
	lis    net.Listener
}

func NewGrpcService(port int) *GrpcService {
	return &GrpcService{
		port: port,
	}
}

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

type HttpService struct {
	port   int
	server *http.Server
}

func NewHttpService(port int) *HttpService {
	return &HttpService{
		port: port,
	}
}

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

type HttpsService struct {
	port   int
	server *http.Server
}

func NewHttpsService(port int) *HttpsService {
	return &HttpsService{
		port: port,
	}
}

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

type WebSocketService struct {
	port   int
	server *http.Server
}

func NewWebSocketService(port int) *WebSocketService {
	return &WebSocketService{
		port: port,
	}
}

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

type WebServiceFactory struct{}

func NewWebServiceFactory() *WebServiceFactory {
	return &WebServiceFactory{}
}

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
