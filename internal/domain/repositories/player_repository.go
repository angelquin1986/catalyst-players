package repositories

import "catalyst-players/internal/domain/entities"

// PlayerRepository defines the interface for player data operations
type PlayerRepository interface {
	Create(player *entities.Player) error
	GetByID(id string) (*entities.Player, error)
	GetAll() ([]entities.Player, error)
	Update(player *entities.Player) error
	Delete(id string) error
	GetByTeamID(teamID uint) ([]entities.Player, error)
	GetWithTeam(id string) (*entities.Player, error)
	GetWithTags(id string) (*entities.Player, error)
	GetTopScorers(seasonID uint, limit int) ([]entities.Player, error)
}
