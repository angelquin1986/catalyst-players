package repositories

import "catalyst-players/internal/domain/entities"

// SeasonRepository defines the interface for season data operations
type SeasonRepository interface {
	Create(season *entities.Season) error
	GetByID(id uint) (*entities.Season, error)
	GetAll() ([]entities.Season, error)
	GetAllWithTeams() ([]entities.Season, error)
	Update(season *entities.Season) error
	Delete(id uint) error
	GetWithLeague(id uint) (*entities.Season, error)
	GetWithTeams(id uint) (*entities.Season, error)
	GetWithMatches(id uint) (*entities.Season, error)
	GetActiveSeasons() ([]entities.Season, error)
	GetByLeagueID(leagueID uint) ([]entities.Season, error)
}
