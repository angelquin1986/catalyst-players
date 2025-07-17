package repositories

import "catalyst-players/internal/domain/entities"

// LeagueRepository defines the interface for league data operations
type LeagueRepository interface {
	Create(league *entities.League) error
	GetByID(id uint) (*entities.League, error)
	GetAll() ([]entities.League, error)
	Update(league *entities.League) error
	Delete(id uint) error
	GetWithSeasons(id uint) (*entities.League, error)
} 