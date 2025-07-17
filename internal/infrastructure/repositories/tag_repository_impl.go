package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// TagRepositoryImpl implements the TagRepository interface
type TagRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewTagRepositoryImpl creates a new tag repository implementation
func NewTagRepositoryImpl(db *gorm.DB) repositories.TagRepository {
	return &TagRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new tag
func (r *TagRepositoryImpl) Create(tag *entities.Tag) error {
	return r.db.Create(tag).Error
}

// GetByID retrieves a tag by ID
func (r *TagRepositoryImpl) GetByID(id uint) (*entities.Tag, error) {
	var tag entities.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// GetAll retrieves all tags
func (r *TagRepositoryImpl) GetAll() ([]entities.Tag, error) {
	var tags []entities.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

// Update updates a tag
func (r *TagRepositoryImpl) Update(tag *entities.Tag) error {
	r.logger.Info("Updating tag with ID: %d", tag.ID)
	err := r.db.Model(tag).Updates(tag).Error
	if err != nil {
		r.logger.Error("Failed to update tag with ID %d: %v", tag.ID, err)
		return err
	}
	r.logger.Info("Successfully updated tag with ID: %d", tag.ID)
	return nil
}

// Delete deletes a tag by ID
func (r *TagRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Tag{}, id).Error
}

// GetByPlayerID retrieves all tags associated with a player
func (r *TagRepositoryImpl) GetByPlayerID(playerID uint) ([]entities.Tag, error) {
	var tags []entities.Tag
	err := r.db.Joins("JOIN tag_player ON tag_player.tag_id = tag.id").
		Where("tag_player.player_id = ?", playerID).
		Find(&tags).Error
	return tags, err
}

// GetByTeamID retrieves all tags associated with a team
func (r *TagRepositoryImpl) GetByTeamID(teamID uint) ([]entities.Tag, error) {
	var tags []entities.Tag
	err := r.db.Joins("JOIN tag_team ON tag_team.tag_id = tag.id").
		Where("tag_team.team_id = ?", teamID).
		Find(&tags).Error
	return tags, err
}
