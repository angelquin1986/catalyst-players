package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// LeagueRepositoryImpl implements the LeagueRepository interface using GORM
type LeagueRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewLeagueRepositoryImpl creates a new league repository implementation
func NewLeagueRepositoryImpl(db *gorm.DB) repositories.LeagueRepository {
	return &LeagueRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new league
func (r *LeagueRepositoryImpl) Create(league *entities.League) error {
	return r.db.Create(league).Error
}

// GetByID retrieves a league by ID
func (r *LeagueRepositoryImpl) GetByID(id uint) (*entities.League, error) {
	var league entities.League
	err := r.db.First(&league, id).Error
	if err != nil {
		return nil, err
	}
	return &league, nil
}

// GetAll retrieves all leagues
func (r *LeagueRepositoryImpl) GetAll() ([]entities.League, error) {
	var leagues []entities.League
	err := r.db.Find(&leagues).Error
	return leagues, err
}

// GetWithSeasons retrieves a league with its seasons
func (r *LeagueRepositoryImpl) GetWithSeasons(id uint) (*entities.League, error) {
	var league entities.League
	err := r.db.Preload("Seasons").First(&league, id).Error
	if err != nil {
		return nil, err
	}
	return &league, nil
}

// Update updates an existing league
func (r *LeagueRepositoryImpl) Update(league *entities.League) error {
	r.logger.Info("Updating league with ID: %d", league.ID)
	err := r.db.Model(league).Updates(league).Error
	if err != nil {
		r.logger.Error("Failed to update league with ID %d: %v", league.ID, err)
		return err
	}
	r.logger.Info("Successfully updated league with ID: %d", league.ID)
	return nil
}

// Delete deletes a league by ID
func (r *LeagueRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.League{}, id).Error
}
