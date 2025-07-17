package repositories

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"catalyst-players/internal/infrastructure/logger"

	"gorm.io/gorm"
)

// StadiumRepositoryImpl implements the StadiumRepository interface
type StadiumRepositoryImpl struct {
	db     *gorm.DB
	logger logger.Logger
}

// NewStadiumRepositoryImpl creates a new stadium repository implementation
func NewStadiumRepositoryImpl(db *gorm.DB) repositories.StadiumRepository {
	return &StadiumRepositoryImpl{
		db:     db,
		logger: logger.NewLogger(),
	}
}

// Create creates a new stadium
func (r *StadiumRepositoryImpl) Create(stadium *entities.Stadium) error {
	return r.db.Create(stadium).Error
}

// GetByID retrieves a stadium by ID
func (r *StadiumRepositoryImpl) GetByID(id uint) (*entities.Stadium, error) {
	var stadium entities.Stadium
	err := r.db.First(&stadium, id).Error
	if err != nil {
		return nil, err
	}
	return &stadium, nil
}

// GetAll retrieves all stadiums
func (r *StadiumRepositoryImpl) GetAll() ([]entities.Stadium, error) {
	var stadiums []entities.Stadium
	err := r.db.Find(&stadiums).Error
	return stadiums, err
}

// Update updates a stadium
func (r *StadiumRepositoryImpl) Update(stadium *entities.Stadium) error {
	r.logger.Info("Updating stadium with ID: %d", stadium.ID)
	err := r.db.Model(stadium).Updates(stadium).Error
	if err != nil {
		r.logger.Error("Failed to update stadium with ID %d: %v", stadium.ID, err)
		return err
	}
	r.logger.Info("Successfully updated stadium with ID: %d", stadium.ID)
	return nil
}

// Delete deletes a stadium by ID
func (r *StadiumRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&entities.Stadium{}, id).Error
}
