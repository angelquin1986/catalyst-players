package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// TeamRepositoryImpl implements the TeamRepository interface
type TeamRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewTeamRepositoryImpl creates a new team repository implementation
func NewTeamRepositoryImpl(db *gorm.DB) repositories.TeamRepository {
	return &TeamRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new team
func (r *TeamRepositoryImpl) Create(team *entities.Team) error {
	return r.db.Create(team).Error
}

// GetByID retrieves a team by ID
func (r *TeamRepositoryImpl) GetByID(id uint) (*entities.Team, error) {
	var team entities.Team
	err := r.db.First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// GetAll retrieves all teams
func (r *TeamRepositoryImpl) GetAll() ([]entities.Team, error) {
	var teams []entities.Team
	err := r.db.Find(&teams).Error
	return teams, err
}

// GetWithPlayers retrieves a team with its players
func (r *TeamRepositoryImpl) GetWithPlayers(id uint) (*entities.Team, error) {
	var team entities.Team
	err := r.db.Preload("Players").First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// GetWithTags retrieves a team with its tags
func (r *TeamRepositoryImpl) GetWithTags(id uint) (*entities.Team, error) {
	var team entities.Team
	err := r.db.Preload("Tags").First(&team, id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// Update updates a team
func (r *TeamRepositoryImpl) Update(team *entities.Team) error {
	r.logger.Info("Updating team with ID: %d", team.ID)
	err := r.db.Model(team).Updates(team).Error
	if err != nil {
		r.logger.Error("Failed to update team with ID %d: %v", team.ID, err)
		return err
	}
	r.logger.Info("Successfully updated team with ID: %d", team.ID)
	return nil
}

// Delete deletes a team by ID
func (r *TeamRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Team{}, id).Error
}

// GetBySeasonID retrieves all teams in a season
func (r *TeamRepositoryImpl) GetBySeasonID(seasonID uint) ([]entities.Team, error) {
	var teams []entities.Team
	err := r.db.Joins("JOIN season_team ON season_team.team_id = team.id").
		Where("season_team.season_id = ?", seasonID).
		Find(&teams).Error
	return teams, err
}

// GetStandings retrieves team standings for a season
func (r *TeamRepositoryImpl) GetStandings(seasonID uint) ([]entities.Team, error) {
	var teams []entities.Team
	err := r.db.Joins("JOIN season_team ON season_team.team_id = team.id").
		Where("season_team.season_id = ?", seasonID).
		Order("team.id").
		Find(&teams).Error
	return teams, err
}
