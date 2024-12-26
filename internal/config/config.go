package config

type Config struct {
	ServerAddr string
	AssetsDir  string
	DBConnStr  string
}

func NewConfig() Config {
	return Config{
		ServerAddr: ":8080",
		AssetsDir:  "./web/public/assets",
		DBConnStr:  "postgres://nebc:1423@localhost:5432/my_app?sslmode=disable",
	}
}
