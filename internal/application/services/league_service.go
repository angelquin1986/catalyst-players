package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
)

// LeagueService handles business logic for league operations
type LeagueService struct {
	leagueRepo repositories.LeagueRepository
}

// NewLeagueService creates a new league service instance
func NewLeagueService(leagueRepo repositories.LeagueRepository) *LeagueService {
	return &LeagueService{
		leagueRepo: leagueRepo,
	}
}

// CreateLeague creates a new league
func (s *LeagueService) CreateLeague(league *entities.League) error {
	if league.Name == "" {
		return errors.New("league name is required")
	}
	
	return s.leagueRepo.Create(league)
}

// GetLeagueByID retrieves a league by ID
func (s *LeagueService) GetLeagueByID(id uint) (*entities.League, error) {
	if id == 0 {
		return nil, errors.New("invalid league ID")
	}
	
	return s.leagueRepo.GetByID(id)
}

// GetLeagueWithSeasons retrieves a league with its seasons
func (s *LeagueService) GetLeagueWithSeasons(id uint) (*entities.League, error) {
	if id == 0 {
		return nil, errors.New("invalid league ID")
	}
	
	return s.leagueRepo.GetWithSeasons(id)
}

// GetAllLeagues retrieves all leagues
func (s *LeagueService) GetAllLeagues() ([]entities.League, error) {
	return s.leagueRepo.GetAll()
}

// UpdateLeague updates an existing league
func (s *LeagueService) UpdateLeague(league *entities.League) error {
	if league.ID == 0 {
		return errors.New("invalid league ID")
	}
	
	if league.Name == "" {
		return errors.New("league name is required")
	}
	
	return s.leagueRepo.Update(league)
}

// DeleteLeague deletes a league by ID
func (s *LeagueService) DeleteLeague(id uint) error {
	if id == 0 {
		return errors.New("invalid league ID")
	}
	
	return s.leagueRepo.Delete(id)
} 