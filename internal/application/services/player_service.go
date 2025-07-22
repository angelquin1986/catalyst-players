package services

import (
	"bytes"
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
	"image"
	"regexp"
	"strings"
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

// validateIdentityCard validates the identity card format
func (s *PlayerService) validateIdentityCard(id string) error {
	if id == "" {
		return errors.New("identity card is required")
	}

	// Remove spaces and dashes
	id = strings.ReplaceAll(id, " ", "")
	id = strings.ReplaceAll(id, "-", "")

	// Check length (typically 10-13 digits for Ecuadorian cedula)
	if len(id) < 10 || len(id) > 13 {
		return errors.New("identity card must be between 10 and 13 digits")
	}

	// Check if it contains only digits
	matched, _ := regexp.MatchString(`^\d+$`, id)
	if !matched {
		return errors.New("identity card must contain only digits")
	}

	return nil
}

// validatePhoto validates the photo data
func (s *PlayerService) validatePhoto(photoData []byte) error {
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

// CreatePlayer creates a new player
func (s *PlayerService) CreatePlayer(player *entities.Player) error {
	// Validate identity card
	if err := s.validateIdentityCard(player.ID); err != nil {
		return err
	}

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

	// Validate photo
	if err := s.validatePhoto(player.Photo); err != nil {
		return err
	}

	return s.playerRepo.Create(player)
}

// GetPlayerByID retrieves a player by ID (cedula)
func (s *PlayerService) GetPlayerByID(id string) (*entities.Player, error) {
	if id == "" {
		return nil, errors.New("invalid player ID")
	}

	return s.playerRepo.GetByID(id)
}

// GetPlayerWithTeam retrieves a player with team information
func (s *PlayerService) GetPlayerWithTeam(id string) (*entities.Player, error) {
	if id == "" {
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
	if player.ID == "" {
		return errors.New("invalid player ID")
	}

	// Validate identity card
	if err := s.validateIdentityCard(player.ID); err != nil {
		return err
	}

	if player.Name == "" {
		return errors.New("player name is required")
	}

	if player.LastName == "" {
		return errors.New("player last name is required")
	}

	// Validate photo
	if err := s.validatePhoto(player.Photo); err != nil {
		return err
	}

	return s.playerRepo.Update(player)
}

// DeletePlayer deletes a player by ID (cedula)
func (s *PlayerService) DeletePlayer(id string) error {
	if id == "" {
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
