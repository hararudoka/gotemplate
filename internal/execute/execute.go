package execute

import (
	"github.com/hararudoka/gotemplate/internal/config"
	"github.com/hararudoka/gotemplate/internal/delivery/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Execute() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		logger.Fatal("Environment loading fails", zap.Error(err))
	}

	//environment, err := env.LoadEnv(".env")
	//if err != nil {
	//	logger.Fatal("Environment loading fails", zap.Error(err))
	//}

	//_, err = storage.Open(&environment)
	//if err != nil {
	//	logger.Fatal("DB connection dead", zap.Error(err))
	//}

	h, err := handler.New()
	if err != nil {
		logger.Fatal("Handler init error", zap.Error(err))
	}

	e := echo.New()
	e.Use(middleware.Logger())

	h.HelloHandler.Call(e.Group("/hello"))

	e.Logger.Fatal(e.Start(":"+conf.Port))
}
