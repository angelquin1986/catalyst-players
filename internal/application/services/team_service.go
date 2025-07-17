package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
)

// TeamService handles business logic for team operations
type TeamService struct {
	teamRepo repositories.TeamRepository
}

// NewTeamService creates a new team service instance
func NewTeamService(teamRepo repositories.TeamRepository) *TeamService {
	return &TeamService{
		teamRepo: teamRepo,
	}
}

// CreateTeam creates a new team
func (s *TeamService) CreateTeam(team *entities.Team) error {
	if team.Name == "" {
		return errors.New("team name is required")
	}
	
	if team.Category == "" {
		return errors.New("team category is required")
	}
	
	return s.teamRepo.Create(team)
}

// GetTeamByID retrieves a team by ID
func (s *TeamService) GetTeamByID(id uint) (*entities.Team, error) {
	if id == 0 {
		return nil, errors.New("invalid team ID")
	}
	
	return s.teamRepo.GetByID(id)
}

// GetTeamWithPlayers retrieves a team with its players
func (s *TeamService) GetTeamWithPlayers(id uint) (*entities.Team, error) {
	if id == 0 {
		return nil, errors.New("invalid team ID")
	}
	
	return s.teamRepo.GetWithPlayers(id)
}

// GetAllTeams retrieves all teams
func (s *TeamService) GetAllTeams() ([]entities.Team, error) {
	return s.teamRepo.GetAll()
}

// UpdateTeam updates an existing team
func (s *TeamService) UpdateTeam(team *entities.Team) error {
	if team.ID == 0 {
		return errors.New("invalid team ID")
	}
	
	if team.Name == "" {
		return errors.New("team name is required")
	}
	
	return s.teamRepo.Update(team)
}

// DeleteTeam deletes a team by ID
func (s *TeamService) DeleteTeam(id uint) error {
	if id == 0 {
		return errors.New("invalid team ID")
	}
	
	return s.teamRepo.Delete(id)
}

// GetTeamsBySeasonID retrieves all teams in a season
func (s *TeamService) GetTeamsBySeasonID(seasonID uint) ([]entities.Team, error) {
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}
	
	return s.teamRepo.GetBySeasonID(seasonID)
}

// GetTeamStandings retrieves team standings for a season
func (s *TeamService) GetTeamStandings(seasonID uint) ([]entities.Team, error) {
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}
	
	return s.teamRepo.GetStandings(seasonID)
} 