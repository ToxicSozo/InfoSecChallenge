package config

type Config struct {
	ServerAddr string
	AssetsDir  string
	DBConnStr  string
}

func NewConfig() Config {
	return Config{
		ServerAddr: "0.0.0.0:8080",
		AssetsDir:  "./web/public/assets",
		DBConnStr:  "postgres://admin:1423@localhost:5432/users?sslmode=disable",
	}
}
