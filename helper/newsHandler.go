package helper

import (
	"database/sql"
	"errors"
	"log"
	"task/models"
)

func GetNews() ([]models.News, error) {
	var news []models.News
	err := DB.Select(&news, "SELECT * FROM news")
	if err != nil {
		log.Println(err) // Log the error
		return nil, err
	}
	return news, nil
}

func CreateNews(news models.News) error {
	_, err := DB.Exec("INSERT INTO news (id, title, body) VALUES ($1, $2, $3)", news.ID, news.Title, news.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func GetNewsByID(newsID string) (*models.News, error) {
	var news models.News

	err := DB.Get(&news, "SELECT * FROM news WHERE id = $1", newsID)
	if err != nil {
		if err == sql.ErrNoRows {
			// News not found
			return nil, nil
		}

		log.Println(err)
		return nil, err
	}

	return &news, nil
}

func UpdateNewsByID(NewsId string, updatedNews models.News) error {
	result, err := DB.Exec("UPDATE blogs SET title = $1, body = $2 WHERE id = $3", updatedNews.Title, updatedNews.Body, NewsId)
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
		return errors.New("News not found")
	}

	return nil
}

func DeleteNewsByID(NewsId string) error {
	result, err := DB.Exec("DELETE FROM News WHERE id = $1", NewsId)
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
		return errors.New("News not found")
	}

	return nil
}
