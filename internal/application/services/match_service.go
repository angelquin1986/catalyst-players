package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
	"time"
)

// MatchService handles business logic for match operations
type MatchService struct {
	matchRepo repositories.MatchRepository
}

// NewMatchService creates a new match service instance
func NewMatchService(matchRepo repositories.MatchRepository) *MatchService {
	return &MatchService{
		matchRepo: matchRepo,
	}
}

// CreateMatch creates a new match
func (s *MatchService) CreateMatch(match *entities.Match) error {
	if match.HomeTeamID == 0 {
		return errors.New("home team ID is required")
	}

	if match.AwayTeamID == 0 {
		return errors.New("away team ID is required")
	}

	if match.HomeTeamID == match.AwayTeamID {
		return errors.New("home team and away team cannot be the same")
	}

	if match.SeasonID == 0 {
		return errors.New("season ID is required")
	}

	if match.StadiumID == 0 {
		return errors.New("stadium ID is required")
	}

	if match.Date.IsZero() {
		return errors.New("match date is required")
	}

	// Validate hour field if provided
	if match.Hour != nil {
		if *match.Hour < 1 || *match.Hour > 24 {
			return errors.New("hour must be between 1 and 24")
		}
	}

	/*if match.Date.Before(time.Now()) {
		return errors.New("match date cannot be in the past")
	}*/

	return s.matchRepo.Create(match)
}

// GetMatchByID retrieves a match by ID
func (s *MatchService) GetMatchByID(id uint) (*entities.Match, error) {
	if id == 0 {
		return nil, errors.New("invalid match ID")
	}

	return s.matchRepo.GetByID(id)
}

// GetMatchWithDetails retrieves a match with all related details
func (s *MatchService) GetMatchWithDetails(id uint) (*entities.Match, error) {
	if id == 0 {
		return nil, errors.New("invalid match ID")
	}

	return s.matchRepo.GetWithDetails(id)
}

// GetAllMatches retrieves all matches
func (s *MatchService) GetAllMatches() ([]entities.Match, error) {
	return s.matchRepo.GetAll()
}

// GetMatchesBySeasonID retrieves all matches in a season
func (s *MatchService) GetMatchesBySeasonID(seasonID uint) ([]entities.Match, error) {
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}

	return s.matchRepo.GetBySeasonID(seasonID)
}

// GetMatchesByStage retrieves matches by stage for a season
func (s *MatchService) GetMatchesByStage(seasonID uint, stage entities.MatchStage) ([]entities.Match, error) {
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}

	return s.matchRepo.GetByStage(seasonID, stage)
}

// GetMatchesByDateRange retrieves matches within a date range
func (s *MatchService) GetMatchesByDateRange(startDate, endDate time.Time) ([]entities.Match, error) {
	if startDate.IsZero() || endDate.IsZero() {
		return nil, errors.New("start date and end date are required")
	}

	if startDate.After(endDate) {
		return nil, errors.New("start date must be before end date")
	}

	return s.matchRepo.GetByDateRange(startDate, endDate)
}

// GetMatchesByTeamID retrieves all matches for a team
func (s *MatchService) GetMatchesByTeamID(teamID uint) ([]entities.Match, error) {
	if teamID == 0 {
		return nil, errors.New("invalid team ID")
	}

	return s.matchRepo.GetByTeamID(teamID, 0) // Pass 0 for no limit
}

// GetUpcoming retrieves upcoming matches
func (s *MatchService) GetUpcoming(limit int) ([]entities.Match, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}

	return s.matchRepo.GetUpcoming(limit)
}

// GetCompleted retrieves completed matches for a season
func (s *MatchService) GetCompleted(seasonID uint) ([]entities.Match, error) {
	if seasonID == 0 {
		return nil, errors.New("invalid season ID")
	}

	return s.matchRepo.GetCompleted(seasonID)
}

// UpdateMatch updates an existing match
func (s *MatchService) UpdateMatch(match *entities.Match) error {
	if match.ID == 0 {
		return errors.New("invalid match ID")
	}

	if match.HomeTeamID == 0 {
		return errors.New("home team ID is required")
	}

	if match.AwayTeamID == 0 {
		return errors.New("away team ID is required")
	}

	if match.HomeTeamID == match.AwayTeamID {
		return errors.New("home team and away team cannot be the same")
	}

	// Validate hour field if provided
	if match.Hour != nil {
		if *match.Hour < 1 || *match.Hour > 24 {
			return errors.New("hour must be between 1 and 24")
		}
	}

	return s.matchRepo.Update(match)
}

// UpdateMatchScore updates the score of a match
func (s *MatchService) UpdateMatchScore(matchID uint, homeScore, awayScore int) error {
	match, err := s.matchRepo.GetByID(matchID)
	if err != nil {
		return err
	}

	if homeScore < 0 || awayScore < 0 {
		return errors.New("scores cannot be negative")
	}

	match.HomeTeamScore = &homeScore
	match.AwayTeamScore = &awayScore

	// Calculate points based on result
	if homeScore > awayScore {
		homePoints := 3
		awayPoints := 0
		match.HomeTeamPoints = &homePoints
		match.AwayTeamPoints = &awayPoints
	} else if homeScore < awayScore {
		homePoints := 0
		awayPoints := 3
		match.HomeTeamPoints = &homePoints
		match.AwayTeamPoints = &awayPoints
	} else {
		homePoints := 1
		awayPoints := 1
		match.HomeTeamPoints = &homePoints
		match.AwayTeamPoints = &awayPoints
	}

	return s.matchRepo.Update(match)
}

// DeleteMatch deletes a match by ID
func (s *MatchService) DeleteMatch(id uint) error {
	if id == 0 {
		return errors.New("invalid match ID")
	}

	return s.matchRepo.Delete(id)
}
