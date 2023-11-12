package farmcontroller

import (
	"delos-test/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var farms []models.Farm
	var log models.Log

	if err := models.DB.Where("name = ?", "GET /api/farms").First(&log).Error; err != nil {
		log := models.Log{Name: "GET /api/farms", Count: 1, UserAgent: c.Request.UserAgent()}
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

	var find = models.DB.Preload("Pond").Find(&farms)

	if find == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"farms": farms})
}

func GetById(c *gin.Context) {
	var farm models.Farm
	var log models.Log

	if err := models.DB.Where("name = ?", "GET /api/farms/:id").First(&log).Error; err != nil {
		log := models.Log{Name: "GET /api/farms/:id", Count: 1, UserAgent: c.Request.UserAgent()}
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
	var find = models.DB.Where("id = ?", c.Param("id")).Preload("Pond").First(&farm)

	if find.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"farm": farm})

}

func Create(c *gin.Context) {
	var input models.CreateFarmInput
	var log models.Log

	if err := models.DB.Where("name = ?", "POST /api/farms").First(&log).Error; err != nil {
		log := models.Log{Name: "POST /api/farms", Count: 1, UserAgent: c.Request.UserAgent()}
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	farm := models.Farm{Name: input.Name}

	result := models.DB.Create(&farm)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Success", "result": &farm})
}

func Update(c *gin.Context) {
	var farms models.Farm
	var log models.Log

	if err := models.DB.Where("name = ?", "PUT /api/farms/:id").First(&log).Error; err != nil {
		log := models.Log{Name: "PUT /api/farms/:id", Count: 1, UserAgent: c.Request.UserAgent()}
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
	if err := models.DB.Where("id = ?", c.Param("id")).First(&farms).Error; err != nil {
		var input models.CreateFarmInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		farm := models.Farm{Name: input.Name}
		result := models.DB.Create(&farm)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "Data Created", "result": &farm})
		return
	}

	var input models.UpdateFarmInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}

	result := models.DB.Model(&farms).Updates(input)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Updated", "result": &farms})
}

func Delete(c *gin.Context) {
	var farm models.Farm
	var log models.Log

	if err := models.DB.Where("name = ?", "DELETE /api/farms/:id").First(&log).Error; err != nil {
		log := models.Log{Name: "DELETE /api/farms/:id", Count: 1, UserAgent: c.Request.UserAgent()}
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
	if err := models.DB.Where("id = ?", c.Param("id")).First(&farm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Delete(&farm)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Deleted", "result": true})
}
