package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project/internal/router"
	"github.com/ljinf/template_project/pkg/config"
	"github.com/ljinf/template_project/pkg/enum"
	"github.com/ljinf/template_project/pkg/logger"
)

func main() {

	var envConf = flag.String("conf", "config/application.dev.yml", "config path, eg: -conf ./config/application.dev.yml")
	flag.Parse()

	config.InitConfig(*envConf)
	logger.InitLogger()

	if config.App.Env == enum.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()

	router.RegisterRoutes(g)

	g.Run(":8080")
}
