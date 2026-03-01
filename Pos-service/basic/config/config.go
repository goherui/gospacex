package config

type AppConfig struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Consul struct {
		Host        string
		Port        int
		ServiceName string
		ServicePort int
		TTL         int
	}
}
