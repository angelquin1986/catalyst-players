package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// SeasonRepositoryImpl implements the SeasonRepository interface using GORM
type SeasonRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewSeasonRepositoryImpl creates a new season repository implementation
func NewSeasonRepositoryImpl(db *gorm.DB) repositories.SeasonRepository {
	return &SeasonRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new season
func (r *SeasonRepositoryImpl) Create(season *entities.Season) error {
	r.logger.Info("Creating new season with ID: %d", season.ID)
	err := r.db.Create(season).Error
	if err != nil {
		r.logger.Error("Failed to create season: %v", err)
		return err
	}
	r.logger.Info("Successfully created season with ID: %d", season.ID)
	return nil
}

// GetByID retrieves a season by ID
func (r *SeasonRepositoryImpl) GetByID(id uint) (*entities.Season, error) {
	r.logger.Info("Retrieving season by ID: %d", id)
	var season entities.Season
	err := r.db.First(&season, id).Error
	if err != nil {
		r.logger.Error("Failed to retrieve season with ID %d: %v", id, err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved season with ID: %d", id)
	return &season, nil
}

// GetAll retrieves all seasons
func (r *SeasonRepositoryImpl) GetAll() ([]entities.Season, error) {
	r.logger.Info("Retrieving all seasons")
	var seasons []entities.Season
	err := r.db.Preload("Teams").Find(&seasons).Error
	if err != nil {
		r.logger.Error("Failed to retrieve all seasons: %v", err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved %d seasons", len(seasons))
	return seasons, err
}

// GetAllWithTeams retrieves all seasons with their associated teams
func (r *SeasonRepositoryImpl) GetAllWithTeams() ([]entities.Season, error) {
	r.logger.Info("Retrieving all seasons with teams")
	var seasons []entities.Season
	err := r.db.Preload("Teams").Find(&seasons).Error
	if err != nil {
		r.logger.Error("Failed to retrieve all seasons with teams: %v", err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved %d seasons with teams", len(seasons))
	return seasons, nil
}

// GetWithLeague retrieves a season with league information
func (r *SeasonRepositoryImpl) GetWithLeague(id uint) (*entities.Season, error) {
	r.logger.Info("Retrieving season with league information for ID: %d", id)
	var season entities.Season
	err := r.db.Preload("League").First(&season, id).Error
	if err != nil {
		r.logger.Error("Failed to retrieve season with league for ID %d: %v", id, err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved season with league for ID: %d", id)
	return &season, nil
}

// GetWithTeams retrieves a season with its teams
func (r *SeasonRepositoryImpl) GetWithTeams(id uint) (*entities.Season, error) {
	r.logger.Info("Retrieving season with teams for ID: %d", id)
	var season entities.Season
	err := r.db.Preload("Teams").First(&season, id).Error
	if err != nil {
		r.logger.Error("Failed to retrieve season with teams for ID %d: %v", id, err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved season with teams for ID: %d", id)
	return &season, nil
}

// GetWithMatches retrieves a season with its matches
func (r *SeasonRepositoryImpl) GetWithMatches(id uint) (*entities.Season, error) {
	r.logger.Info("Retrieving season with matches for ID: %d", id)
	var season entities.Season
	err := r.db.Preload("Matches").First(&season, id).Error
	if err != nil {
		r.logger.Error("Failed to retrieve season with matches for ID %d: %v", id, err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved season with matches for ID: %d", id)
	return &season, nil
}

// GetActiveSeasons retrieves all active seasons
func (r *SeasonRepositoryImpl) GetActiveSeasons() ([]entities.Season, error) {
	r.logger.Info("Retrieving all active seasons")
	var seasons []entities.Season
	err := r.db.Where("status = ?", entities.SeasonStatusActive).Find(&seasons).Error
	if err != nil {
		r.logger.Error("Failed to retrieve active seasons: %v", err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved %d active seasons", len(seasons))
	return seasons, err
}

// GetByLeagueID retrieves all seasons for a league
func (r *SeasonRepositoryImpl) GetByLeagueID(leagueID uint) ([]entities.Season, error) {
	r.logger.Info("Retrieving seasons for league ID: %d", leagueID)
	var seasons []entities.Season
	err := r.db.Where("league_id = ?", leagueID).Find(&seasons).Error
	if err != nil {
		r.logger.Error("Failed to retrieve seasons for league ID %d: %v", leagueID, err)
		return nil, err
	}
	r.logger.Info("Successfully retrieved %d seasons for league ID: %d", len(seasons), leagueID)
	return seasons, err
}

// Update updates an existing season
func (r *SeasonRepositoryImpl) Update(season *entities.Season) error {
	r.logger.Info("Updating season with ID: %d", season.ID)
	err := r.db.Model(season).Updates(season).Error
	if err != nil {
		r.logger.Error("Failed to update season with ID %d: %v", season.ID, err)
		return err
	}
	r.logger.Info("Successfully updated season with ID: %d", season.ID)
	return nil
}

// Delete deletes a season by ID
func (r *SeasonRepositoryImpl) Delete(id uint) error {
	r.logger.Info("Deleting season with ID: %d", id)
	err := r.db.Delete(&entities.Season{}, id).Error
	if err != nil {
		r.logger.Error("Failed to delete season with ID %d: %v", id, err)
		return err
	}
	r.logger.Info("Successfully deleted season with ID: %d", id)
	return nil
}
