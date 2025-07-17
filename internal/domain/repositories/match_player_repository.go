package repositories

import "catalyst-players/internal/domain/entities"

// MatchPlayerRepository defines the interface for match player statistics operations
type MatchPlayerRepository interface {
	Create(matchPlayer *entities.MatchPlayer) error
	GetByID(id uint) (*entities.MatchPlayer, error)
	GetAll() ([]entities.MatchPlayer, error)
	Update(matchPlayer *entities.MatchPlayer) error
	Delete(id uint) error
	GetByMatchID(matchID uint) ([]entities.MatchPlayer, error)
	GetByPlayerID(playerID uint) ([]entities.MatchPlayer, error)
	GetByTeamID(teamID uint) ([]entities.MatchPlayer, error)
	GetPlayerStats(playerID uint, seasonID uint) ([]entities.MatchPlayer, error)
} 