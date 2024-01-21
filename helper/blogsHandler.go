package helper

import (
	"database/sql"
	"errors"
	"log"
	"task/models"
)

func GetBlogs() ([]models.Blog, error) {
	var blogs []models.Blog
	err := DB.Select(&blogs, "SELECT * FROM blogs")
	if err != nil {
		log.Println(err) // Log the error
		return nil, err
	}
	return blogs, nil
}

func CreateBlog(blog models.Blog) error {
	_, err := DB.Exec("INSERT INTO blogs (id, title, body) VALUES ($1, $2, $3)", blog.ID, blog.Title, blog.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetBlogByID(blogID string) (*models.Blog, error) {
	var blog models.Blog

	err := DB.Get(&blog, "SELECT * FROM blogs WHERE id = $1", blogID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Blog not found
			return nil, nil
		}

		log.Println(err)
		return nil, err
	}

	return &blog, nil
}

func UpdateBlogByID(BlogId string, updatedBlog models.Blog) error {
	result, err := DB.Exec("UPDATE blogs SET title = $1, body = $2 WHERE id = $3", updatedBlog.Title, updatedBlog.Body, BlogId)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Blog not found")
	}

	return nil
}

func DeleteBlogByID(BlogId string) error {
	result, err := DB.Exec("DELETE FROM blogs WHERE id = $1", BlogId)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Blog not found")
	}

	return nil
}
