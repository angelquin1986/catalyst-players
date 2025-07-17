package entities

import (
	"time"
)

// League represents a soccer league entity
type League struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	BirthDate time.Time `json:"birth_date" gorm:"type:timestamp"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
	// Relationships
	Seasons []Season `json:"seasons,omitempty" gorm:"foreignKey:LeagueID"`
}

// TableName specifies the table name for League
func (League) TableName() string {
	return "league"
} 