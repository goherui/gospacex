package config

type AppConfig struct {
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Redis struct {
		Host     string
		Port     int
		Password string
		Database int
	}
	Es struct {
		Host string
	}
	Consul struct {
		Host        string
		Port        int
		ServiceName string
		ServicePort int
		TTL         int
	}
}
type NacosConfig struct {
	Addr      string
	Port      int
	Namespace string
	DataID    string
	Group     string
	Username  string
	Password  string
}
