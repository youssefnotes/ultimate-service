package platform

import "time"

type DBCfg struct {
	SSLmode    string
	Timezone   string
	Scheme     string
	Username   string
	Password   string
	IP         string
	Path       string
	DriverName string
}
type WebCfg struct {
	APIHost         string
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

//TODO consolidate all config into one
type Config struct {
	DB  DBCfg
	Web WebCfg
}
