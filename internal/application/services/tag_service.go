package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"errors"
)

// TagService handles business logic for tag operations
type TagService struct {
	tagRepo repositories.TagRepository
}

// NewTagService creates a new tag service instance
func NewTagService(tagRepo repositories.TagRepository) *TagService {
	return &TagService{
		tagRepo: tagRepo,
	}
}

// CreateTag creates a new tag
func (s *TagService) CreateTag(tag *entities.Tag) error {
	if tag.Name == "" {
		return errors.New("tag name is required")
	}
	
	return s.tagRepo.Create(tag)
}

// GetTagByID retrieves a tag by ID
func (s *TagService) GetTagByID(id uint) (*entities.Tag, error) {
	if id == 0 {
		return nil, errors.New("invalid tag ID")
	}
	
	return s.tagRepo.GetByID(id)
}

// GetAllTags retrieves all tags
func (s *TagService) GetAllTags() ([]entities.Tag, error) {
	return s.tagRepo.GetAll()
}

// UpdateTag updates an existing tag
func (s *TagService) UpdateTag(tag *entities.Tag) error {
	if tag.ID == 0 {
		return errors.New("invalid tag ID")
	}
	
	if tag.Name == "" {
		return errors.New("tag name is required")
	}
	
	return s.tagRepo.Update(tag)
}

// DeleteTag deletes a tag by ID
func (s *TagService) DeleteTag(id uint) error {
	if id == 0 {
		return errors.New("invalid tag ID")
	}
	
	return s.tagRepo.Delete(id)
}

// GetTagsByPlayerID retrieves all tags associated with a player
func (s *TagService) GetTagsByPlayerID(playerID uint) ([]entities.Tag, error) {
	if playerID == 0 {
		return nil, errors.New("invalid player ID")
	}
	
	return s.tagRepo.GetByPlayerID(playerID)
}

// GetTagsByTeamID retrieves all tags associated with a team
func (s *TagService) GetTagsByTeamID(teamID uint) ([]entities.Tag, error) {
	if teamID == 0 {
		return nil, errors.New("invalid team ID")
	}
	
	return s.tagRepo.GetByTeamID(teamID)
} 