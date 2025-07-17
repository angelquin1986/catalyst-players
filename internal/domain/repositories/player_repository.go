package repositories

import "catalyst-players/internal/domain/entities"

// PlayerRepository defines the interface for player data operations
type PlayerRepository interface {
	Create(player *entities.Player) error
	GetByID(id uint) (*entities.Player, error)
	GetAll() ([]entities.Player, error)
	Update(player *entities.Player) error
	Delete(id uint) error
	GetByTeamID(teamID uint) ([]entities.Player, error)
	GetWithTeam(id uint) (*entities.Player, error)
	GetWithTags(id uint) (*entities.Player, error)
	GetTopScorers(seasonID uint, limit int) ([]entities.Player, error)
} 