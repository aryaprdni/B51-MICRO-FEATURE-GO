package services

import (
	"B51-MICRO-FEATURE-GO/models"

	"gorm.io/gorm"
)

type PartaiService interface {
	GetPartais() ([]models.Partai, error)
	GetPartaiById(ID int) (models.Partai, error)
}

type partaiRepository struct {
	db *gorm.DB
}

func PartaiRepository(db *gorm.DB) *partaiRepository {
	return &partaiRepository{db}
}

func (rpt *partaiRepository) GetPartais() ([]models.Partai, error) {
	var partais []models.Partai
	err := rpt.db.Find(&partais).Error
	return partais, err
}

func (rpt *partaiRepository) GetPartaiById(ID int) (models.Partai, error) {
	var partai models.Partai
	err := rpt.db.First(&partai, ID).Error
	return partai, err
}
