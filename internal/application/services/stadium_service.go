package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
)

// StadiumService handles business logic for stadium operations
type StadiumService struct {
	stadiumRepo repositories.StadiumRepository
}

// NewStadiumService creates a new stadium service instance
func NewStadiumService(stadiumRepo repositories.StadiumRepository) *StadiumService {
	return &StadiumService{
		stadiumRepo: stadiumRepo,
	}
}

// CreateStadium creates a new stadium
func (s *StadiumService) CreateStadium(stadium *entities.Stadium) error {
	if stadium.Name == "" {
		return errors.New("stadium name is required")
	}
	
	return s.stadiumRepo.Create(stadium)
}

// GetStadiumByID retrieves a stadium by ID
func (s *StadiumService) GetStadiumByID(id uint) (*entities.Stadium, error) {
	if id == 0 {
		return nil, errors.New("invalid stadium ID")
	}
	
	return s.stadiumRepo.GetByID(id)
}

// GetAllStadiums retrieves all stadiums
func (s *StadiumService) GetAllStadiums() ([]entities.Stadium, error) {
	return s.stadiumRepo.GetAll()
}

// UpdateStadium updates an existing stadium
func (s *StadiumService) UpdateStadium(stadium *entities.Stadium) error {
	if stadium.ID == 0 {
		return errors.New("invalid stadium ID")
	}
	
	if stadium.Name == "" {
		return errors.New("stadium name is required")
	}
	
	return s.stadiumRepo.Update(stadium)
}

// DeleteStadium deletes a stadium by ID
func (s *StadiumService) DeleteStadium(id uint) error {
	if id == 0 {
		return errors.New("invalid stadium ID")
	}
	
	return s.stadiumRepo.Delete(id)
} 