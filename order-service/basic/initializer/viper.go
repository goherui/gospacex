package initializer

import (
	"fmt"
	"gospacex/order-service/basic/config"
	"path/filepath"

	"github.com/spf13/viper"
)

func ViperInit() {
	var err error
	configPath := filepath.Join(GetProjectRoot(), "config.yaml")
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config.GlobalConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("配置加载成功", config.GlobalConfig)
}
