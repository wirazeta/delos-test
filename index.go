package main

import (
	"delos-test/controllers/farmcontroller"
	"delos-test/controllers/pondcontroller"
	"delos-test/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// Call Farm
	r.GET("/api/farms", farmcontroller.Index)
	r.GET("/api/farms/:id", farmcontroller.GetById)
	r.POST("/api/farms", farmcontroller.Create)
	r.PUT("/api/farms/:id", farmcontroller.Update)
	r.DELETE("/api/farms/:id", farmcontroller.Delete)
	// Call Pond
	r.GET("/api/ponds", pondcontroller.Index)
	r.GET("/api/ponds/:id", pondcontroller.GetById)
	r.POST("/api/ponds", pondcontroller.Create)
	r.PUT("/api/ponds/:id", pondcontroller.Update)
	r.DELETE("/api/ponds/:id", pondcontroller.Delete)

	r.Run()
}
