package services

import (
	"B51-MICRO-FEATURE-GO/models"

	"gorm.io/gorm"
)

type VoteService interface {
	GetVotes() ([]models.Vote, error)
	GetVoteById(ID int) (models.Vote, error)
}

type voteRepository struct {
	db *gorm.DB
}

func VoteRepository(db *gorm.DB) *voteRepository {
	return &voteRepository{db}
}

func (rv *voteRepository) GetVotes() ([]models.Vote, error) {
	var votes []models.Vote
	err := rv.db.Find(&votes).Error

	return votes, err
}

func (rv *voteRepository) GetVoteById(ID int) (models.Vote, error) {
	var vote models.Vote
	err := rv.db.First(&vote, ID).Error
	return vote, err
}
