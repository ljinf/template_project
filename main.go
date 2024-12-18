package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project/api/router"
	"github.com/ljinf/template_project/common/enum"
	"github.com/ljinf/template_project/config"
)

func main() {
	if config.App.Env == enum.ModeProd {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()

	router.RegisterRoutes(g)

	g.Run(":8080")
}
