package models

import "time"

type User struct {
	ID        uint      `gorm:"primary_key;autoIncrement" json:"id"`
	Fullname  string    `json:"fullname"`
	Address   string    `json:"address"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Articles  []Article `json:"articles"`
	Vote      Vote      `json:"vote" gorm:"foreignKey:UserID"`
}
