package pondcontroller

import (
	"delos-test/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var ponds []models.Pond
	var log models.Log

	if err := models.DB.Where("name = ?", "GET /api/ponds").First(&log).Error; err != nil {
		log := models.Log{Name: "GET /api/ponds", Count: 1, UserAgent: c.Request.UserAgent()}
		if err := models.DB.Create(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	} else {
		log.Count++
		fmt.Print(models.DB.Where("name = ?", "GET /api/ponds").First(&log))
		if err := models.DB.Save(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	}
	find := models.DB.Preload("Farm").Find(&ponds)

	if err := find.Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": &ponds})
}

func GetById(c *gin.Context) {
	var pond models.Pond
	var log models.Log

	if err := models.DB.Where("name = ?", "GET /api/ponds/:id").First(&log).Error; err != nil {
		log := models.Log{Name: "GET /api/ponds/:id", Count: 1, UserAgent: c.Request.UserAgent()}
		if err := models.DB.Create(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	} else {
		log.Count++
		if err := models.DB.Save(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	}
	find := models.DB.Where("id = ?", c.Param("id")).Preload("Farm").First(&pond)

	if err := find.Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": &pond})
}

func Create(c *gin.Context) {
	var input models.CreatePond
	var log models.Log

	if err := models.DB.Where("name = ?", "POST /api/ponds").First(&log).Error; err != nil {
		log := models.Log{Name: "POST /api/ponds", Count: 1, UserAgent: c.Request.UserAgent()}
		if err := models.DB.Create(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	} else {
		log.Count++
		if err := models.DB.Save(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pond := models.Pond{FarmID: input.FarmID, Name: input.Name}
	result := models.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&pond)
	if err := result.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Data Created", "result": &pond})
}

func Update(c *gin.Context) {
	var pond models.Pond
	var log models.Log

	if err := models.DB.Where("name = ?", "PUT /api/ponds/:id").First(&log).Error; err != nil {
		log := models.Log{Name: "PUT /api/ponds/:id", Count: 1, UserAgent: c.Request.UserAgent()}
		if err := models.DB.Create(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	} else {
		log.Count++
		if err := models.DB.Save(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	}
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pond).Error; err != nil {
		var input models.UpdatePond
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}
		pond := models.Pond{FarmID: input.FarmID, Name: input.Name, ID: uint(id)}
		result := models.DB.Create(&pond)
		if err := result.Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Data Created", "result": pond})
		return
	}
	var input models.UpdatePond
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	result := models.DB.Model(&pond).Updates(input)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Updated", "result": &pond})
}

func Delete(c *gin.Context) {
	var pond models.Pond
	var log models.Log

	if err := models.DB.Where("name = ?", "DELETE /api/ponds/:id").First(&log).Error; err != nil {
		log := models.Log{Name: "DELETE /api/ponds/:id", Count: 1, UserAgent: c.Request.UserAgent()}
		if err := models.DB.Create(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	} else {
		log.Count++
		if err := models.DB.Save(&log).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Print(gin.H{"result": &log})
	}
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pond).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Delete(&pond)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Deleted", "result": true})
}
