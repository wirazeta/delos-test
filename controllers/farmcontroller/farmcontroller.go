package farmcontroller

import (
	"delos-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var farms []models.Farm

	var find = models.DB.Find(&farms)

	if find == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	}

	c.JSON(http.StatusOK, gin.H{"farms": farms})
}

func Create(c *gin.Context) {
	var farms models.Farm

	if err := c.ShouldBindJSON(&farms); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := models.DB.Create(&farms)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Success", "result": result})
}

func Update(c *gin.Context) {
	var farms models.Farm

	if err := c.ShouldBindJSON(&farms); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	result := models.DB.Where("id = ?", c.Param("id")).Save(&farms)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	if models.DB.First(&farms, c.Param("id")) != nil {
		c.JSON(http.StatusCreated, gin.H{"message": "Created", "result": result})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Farm Updated"})
}
