package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// MatchPlayerRepositoryImpl implements the MatchPlayerRepository interface
type MatchPlayerRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewMatchPlayerRepositoryImpl creates a new match player repository implementation
func NewMatchPlayerRepositoryImpl(db *gorm.DB) repositories.MatchPlayerRepository {
	return &MatchPlayerRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new match player record
func (r *MatchPlayerRepositoryImpl) Create(matchPlayer *entities.MatchPlayer) error {
	return r.db.Create(matchPlayer).Error
}

// GetByID retrieves a match player statistic by ID
func (r *MatchPlayerRepositoryImpl) GetByID(id uint) (*entities.MatchPlayer, error) {
	var matchPlayer entities.MatchPlayer
	err := r.db.First(&matchPlayer, id).Error
	if err != nil {
		return nil, err
	}
	return &matchPlayer, nil
}

// GetAll retrieves all match player statistics
func (r *MatchPlayerRepositoryImpl) GetAll() ([]entities.MatchPlayer, error) {
	var matchPlayers []entities.MatchPlayer
	err := r.db.Find(&matchPlayers).Error
	return matchPlayers, err
}

// GetByMatchID retrieves all player statistics for a match
func (r *MatchPlayerRepositoryImpl) GetByMatchID(matchID uint) ([]entities.MatchPlayer, error) {
	var matchPlayers []entities.MatchPlayer
	err := r.db.Where("match_id = ?", matchID).Find(&matchPlayers).Error
	return matchPlayers, err
}

// GetByPlayerID retrieves all match statistics for a player
func (r *MatchPlayerRepositoryImpl) GetByPlayerID(playerID uint) ([]entities.MatchPlayer, error) {
	var matchPlayers []entities.MatchPlayer
	err := r.db.Where("player_id = ?", playerID).Find(&matchPlayers).Error
	return matchPlayers, err
}

// GetByTeamID retrieves all match statistics for a team
func (r *MatchPlayerRepositoryImpl) GetByTeamID(teamID uint) ([]entities.MatchPlayer, error) {
	var matchPlayers []entities.MatchPlayer
	err := r.db.Where("team_id = ?", teamID).Find(&matchPlayers).Error
	return matchPlayers, err
}

// GetPlayerStats retrieves player statistics for a specific season
func (r *MatchPlayerRepositoryImpl) GetPlayerStats(playerID uint, seasonID uint) ([]entities.MatchPlayer, error) {
	var matchPlayers []entities.MatchPlayer
	err := r.db.Joins("JOIN `match` ON `match`.id = match_player.match_id").
		Where("match_player.player_id = ? AND `match`.season_id = ?", playerID, seasonID).
		Find(&matchPlayers).Error
	return matchPlayers, err
}

// Update updates a match player record
func (r *MatchPlayerRepositoryImpl) Update(matchPlayer *entities.MatchPlayer) error {
	r.logger.Info("Updating match player with ID: %d", matchPlayer.ID)
	err := r.db.Model(matchPlayer).Updates(matchPlayer).Error
	if err != nil {
		r.logger.Error("Failed to update match player with ID %d: %v", matchPlayer.ID, err)
		return err
	}
	r.logger.Info("Successfully updated match player with ID: %d", matchPlayer.ID)
	return nil
}

// Delete deletes a match player record by ID
func (r *MatchPlayerRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.MatchPlayer{}, id).Error
}
