package database

import (
	"database/sql"
	"errors"
)

type Dish struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"img"`
	Calorie     int     `json:"calorie"`
	Runtime     int     `json:"runtime"`
	Like        float64 `json:"like"`
	Dislike     float64 `json:"dislike"`
	Nationality string  `json:"nationality"`
}

type Logger interface {
	Printf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

var ErrDishNotFound = errors.New("dish not found")

func GetDishDetails(db *sql.DB, logger Logger, dishName string, nationality string) (*Dish, error) {
	dishQuery := `
		SELECT name, img, calorie, runtime 
		FROM Dishes 
		WHERE name = ?
	`
	
	var dish Dish
	err := db.QueryRow(dishQuery, dishName).Scan(
		&dish.Name,
		&dish.ImageURL,
		&dish.Calorie,
		&dish.Runtime,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Printf("Dish '%s' not found", dishName)
			return nil, ErrDishNotFound
		}
		logger.Printf("Error querying dish: %v", err)
		return nil, err
	}
	
	feedbackQuery := `
		SELECT like, dislike 
		FROM Feedbacks 
		WHERE dishName = ? AND nationality = ?
	`
	
	err = db.QueryRow(feedbackQuery, dishName, nationality).Scan(
		&dish.Like,
		&dish.Dislike,
	)
	
	if err != nil && err != sql.ErrNoRows {
		logger.Printf("Error querying feedback: %v", err)
		return nil, err
	}
	
	dish.Nationality = nationality
	logger.Printf("Successfully retrieved details for dish '%s' and nationality '%s'", dishName, nationality)
	return &dish, nil
}

func UpdateDishFeedback(db *sql.DB, logger Logger, dishName string, nationality string, feedback string) error {
	dishQuery := `SELECT COUNT(*) FROM Dishes WHERE name = ?`
	var count int
	err := db.QueryRow(dishQuery, dishName).Scan(&count)
	if err != nil {
		logger.Printf("Error checking if dish exists: %v", err)
		return err
	}
	
	if count == 0 {
		logger.Printf("Dish '%s' not found", dishName)
		return ErrDishNotFound
	}

	checkQuery := `SELECT COUNT(*) FROM Feedbacks WHERE dishName = ? AND nationality = ?`
	err = db.QueryRow(checkQuery, dishName, nationality).Scan(&count)
	if err != nil {
		logger.Printf("Error checking existing feedback: %v", err)
		return err
	}
	
	var queryStr string
	if count == 0 {
		if feedback == "like" {
			queryStr = `INSERT INTO Feedbacks (dishName, nationality, like, dislike) VALUES (?, ?, 1, 0)`
		} else {
			queryStr = `INSERT INTO Feedbacks (dishName, nationality, like, dislike) VALUES (?, ?, 0, 1)`
		}
		
		_, err = db.Exec(queryStr, dishName, nationality)
		if err != nil {
			logger.Printf("Error inserting new feedback: %v", err)
			return err
		}
	} else {
		if feedback == "like" {
			queryStr = `UPDATE Feedbacks SET like = like + 1 WHERE dishName = ? AND nationality = ?`
		} else {
			queryStr = `UPDATE Feedbacks SET dislike = dislike + 1 WHERE dishName = ? AND nationality = ?`
		}
		
		_, err = db.Exec(queryStr, dishName, nationality)
		if err != nil {
			logger.Printf("Error updating feedback: %v", err)
			return err
		}
	}
	
	logger.Printf("Successfully updated '%s' feedback for dish '%s' and nationality '%s'", 
		feedback, dishName, nationality)
	return nil
}