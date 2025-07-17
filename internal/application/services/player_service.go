package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
	"time"
)

// PlayerService handles business logic for player operations
type PlayerService struct {
	playerRepo repositories.PlayerRepository
}

// NewPlayerService creates a new player service instance
func NewPlayerService(playerRepo repositories.PlayerRepository) *PlayerService {
	return &PlayerService{
		playerRepo: playerRepo,
	}
}

// CreatePlayer creates a new player
func (s *PlayerService) CreatePlayer(player *entities.Player) error {
	if player.Name == "" {
		return errors.New("player name is required")
	}
	
	if player.LastName == "" {
		return errors.New("player last name is required")
	}
	
	if player.TeamID == 0 {
		return errors.New("team ID is required")
	}
	
	if player.Number <= 0 {
		return errors.New("player number must be greater than 0")
	}
	
	// Validate birth date (player must be at least 5 years old)
	if player.BirthDate.After(time.Now().AddDate(-5, 0, 0)) {
		return errors.New("player must be at least 5 years old")
	}
	
	return s.playerRepo.Create(player)
}

// GetPlayerByID retrieves a player by ID
func (s *PlayerService) GetPlayerByID(id uint) (*entities.Player, error) {
	if id == 0 {
		return nil, errors.New("invalid player ID")
	}
	
	return s.playerRepo.GetByID(id)
}

// GetPlayerWithTeam retrieves a player with team information
func (s *PlayerService) GetPlayerWithTeam(id uint) (*entities.Player, error) {
	if id == 0 {
		return nil, errors.New("invalid player ID")
	}
	
	return s.playerRepo.GetWithTeam(id)
}

// GetAllPlayers retrieves all players
func (s *PlayerService) GetAllPlayers() ([]entities.Player, error) {
	return s.playerRepo.GetAll()
}

// GetPlayersByTeamID retrieves all players in a team
func (s *PlayerService) GetPlayersByTeamID(teamID uint) ([]entities.Player, error) {
	if teamID == 0 {
		return nil, errors.New("invalid team ID")
	}
	
	return s.playerRepo.GetByTeamID(teamID)
}

// UpdatePlayer updates an existing player
func (s *PlayerService) UpdatePlayer(player *entities.Player) error {
	if player.ID == 0 {
		return errors.New("invalid player ID")
	}
	
	if player.Name == "" {
		return errors.New("player name is required")
	}
	
	if player.LastName == "" {
		return errors.New("player last name is required")
	}
	
	return s.playerRepo.Update(player)
}

// DeletePlayer deletes a player by ID
func (s *PlayerService) DeletePlayer(id uint) error {
	if id == 0 {
		return errors.New("invalid player ID")
	}
	
	return s.playerRepo.Delete(id)
}

// GetTopScorers retrieves top scoring players for a season
func (s *PlayerService) GetTopScorers(seasonID uint, limit int) ([]entities.Player, error) {
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}
	
	if limit <= 0 {
		limit = 10 // Default limit
	}
	
	return s.playerRepo.GetTopScorers(seasonID, limit)
} 