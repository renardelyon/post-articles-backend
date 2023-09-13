package config

type Config struct {
	Application     ApplicationConfig
	ArticleDatabase DBConfig
}

type ApplicationConfig struct {
	ServiceName    string
	ServiceVersion string
	ServerPort     string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Protocol string
	Password string
	Database string
}
