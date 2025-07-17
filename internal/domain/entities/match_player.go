package entities

import (
	"time"
)

// MatchPlayer represents individual player statistics in a match
type MatchPlayer struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	MatchID    uint      `json:"match_id" gorm:"not null"`
	TeamID     uint      `json:"team_id" gorm:"not null"`
	PlayerID   uint      `json:"player_id" gorm:"not null"`
	RedCard    int       `json:"red_card" gorm:"type:int;default:0"`
	YellowCard int       `json:"yellow_card" gorm:"type:int;default:0"`
	Goals      int       `json:"goals" gorm:"type:int;default:0"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	
	// Relationships
	Match      Match     `json:"match,omitempty" gorm:"foreignKey:MatchID"`
	Team       Team      `json:"team,omitempty" gorm:"foreignKey:TeamID"`
	Player     Player    `json:"player,omitempty" gorm:"foreignKey:PlayerID"`
}

// TableName specifies the table name for MatchPlayer
func (MatchPlayer) TableName() string {
	return "match_player"
} 