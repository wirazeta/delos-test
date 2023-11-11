package main

import (
	"delos-test/controllers/farmcontroller"
	"delos-test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/api/farms", farmcontroller.Index)
	r.POST("/api/farms", farmcontroller.Create)
	r.PUT("/api/farms/:id", farmcontroller.Update)
	r.Run()
}
