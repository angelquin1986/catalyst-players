package repositories

import "catalyst-players/internal/domain/entities"

// TagRepository defines the interface for tag data operations
type TagRepository interface {
	Create(tag *entities.Tag) error
	GetByID(id uint) (*entities.Tag, error)
	GetAll() ([]entities.Tag, error)
	Update(tag *entities.Tag) error
	Delete(id uint) error
	GetByPlayerID(playerID uint) ([]entities.Tag, error)
	GetByTeamID(teamID uint) ([]entities.Tag, error)
} 