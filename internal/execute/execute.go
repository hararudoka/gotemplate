package execute

import (
	"github.com/hararudoka/gotemplate/internal/config"
	"github.com/hararudoka/gotemplate/internal/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func Execute() {
	log := logrus.New()

	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("Environment loading fails")
	}

	//environment, err := env.LoadEnv(".env")
	//if err != nil {
	//	log.Fatal("Environment loading fails")
	//}

	//_, err = storage.Open(&environment)
	//if err != nil {
	//	log.Fatal("DB connection dead")
	//}

	h, err := handler.New()
	if err != nil {
		log.Fatal("Handler init error")
	}

	e := echo.New()
	e.Use(middleware.Logger())

	h.HelloHandler.Call(e.Group("/hello"))

	e.Logger.Fatal(e.Start(":"+conf.Port))
}
