package initializer

import (
	"fmt"
	"gospaacex/Pos-service/basic/config"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
)

var (
	consulClient    *api.Client
	serviceID       string
	serviceCache    map[string][]*api.AgentService
	cacheMutex      sync.RWMutex
	cacheExpiration time.Time
	cacheDuration   = 30 * time.Second
)

// InitConsul 初始化Consul并注册服务
func InitConsul() error {
	// 创建Consul客户端
	consulConfig := api.DefaultConfig()
	consulConfig.Address = fmt.Sprintf("%s:%d", config.GlobalConfig.Consul.Host, config.GlobalConfig.Consul.Port)

	var err error
	if consulClient, err = api.NewClient(consulConfig); err != nil {
		return fmt.Errorf("创建Consul客户端失败: %w", err)
	}
	// 初始化服务缓存
	serviceCache = make(map[string][]*api.AgentService)
	// 生成服务ID
	serviceID = fmt.Sprintf("%s-%d", config.GlobalConfig.Consul.ServiceName, time.Now().Unix())
	checkID := fmt.Sprintf("%s-health", serviceID)
	// 注册服务
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    config.GlobalConfig.Consul.ServiceName,
		Address: "localhost",
		Port:    config.GlobalConfig.Consul.ServicePort,
		Checks: []*api.AgentServiceCheck{
			{
				CheckID:                        checkID,
				Name:                           "TTL Health Check",
				TTL:                            fmt.Sprintf("%ds", config.GlobalConfig.Consul.TTL),
				DeregisterCriticalServiceAfter: "1m", // 自动注销
			},
		},
	}
	if err := consulClient.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("注册服务失败: %w", err)
	}
	// 启动健康检查
	go func() {
		ticker := time.NewTicker(time.Duration(config.GlobalConfig.Consul.TTL/2) * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			if err := consulClient.Agent().UpdateTTL(checkID, "服务正常", api.HealthPassing); err != nil {
				log.Printf("健康检查更新失败: %v", err)
			}
		}
	}()
	// 启动服务缓存更新
	go func() {
		ticker := time.NewTicker(cacheDuration)
		defer ticker.Stop()

		for range ticker.C {
			updateServiceCache()
		}
	}()
	return nil
}

// DeregisterService 注销服务
func DeregisterService() error {
	if consulClient == nil || serviceID == "" {
		return nil
	}
	if err := consulClient.Agent().ServiceDeregister(serviceID); err != nil {
		return fmt.Errorf("注销服务失败: %w", err)
	}
	log.Println("服务已从Consul注销")
	return nil
}

// GetService 获取服务实例列表
func GetService(serviceName string) ([]*api.AgentService, error) {
	// 先尝试从缓存获取
	cacheMutex.RLock()
	if time.Now().Before(cacheExpiration) && serviceCache != nil {
		if services, ok := serviceCache[serviceName]; ok {
			cacheMutex.RUnlock()
			return services, nil
		}
	}
	cacheMutex.RUnlock()
	// 缓存过期或不存在，重新获取
	services, err := consulClient.Agent().Services()
	if err != nil {
		return nil, fmt.Errorf("获取服务列表失败: %w", err)
	}
	var serviceInstances []*api.AgentService
	for _, service := range services {
		if service.Service == serviceName {
			serviceInstances = append(serviceInstances, service)
		}
	}
	// 更新缓存
	cacheMutex.Lock()
	serviceCache[serviceName] = serviceInstances
	cacheExpiration = time.Now().Add(cacheDuration)
	cacheMutex.Unlock()
	return serviceInstances, nil
}

// GetServiceWithLoadBalancer 带负载均衡的服务获取
func GetServiceWithLoadBalancer(serviceName string) (*api.AgentService, error) {
	// 获取健康的服务实例
	healthyServices, err := GetHealthyService(serviceName)
	if err != nil {
		return nil, err
	}
	if len(healthyServices) == 0 {
		return nil, fmt.Errorf("没有可用的健康服务实例")
	}
	// 随机负载均衡
	randomIndex := rand.Intn(len(healthyServices))
	return healthyServices[randomIndex], nil
}

// GetHealthyService 获取健康的服务实例
func GetHealthyService(serviceName string) ([]*api.AgentService, error) {
	healthChecks, _, err := consulClient.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, fmt.Errorf("健康检查失败: %w", err)
	}
	var healthyServices []*api.AgentService
	for _, check := range healthChecks {
		if check.Checks.AggregatedStatus() == api.HealthPassing {
			healthyServices = append(healthyServices, check.Service)
		} else {
			log.Printf("服务 %s 状态异常: %s", check.Service.ID, check.Checks.AggregatedStatus())
		}
	}
	return healthyServices, nil
}

// GetAllServices 获取所有注册的服务
func GetAllServices() (map[string][]*api.AgentService, error) {
	services, err := consulClient.Agent().Services()
	if err != nil {
		return nil, fmt.Errorf("获取所有服务失败: %w", err)
	}
	serviceMap := make(map[string][]*api.AgentService)
	for _, service := range services {
		serviceMap[service.Service] = append(serviceMap[service.Service], service)
	}
	return serviceMap, nil
}

// updateServiceCache 更新服务缓存
func updateServiceCache() {
	services, err := consulClient.Agent().Services()
	if err != nil {
		log.Printf("更新服务缓存失败: %v", err)
		return
	}
	cacheMutex.Lock()
	defer cacheMutex.Unlock()
	serviceMap := make(map[string][]*api.AgentService)
	for _, service := range services {
		serviceMap[service.Service] = append(serviceMap[service.Service], service)
	}
	serviceCache = serviceMap
	cacheExpiration = time.Now().Add(cacheDuration)
	log.Println("服务缓存已更新")
}
