package entities

import (
	"time"
)

// Player represents a soccer player entity
type Player struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:255;not null"`
	LastName  string    `json:"last_name" gorm:"size:255;not null"`
	BirthDate time.Time `json:"birth_date" gorm:"type:timestamp"`
	TeamID    uint      `json:"team_id" gorm:"not null"`
	Number    int       `json:"number" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
	// Relationships
	Team      Team      `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Tags      []Tag     `json:"tags,omitempty" gorm:"many2many:tag_player;"`
	MatchStats []MatchPlayer `json:"match_stats,omitempty" gorm:"foreignKey:PlayerID"`
}

// TableName specifies the table name for Player
func (Player) TableName() string {
	return "player"
} 