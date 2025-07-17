package repositories

import "catalyst-players/internal/domain/entities"

// TeamRepository defines the interface for team data operations
type TeamRepository interface {
	Create(team *entities.Team) error
	GetByID(id uint) (*entities.Team, error)
	GetAll() ([]entities.Team, error)
	Update(team *entities.Team) error
	Delete(id uint) error
	GetWithPlayers(id uint) (*entities.Team, error)
	GetWithTags(id uint) (*entities.Team, error)
	GetBySeasonID(seasonID uint) ([]entities.Team, error)
	GetStandings(seasonID uint) ([]entities.Team, error)
} 