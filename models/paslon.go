package models

import "time"

type Paslon struct {
	ID            uint      `gorm:"primary_key;autoIncrement" json:"id"`
	Name          string    `json:"name"`
	NumberPaslon  int       `json:"numberPaslon"`
	VisionMission string    `json:"visionMission"`
	Image         string    `json:"image"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	Koalisi       []Partai  `gorm:"foreignKey:PaslonID" json:"koalisi"`
	Vote          uint      `json:"vote"`
}
