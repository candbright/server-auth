package main

import (
	"github.com/candbright/go-core/rest/handler"
	"github.com/candbright/go-log/log"
	"github.com/candbright/go-log/options"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"piano-server/config"
	"piano-server/server/service"
	"strconv"
)

func main() {
	handler.AppName(config.Get("application.name"))
	err := log.Init(
		options.Path(config.Get("log.path")),
		options.Level(func() logrus.Level {
			level, err := logrus.ParseLevel(config.Get("log.level"))
			if err != nil {
				return logrus.InfoLevel
			}
			return level
		}),
		options.Format(&logrus.TextFormatter{}),
	)
	if err != nil {
		panic(err)
	}
	engine := gin.New()
	service.RegisterHandlers(engine)
	log.Debug("start application " + config.Get("application.name"))
	_ = engine.Run(":" + strconv.Itoa(config.GetInt("application.port")))
}
