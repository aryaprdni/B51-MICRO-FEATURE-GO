package models

import "time"

type Paslon struct {
	ID            int       `gorm:"primary key;autoIncrement" json:"id"`
	Name          string    `json:"name"`
	NumberPaslon  string    `json:"numberPaslon"`
	VisionMission string    `json:"visionMission"`
	Image         string    `json:"image"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
