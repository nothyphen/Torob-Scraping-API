package main

import (
	"webscrap/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := routes.Urls()
	gin.SetMode(gin.DebugMode)
	err := r.Run()

	if err != nil {
		panic("cant start webserver !!!!!!!!!!11")
	}
}