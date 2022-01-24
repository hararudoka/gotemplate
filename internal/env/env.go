package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	// Env fields e.g. password or hostname
	Password string `env:"PASSWORD"` // not sure about this
}

func LoadEnv(name string) (env Env, err error) { // .env
	err = godotenv.Load("../../configs/"+name) //?

	env.Password = os.Getenv("PASSWORD")
	//...

	return
}
