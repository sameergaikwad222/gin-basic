package main

import (
	"gin-basic/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterAllRoutes(router)

	router.Run(":3000")
}