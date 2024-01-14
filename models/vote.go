package models

import "time"

type Vote struct {
	ID        int       `gorm:"primary_key;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PaslonID  int       `json:"paslonID"`
	UserID    uint      `json:"userID" gorm:"foreignKey:UserID"`
}
