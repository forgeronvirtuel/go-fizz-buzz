package main

import (
	"github.com/forgeronvirtuel/fizzbuzzrest/api"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	addr := "localhost:80"
	if len(os.Args) >= 2 {
		addr = os.Args[1]
	}

	router := setupRouter()
	if err := router.Run(addr); err != nil {
		log.Fatalln(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	counter := api.NewRequestCounter()
	router.GET("/", api.CreateFizzbuzzRoute(counter))
	router.GET("/statistics", api.CreateStatisticsRoute(counter))
	return router
}
