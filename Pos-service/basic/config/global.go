package config

import "gorm.io/gorm"

var (
	GlobalConfig *AppConfig
	DB           *gorm.DB
)
