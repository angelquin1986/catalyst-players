package repositories

import "catalyst-players/internal/domain/entities"

// StadiumRepository defines the interface for stadium data operations
type StadiumRepository interface {
	Create(stadium *entities.Stadium) error
	GetByID(id uint) (*entities.Stadium, error)
	GetAll() ([]entities.Stadium, error)
	Update(stadium *entities.Stadium) error
	Delete(id uint) error
} 