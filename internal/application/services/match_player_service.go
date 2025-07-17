package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
)

// MatchPlayerService handles business logic for match player statistics operations
type MatchPlayerService struct {
	matchPlayerRepo repositories.MatchPlayerRepository
}

// NewMatchPlayerService creates a new match player service instance
func NewMatchPlayerService(matchPlayerRepo repositories.MatchPlayerRepository) *MatchPlayerService {
	return &MatchPlayerService{
		matchPlayerRepo: matchPlayerRepo,
	}
}

// CreateMatchPlayer creates a new match player statistic
func (s *MatchPlayerService) CreateMatchPlayer(matchPlayer *entities.MatchPlayer) error {
	if matchPlayer.MatchID == 0 {
		return errors.New("match ID is required")
	}
	
	if matchPlayer.TeamID == 0 {
		return errors.New("team ID is required")
	}
	
	if matchPlayer.PlayerID == 0 {
		return errors.New("player ID is required")
	}
	
	if matchPlayer.RedCard < 0 || matchPlayer.YellowCard < 0 || matchPlayer.Goals < 0 {
		return errors.New("statistics cannot be negative")
	}
	
	return s.matchPlayerRepo.Create(matchPlayer)
}

// GetMatchPlayerByID retrieves a match player statistic by ID
func (s *MatchPlayerService) GetMatchPlayerByID(id uint) (*entities.MatchPlayer, error) {
	if id == 0 {
		return nil, errors.New("invalid match player ID")
	}
	
	return s.matchPlayerRepo.GetByID(id)
}

// GetAllMatchPlayers retrieves all match player statistics
func (s *MatchPlayerService) GetAllMatchPlayers() ([]entities.MatchPlayer, error) {
	return s.matchPlayerRepo.GetAll()
}

// GetMatchPlayersByMatchID retrieves all player statistics for a match
func (s *MatchPlayerService) GetMatchPlayersByMatchID(matchID uint) ([]entities.MatchPlayer, error) {
	if matchID == 0 {
		return nil, errors.New("invalid match ID")
	}
	
	return s.matchPlayerRepo.GetByMatchID(matchID)
}

// GetMatchPlayersByPlayerID retrieves all match statistics for a player
func (s *MatchPlayerService) GetMatchPlayersByPlayerID(playerID uint) ([]entities.MatchPlayer, error) {
	if playerID == 0 {
		return nil, errors.New("invalid player ID")
	}
	
	return s.matchPlayerRepo.GetByPlayerID(playerID)
}

// GetMatchPlayersByTeamID retrieves all match statistics for a team
func (s *MatchPlayerService) GetMatchPlayersByTeamID(teamID uint) ([]entities.MatchPlayer, error) {
	if teamID == 0 {
		return nil, errors.New("invalid team ID")
	}
	
	return s.matchPlayerRepo.GetByTeamID(teamID)
}

// GetPlayerStats retrieves player statistics for a specific season
func (s *MatchPlayerService) GetPlayerStats(playerID uint, seasonID uint) ([]entities.MatchPlayer, error) {
	if playerID == 0 {
		return nil, errors.New("invalid player ID")
	}
	
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}
	
	return s.matchPlayerRepo.GetPlayerStats(playerID, seasonID)
}

// UpdateMatchPlayer updates an existing match player statistic
func (s *MatchPlayerService) UpdateMatchPlayer(matchPlayer *entities.MatchPlayer) error {
	if matchPlayer.ID == 0 {
		return errors.New("invalid match player ID")
	}
	
	if matchPlayer.RedCard < 0 || matchPlayer.YellowCard < 0 || matchPlayer.Goals < 0 {
		return errors.New("statistics cannot be negative")
	}
	
	return s.matchPlayerRepo.Update(matchPlayer)
}

// DeleteMatchPlayer deletes a match player statistic by ID
func (s *MatchPlayerService) DeleteMatchPlayer(id uint) error {
	if id == 0 {
		return errors.New("invalid match player ID")
	}
	
	return s.matchPlayerRepo.Delete(id)
} 