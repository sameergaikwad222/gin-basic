package main

import (
	"gin-basic/config"
	"gin-basic/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := config.InitDB()
	routes.RegisterAllRoutes(router, db)
	router.Run(":3000")
}