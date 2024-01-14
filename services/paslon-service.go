package services

import (
	"B51-MICRO-FEATURE-GO/models"

	"gorm.io/gorm"
)

type PaslonService interface {
	GetPaslons() ([]models.Paslon, error)
	GetPaslonById(ID int) (models.Paslon, error)
}

type paslonRepository struct {
	db *gorm.DB
}

func PaslonRepository(db *gorm.DB) *paslonRepository {
	return &paslonRepository{db}
}

func (rp *paslonRepository) GetPaslons() ([]models.Paslon, error) {
	var paslons []models.Paslon
	err := rp.db.Preload("Koalisi").Find(&paslons).Error
	return paslons, err
}

func (rp *paslonRepository) GetPaslonById(ID int) (models.Paslon, error) {
	var paslon models.Paslon
	err := rp.db.Preload("Koalisi").Find(&paslon, ID).Error
	return paslon, err
}
