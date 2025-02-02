package model

import "time"

type Recipe struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title"`
	MakingTime  string    `json:"making_time"`
	Serves      string    `json:"serves"`
	Ingredients string    `json:"ingredients"`
	Cost        uint      `json:"cost"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RecipeRequest struct {
	Title       string `json:"title"`
	MakingTime  string `json:"making_time"`
	Serves      string `json:"serves"`
	Ingredients string `json:"ingredients"`
	Cost        uint   `json:"cost"`
}
