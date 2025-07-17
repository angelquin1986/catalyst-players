package entities

import (
	"time"
)

// Team represents a soccer team entity
type Team struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	BirthDate time.Time `json:"birth_date" gorm:"type:timestamp"`
	Category  string    `json:"category" gorm:"size:255"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
	// Relationships
	Players []Player `json:"players,omitempty" gorm:"foreignKey:TeamID"`
	Tags    []Tag    `json:"tags,omitempty" gorm:"many2many:tag_team;"`
}

// TableName specifies the table name for Team
func (Team) TableName() string {
	return "team"
} 