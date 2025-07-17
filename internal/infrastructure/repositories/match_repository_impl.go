package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"
	"time"

	"gorm.io/gorm"
)

// MatchRepositoryImpl implements the MatchRepository interface using GORM
type MatchRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewMatchRepositoryImpl creates a new match repository implementation
func NewMatchRepositoryImpl(db *gorm.DB) repositories.MatchRepository {
	return &MatchRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new match
func (r *MatchRepositoryImpl) Create(match *entities.Match) error {
	// First create the record
	err := r.db.Create(match).Error
	if err != nil {
		return err
	}
	// Then reload the match with all relationships
	return r.db.Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Season.League").
		Preload("Stadium").
		Preload("PlayerStats").
		First(match, match.ID).Error
}

// GetByID retrieves a match by ID
func (r *MatchRepositoryImpl) GetByID(id uint) (*entities.Match, error) {
	var match entities.Match
	err := r.db.Preload("HomeTeam").Preload("AwayTeam").First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

// GetAll retrieves all matches
func (r *MatchRepositoryImpl) GetAll() ([]entities.Match, error) {
	var matches []entities.Match
	err := r.db.Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Season").
		Preload("Stadium").
		Preload("PlayerStats").Find(&matches).Error
	return matches, err
}

// GetWithDetails retrieves a match with all related details
func (r *MatchRepositoryImpl) GetWithDetails(id uint) (*entities.Match, error) {
	var match entities.Match
	err := r.db.Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Season").
		Preload("Stadium").
		Preload("PlayerStats").
		First(&match, id).Error
	if err != nil {
		return nil, err
	}
	return &match, nil
}

// GetBySeasonID retrieves all matches in a season, with associated teams
func (r *MatchRepositoryImpl) GetBySeasonID(seasonID uint) ([]entities.Match, error) {
	r.logger.Info("Retrieving matches for season ID: %d", seasonID)
	var matches []entities.Match
	err := r.db.Preload("HomeTeam").Preload("AwayTeam").Where("season_id = ?", seasonID).Find(&matches).Error
	if err != nil {
		r.logger.Error("Failed to retrieve matches for season ID %d: %v", seasonID, err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved %d matches for season ID %d", len(matches), seasonID)
	return matches, err
}

// GetByStage retrieves matches by stage for a season
func (r *MatchRepositoryImpl) GetByStage(seasonID uint, stage entities.MatchStage) ([]entities.Match, error) {
	var matches []entities.Match
	err := r.db.Where("season_id = ? AND stage = ?", seasonID, stage).Find(&matches).Error
	return matches, err
}

// GetByDateRange retrieves matches within a date range
func (r *MatchRepositoryImpl) GetByDateRange(startDate, endDate time.Time) ([]entities.Match, error) {
	var matches []entities.Match
	err := r.db.Where("date BETWEEN ? AND ?", startDate, endDate).Find(&matches).Error
	return matches, err
}

// GetByTeamID retrieves all matches for a team
func (r *MatchRepositoryImpl) GetByTeamID(teamID uint, limit int) ([]entities.Match, error) {
	var matches []entities.Match
	query := r.db.Where("home_team_id = ? OR away_team_id = ?", teamID, teamID)
	if limit > 0 {
		query = query.Limit(limit)
	}
	err := query.Find(&matches).Error
	return matches, err
}

// GetUpcoming retrieves upcoming matches
func (r *MatchRepositoryImpl) GetUpcoming(limit int) ([]entities.Match, error) {
	var matches []entities.Match
	err := r.db.Where("date > ?", time.Now()).
		Order("date ASC").
		Limit(limit).
		Find(&matches).Error
	return matches, err
}

// GetCompleted retrieves completed matches for a season
func (r *MatchRepositoryImpl) GetCompleted(seasonID uint) ([]entities.Match, error) {
	var matches []entities.Match
	err := r.db.Preload("HomeTeam").Preload("AwayTeam").Where("season_id = ? AND status = ?", seasonID, entities.MatchStatusFinished).
		Find(&matches).Error
	return matches, err
}

// GetLive retrieves live matches
func (r *MatchRepositoryImpl) GetLive() ([]entities.Match, error) {
	var matches []entities.Match
	err := r.db.Where("status = ?", entities.MatchStatusInProgress).Find(&matches).Error
	return matches, err
}

// Update updates a match
func (r *MatchRepositoryImpl) Update(match *entities.Match) error {
	r.logger.Info("Updating match with ID: %d", match.ID)
	err := r.db.Model(match).Updates(match).Error
	if err != nil {
		r.logger.Error("Failed to update match with ID %d: %v", match.ID, err)
		return err
	}
	r.logger.Info("Successfully updated match with ID: %d", match.ID)
	return nil
}

// Delete deletes a match
func (r *MatchRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Match{}, id).Error
}
