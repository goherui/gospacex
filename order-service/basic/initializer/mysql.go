package initializer

import (
	"fmt"
	"gospacex/order-service/basic/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MySQLInit() {
	var err error
	once := sync.Once{}
	conf := config.GlobalConfig.Mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	once.Do(func() {
		config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		fmt.Println("数据库连接成功")
	})
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := config.DB.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//sqlDB.SetConnMaxLifetime(time.Hour)
	//err = config.DB.AutoMigrate(&model.Order{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("迁移成功")
}
