package infrastructure

type Config struct {
	DBName   string
	Username string
	Password string
	Port     int
}

func NewConfig() *Config {
	c := &Config{
		// TODO: 環境変数に切り出す (docker-compose.ymlも同様)
		DBName:   "gin_gorm_docker_gqlgen_crud_database",
		Username: "user",
		Password: "password",
		Port:     3306,
	}

	return c
}
