package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/wicaksana/go-simple-server/web/controller"
	"github.com/wicaksana/go-simple-server/web/middleware"
)

func main() {
	r := gin.New()
	r.Use(middleware.JSONLogMIddleware())
	r.Use(gin.Recovery())

	r.Static("/css", "./web/static/css")
	r.Static("/js", "./web/static/js")
	r.LoadHTMLGlob("./web/templates/**/*")

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Info().
			Msg(fmt.Sprintf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers))
	}

	controller.Router(r)

	r.Run()
}
