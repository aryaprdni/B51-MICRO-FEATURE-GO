package models

import "time"

type Article struct {
	ID          int       `gorm:"primary key;autoIncrement" json:"id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
