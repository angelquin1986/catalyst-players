package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
)

// SeasonService handles business logic for season operations
type SeasonService struct {
	seasonRepo repositories.SeasonRepository
}

// NewSeasonService creates a new season service instance
func NewSeasonService(seasonRepo repositories.SeasonRepository) *SeasonService {
	return &SeasonService{
		seasonRepo: seasonRepo,
	}
}

// CreateSeason creates a new season
func (s *SeasonService) CreateSeason(season *entities.Season) error {
	if season.Name == "" {
		return errors.New("season name is required")
	}

	if season.LeagueID == 0 {
		return errors.New("league ID is required")
	}

	if season.StartsAt.IsZero() {
		return errors.New("start date is required")
	}

	if season.EndsAt.IsZero() {
		return errors.New("end date is required")
	}

	if season.StartsAt.After(season.EndsAt) {
		return errors.New("start date must be before end date")
	}

	return s.seasonRepo.Create(season)
}

// GetSeasonByID retrieves a season by ID
func (s *SeasonService) GetSeasonByID(id uint) (*entities.Season, error) {
	if id == 0 {
		return nil, errors.New("invalid season ID")
	}

	return s.seasonRepo.GetByID(id)
}

// GetSeasonWithLeague retrieves a season with league information
func (s *SeasonService) GetSeasonWithLeague(id uint) (*entities.Season, error) {
	if id == 0 {
		return nil, errors.New("invalid season ID")
	}

	return s.seasonRepo.GetWithLeague(id)
}

// GetSeasonWithTeams retrieves a season with its teams
func (s *SeasonService) GetSeasonWithTeams(id uint) (*entities.Season, error) {
	if id == 0 {
		return nil, errors.New("invalid season ID")
	}

	return s.seasonRepo.GetWithTeams(id)
}

// GetAllSeasons retrieves all seasons
func (s *SeasonService) GetAllSeasons() ([]entities.Season, error) {
	return s.seasonRepo.GetAll()
}

// GetAllSeasonsWithTeams retrieves all seasons with their teams
func (s *SeasonService) GetAllSeasonsWithTeams() ([]entities.Season, error) {
	return s.seasonRepo.GetAllWithTeams()
}

// GetActiveSeasons retrieves all active seasons
func (s *SeasonService) GetActiveSeasons() ([]entities.Season, error) {
	return s.seasonRepo.GetActiveSeasons()
}

// GetSeasonsByLeagueID retrieves all seasons for a league
func (s *SeasonService) GetSeasonsByLeagueID(leagueID uint) ([]entities.Season, error) {
	if leagueID == 0 {
		return nil, errors.New("invalid league ID")
	}

	return s.seasonRepo.GetByLeagueID(leagueID)
}

// UpdateSeason updates an existing season
func (s *SeasonService) UpdateSeason(season *entities.Season) error {
	if season.ID == 0 {
		return errors.New("invalid season ID")
	}

	if season.Name == "" {
		return errors.New("season name is required")
	}

	return s.seasonRepo.Update(season)
}

// DeleteSeason deletes a season by ID
func (s *SeasonService) DeleteSeason(id uint) error {
	if id == 0 {
		return errors.New("invalid season ID")
	}

	return s.seasonRepo.Delete(id)
}

// ActivateSeason activates a season
func (s *SeasonService) ActivateSeason(id uint) error {
	season, err := s.seasonRepo.GetByID(id)
	if err != nil {
		return err
	}

	season.Status = entities.SeasonStatusActive
	return s.seasonRepo.Update(season)
}

// CompleteSeason completes a season
func (s *SeasonService) CompleteSeason(id uint) error {
	season, err := s.seasonRepo.GetByID(id)
	if err != nil {
		return err
	}

	season.Status = entities.SeasonStatusCompleted
	return s.seasonRepo.Update(season)
}
