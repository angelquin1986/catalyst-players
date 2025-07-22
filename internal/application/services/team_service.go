package services

import (
	"bytes"
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
	"image"
	"strings"
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

// validatePhoto validates the photo data
func (s *TeamService) validatePhoto(photoData []byte) error {
	if len(photoData) == 0 {
		return nil // Photo is optional
	}

	// Check file size (500KB limit)
	const maxSize = 500 * 1024 // 500KB
	if len(photoData) > maxSize {
		return errors.New("photo size exceeds 500KB limit")
	}

	// Check if it's a valid image (JPG or PNG)
	img, format, err := image.Decode(bytes.NewReader(photoData))
	if err != nil {
		return errors.New("invalid image format")
	}

	// Check format
	format = strings.ToLower(format)
	if format != "jpeg" && format != "png" {
		return errors.New("only JPG and PNG formats are allowed")
	}

	// Validate image dimensions (optional check)
	if img.Bounds().Dx() < 50 || img.Bounds().Dy() < 50 {
		return errors.New("image dimensions too small (minimum 50x50 pixels)")
	}

	return nil
}

// CreateTeam creates a new team
func (s *TeamService) CreateTeam(team *entities.Team) error {
	if team.Name == "" {
		return errors.New("team name is required")
	}

	if team.Category == "" {
		return errors.New("team category is required")
	}

	// Validate photo
	if err := s.validatePhoto(team.Photo); err != nil {
		return err
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

	// Validate photo
	if err := s.validatePhoto(team.Photo); err != nil {
		return err
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
