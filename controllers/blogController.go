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

// @Summary Retrieve all blogs
// @Produce json
// @Success 200 {array} models.Blog
// @Failure 500 {string} string "Internal Server Error"
// @Router /blogs [get]
func GetBlogsController(c *gin.Context) {
	blogs, err := helper.GetBlogs()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching blogs"})
		return
	}
	c.JSON(200, blogs)
}

// @Summary Create a new blog
// @Accept json
// @Param blog body models.Blog true "Blog object to be created"
// @Success 201 {object} models.Blog
// @Failure 400 {string} string "Bad Reques"
// @Failure 500 {string} string "Internal Server Error"
// @Router /blogs [post]
func CreateBlogController(c *gin.Context) {

	var blog models.Blog

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// if blog.Title == "" || blog.Body == "" {
	// 	c.JSON(400, gin.H{"error": "Title and Body are required fields"}) // Validate request body
	// 	return
	// }

	blog.ID = uuid.New()
	err := helper.CreateBlog(blog)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating blog"})
		return
	}

	c.JSON(201, blog)
}

// @Summary Retrieve a blog by ID
// @Param id path string true "Blog ID"
// @Produce json
// @Success 200 {object} models.Blog
// @Failure 400 {string} string "Bad Reques"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /blogs/{id} [get]
func GetBlogByIdController(c *gin.Context) {
	// Get id of post
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format for blogID"}) // validate for UUID
		return
	}

	blog, err := helper.GetBlogByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "Error fetching blog"})
		return
	}

	if blog == nil {
		c.JSON(404, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(200, blog)
}

// @Summary Update a blog by ID
// @Param id path string true "Blog ID"
// @Accept json
// @Param blog body models.Blog true "Updated blog data"
// @Success 200 {object} models.Blog
// @Failure 400 {string} string "Bad Reques"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /blogs/{id} [put]
func UpdateBlogByIdController(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format for blogID"}) // validate for UUID
		return
	}

	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := helper.UpdateBlogByID(id, blog)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "Blog not found" {
			c.JSON(404, gin.H{"error": "Blog not found"})
			return
		}

		log.Println(err)
		c.JSON(500, gin.H{"error": "Error updating blog"})
		return
	}

	c.JSON(200, blog)

}

// @Summary Delete a blog by ID
// @Param id path string true "Blog ID"
// @Success 204 {string} string "Blog deleted successfully"
// @Failure 400 {string} string "Bad Reques"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /blogs/{id} [delete]
func DeleteBlogByIdController(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format for blogID"}) // validate for UUID
		return
	}

	err := helper.DeleteBlogByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting blog"})
		return
	}

	c.JSON(200, gin.H{"message": "Blog deleted successfully"})

}
