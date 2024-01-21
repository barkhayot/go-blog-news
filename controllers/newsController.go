package controllers

import (
	"database/sql"
	"errors"
	"log"
	"task/helper"
	"task/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary Retrieve all news
// @Produce json
// @Success 200 {array} models.News
// @Failure 500 {string} string "Internal Server Error"
// @Router /news [get]
func GetNewsController(c *gin.Context) {
	news, err := helper.GetNews()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching news"})
		return
	}
	c.JSON(200, news)
}

// @Summary Create a new news
// @Accept json
// @Param news body models.News true "News object to be created"
// @Success 201 {object} models.News
// @Failure 400 {string} string "Bad Request: Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /news [post]
func CreateNewsController(c *gin.Context) {

	var news models.News

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// if news.Title == "" || news.Body == "" {
	// 	c.JSON(400, gin.H{"error": "Title and Body are required fields"}) // Validate request body
	// 	return
	// }

	news.ID = uuid.New()
	err := helper.CreateNews(news)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating news"})
		return
	}

	c.JSON(201, news)
}

// @Summary Retrieve a news by ID
// @Param id path string true "News ID"
// @Produce json
// @Success 200 {object} models.News
// @Failure 400 {string} string "Bad Request: Invalid UUID format for newsID"
// @Failure 404 {string} string "News not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /news/{id} [get]
func GetNewsByIdController(c *gin.Context) {
	// Get id of post
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format for newsID"}) // validate for UUID
		return
	}

	news, err := helper.GetNewsByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error fetching news"})
		return
	}

	if news == nil {
		c.JSON(404, gin.H{"error": "News not found"})
		return
	}

	c.JSON(200, news)
}

// @Summary Update a news by ID
// @Param id path string true "News ID"
// @Accept json
// @Param news body models.News true "Updated news data"
// @Success 200 {object} models.News
// @Failure 400 {string} string "Bad Request: Invalid UUID format for newsID or Invalid request payload"
// @Failure 404 {string} string "News not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /news/{id} [put]
func UpdateNewsByIdController(c *gin.Context) {
	id := c.Param("id")
	var news models.News

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format for newsID"}) // validate for UUID
		return
	}

	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := helper.UpdateNewsByID(id, news)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "News not found" {
			c.JSON(404, gin.H{"error": "News not found"})
			return
		}

		log.Println(err)
		c.JSON(500, gin.H{"error": "Error updating news"})
		return
	}

	c.JSON(200, news)

}

// @Summary Delete a news by ID
// @Param id path string true "News ID"
// @Success 204 {string} string "News deleted successfully"
// @Failure 400 {string} string "Bad Request: Invalid UUID format for newsID"
// @Failure 404 {string} string "News not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /news/{id} [delete]
func DeleteNewsByIdController(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format for newsID"}) // validate for UUID
		return
	}

	err := helper.DeleteNewsByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting news"})
		return
	}

	c.JSON(200, gin.H{"message": "News deleted successfully"})

}
