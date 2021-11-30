package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noernova/tai-eng-dictionaryAPI/models"
)

// GET /
func Index(c *gin.Context) {
	apiendpoints := []string{
		"/api/v1/apikey={apikey}eng/:text",
		"/api/v1/apikey={apikey}shn/:text",
	}

	c.JSON(http.StatusOK, gin.H{"version": "1.0.0", "message": "Welcome to shn_eng-dic API.", "available_endpoints": apiendpoints })
}

// GET /api/v1/eng/:text
func Eng_to_Shn(c *gin.Context) {
	var Word []models.Word
	var Antonym models.Antonym
	var Definition models.Definition

	var Translate []models.Translate

	text := c.Param("text")

	if err := models.DB.Where("english = ?", text).Find(&Word).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No word found"})
		return
	}
	
	// search anonyms by word id
	for _, word := range Word {
		models.DB.Where("word_id = ?", word.ID).First(&Antonym)
		models.DB.Where("word_id = ?", word.ID).First(&Definition)
		Translate = append(Translate, models.Translate{Word: word, Antonym: Antonym, Definition: Definition})
	}

	c.JSON(http.StatusOK, gin.H{"data": Translate})
}

// GET /api/v1/shn/:text
func Shn_to_Eng(c *gin.Context) {
	var Word []models.Word
	var Antonym models.Antonym
	var Definition models.Definition

	var Translate []models.Translate

	text := c.Param("text")

	if err := models.DB.Where("shan = ?", text).Find(&Word).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No word found"})
		return
	}
	
	// search anonyms by word id
	for _, word := range Word {
		models.DB.Where("word_id = ?", word.ID).First(&Antonym)
		models.DB.Where("word_id = ?", word.ID).First(&Definition)
		Translate = append(Translate, models.Translate{Word: word, Antonym: Antonym, Definition: Definition})
	}

	c.JSON(http.StatusOK, gin.H{"data": Translate})
}