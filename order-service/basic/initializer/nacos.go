package initializer

import (
	"fmt"
	"gospacex/order-service/basic/config"
	"path/filepath"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
)

func NacosInit() {
	var err error
	configPath := filepath.Join(GetProjectRoot(), "config.yaml")
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var NacosConfig config.NacosConfig
	err = viper.UnmarshalKey("Nacos", &NacosConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("nacos配置加载成功", NacosConfig)
	// Nacos服务器地址
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: NacosConfig.Addr,
			Port:   uint64(NacosConfig.Port),
		},
	}
	// 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         NacosConfig.Namespace, // 如果不需要命名空间，可以留空
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	configContent, err := configClient.GetConfig(vo.ConfigParam{
		DataId: NacosConfig.DataID,
		Group:  NacosConfig.Group,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("从nacos读取到的内容\n", configContent)
	config.GlobalConfig = &config.AppConfig{}
	nacosViper := viper.New()
	nacosViper.SetConfigType("yaml")
	err = nacosViper.ReadConfig(strings.NewReader(configContent))
	if err != nil {
		fmt.Println("配置读取失败")
		return
	}
	err = viper.Unmarshal(config.GlobalConfig)
	if err != nil {
		fmt.Println("获取配置失败")
		return
	}
	fmt.Println("配置加载成功", config.GlobalConfig)
}
