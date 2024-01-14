package models

import "time"

type Partai struct {
	ID            uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name          string    `json:"name"`
	PartyLeader   string    `json:"partyLeader"`
	VisionMission string    `json:"visionMission"`
	Image         string    `json:"image"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	PaslonID      uint      `json:"paslon" gorm:"foreignKey:PaslonID"`
}
