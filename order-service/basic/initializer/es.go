package initializer

import (
	"fmt"
	"gospacex/order-service/basic/config"

	"github.com/olivere/elastic/v7"
)

func EsInit() {
	var err error
	config.Es, err = elastic.NewClient(elastic.SetURL(config.GlobalConfig.Es.Host),
		elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("es连接成功")
}
