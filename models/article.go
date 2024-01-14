package models

import "time"

type Article struct {
	ID          int       `gorm:"primary_key;autoIncrement" json:"id"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	UserID      int       `json:"userID" gorm:"column:userID"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
