package config

import (
	"context"

	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	GlobalConfig *AppConfig
	DB           *gorm.DB
	Ctx          = context.Background()
	Rdb          *redis.Client
	Es           *elastic.Client
)
