package main

import (
	"github.com/dollarkillerx/baguette/engine"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	app := gin.New()
	gin.SetMode(gin.ReleaseMode)

	surf := app.Group("/surf")
	{
		suftEngine := engine.Suft{}
		surf.POST("/GET",suftEngine.Get)
	}
	phantom := app.Group("/phantom")
	{
		phantomEngine := engine.Phantom{}
		phantom.POST("/GET",phantomEngine.Get)
	}

	socks := os.Getenv("RUNSOCKS")
	if socks == "" {
		socks = "0.0.0.0:8080"
	}
	if err := app.Run(socks);err != nil {
		log.Fatalln(err)
	}
}
