package repositories

import (
	"catalyst-players/internal/domain/entities"
	"time"
)

// MatchRepository defines the interface for match data operations
type MatchRepository interface {
	Create(match *entities.Match) error
	GetByID(id uint) (*entities.Match, error)
	GetAll() ([]entities.Match, error)
	Update(match *entities.Match) error
	Delete(id uint) error
	GetWithDetails(id uint) (*entities.Match, error)
	GetBySeasonID(seasonID uint) ([]entities.Match, error)
	GetByStage(seasonID uint, stage entities.MatchStage) ([]entities.Match, error)
	GetByDateRange(startDate, endDate time.Time) ([]entities.Match, error)
	GetByTeamID(teamID uint, limit int) ([]entities.Match, error)
	GetUpcoming(limit int) ([]entities.Match, error)
	GetLive() ([]entities.Match, error)
	GetCompleted(seasonID uint) ([]entities.Match, error)
}
