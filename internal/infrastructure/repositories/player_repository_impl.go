package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// PlayerRepositoryImpl implements the PlayerRepository interface using GORM
type PlayerRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewPlayerRepositoryImpl creates a new player repository
func NewPlayerRepositoryImpl(db *gorm.DB) repositories.PlayerRepository {
	return &PlayerRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new player
func (r *PlayerRepositoryImpl) Create(player *entities.Player) error {
	return r.db.Create(player).Error
}

// GetByID retrieves a player by ID
func (r *PlayerRepositoryImpl) GetByID(id uint) (*entities.Player, error) {
	var player entities.Player
	err := r.db.First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// GetAll retrieves all players
func (r *PlayerRepositoryImpl) GetAll() ([]entities.Player, error) {
	var players []entities.Player
	err := r.db.Preload("Team").Find(&players).Error
	return players, err
}

// GetByTeamID retrieves all players in a team
func (r *PlayerRepositoryImpl) GetByTeamID(teamID uint) ([]entities.Player, error) {
	var players []entities.Player
	err := r.db.Where("team_id = ?", teamID).Find(&players).Error
	return players, err
}

// GetWithTeam retrieves a player with team information
func (r *PlayerRepositoryImpl) GetWithTeam(id uint) (*entities.Player, error) {
	var player entities.Player
	err := r.db.Preload("Team").First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// GetWithTags retrieves a player with tags
func (r *PlayerRepositoryImpl) GetWithTags(id uint) (*entities.Player, error) {
	var player entities.Player
	err := r.db.Preload("Tags").First(&player, id).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

// Update updates a player's information
func (r *PlayerRepositoryImpl) Update(player *entities.Player) error {
	r.logger.Info("Updating player with ID: %d", player.ID)
	err := r.db.Model(player).Updates(player).Error
	if err != nil {
		r.logger.Error("Failed to update player with ID %d: %v", player.ID, err)
		return err
	}
	r.logger.Info("Successfully updated player with ID: %d", player.ID)
	return nil
}

// Delete deletes a player by ID
func (r *PlayerRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Player{}, id).Error
}

// GetTopScorers retrieves top scoring players for a season
func (r *PlayerRepositoryImpl) GetTopScorers(seasonID uint, limit int) ([]entities.Player, error) {
	var players []entities.Player
	err := r.db.Joins("JOIN match_player ON match_player.player_id = player.id").
		Joins("JOIN `match` ON `match`.id = match_player.match_id").
		Where("`match`.season_id = ?", seasonID).
		Group("player.id").
		Order("SUM(match_player.goals) DESC").
		Limit(limit).
		Find(&players).Error
	return players, err
}
