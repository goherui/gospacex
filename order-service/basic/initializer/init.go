package initializer

import (
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	NacosInit()
	ViperInit()
	MySQLInit()
	RedisInit()
	EsInit()
	err := ConsulInit()
	if err != nil {
		return
	}
}
func GetProjectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		dir := filepath.Dir(filename)
		for i := 0; i < 3; i++ {
			dir = filepath.Dir(dir)
		}
		return dir
	}
	wd, _ := os.Getwd()
	return wd
}
