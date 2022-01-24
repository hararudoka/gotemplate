package env

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	// Env fields e.g. password or hostname
	Password string
	Username string
	DBName   string
	Hostname string
	Mode     string
}

func LoadEnv(path string) (env Env, err error) { // .env
	err = godotenv.Load(path) //?

	env.Password = os.Getenv("PASSWORD")
	env.Username = os.Getenv("USERNAME")
	env.DBName = os.Getenv("DBNAME")
	env.Hostname = os.Getenv("HOSTNAME")
	env.Mode = os.Getenv("MODE")

	return
}
