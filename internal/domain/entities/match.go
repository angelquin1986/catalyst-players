package entities

import (
	"time"
)

// MatchStatus defines the possible statuses for a match
type MatchStatus string

const (
	MatchStatusScheduled  MatchStatus = "scheduled"
	MatchStatusInProgress MatchStatus = "in_progress"
	MatchStatusFinished   MatchStatus = "finished"
	MatchStatusPostponed  MatchStatus = "postponed"
	MatchStatusCancelled  MatchStatus = "cancelled"
)

// MatchStage represents the stage of a match in the tournament
type MatchStage string

const (
	MatchStageRegular  MatchStage = "regular"
	MatchStageQuarters MatchStage = "quarters"
	MatchStageSemis    MatchStage = "semis"
	MatchStageFinal    MatchStage = "final"
	MatchStageThird    MatchStage = "third"
)

// Match represents a soccer match entity
type Match struct {
	ID             uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	HomeTeamID     uint       `json:"home_team_id" gorm:"not null"`
	AwayTeamID     uint       `json:"away_team_id" gorm:"not null"`
	SeasonID       uint       `json:"season_id" gorm:"not null"`
	StadiumID      uint       `json:"stadium_id" gorm:"not null"`
	Date           time.Time  `json:"date" gorm:"type:timestamp;not null"`
	Hour           *int       `json:"hour" gorm:"type:int"`
	HomeTeamScore  *int       `json:"home_team_score" gorm:"type:int"`
	AwayTeamScore  *int       `json:"away_team_score" gorm:"type:int"`
	HomeTeamPoints *int       `json:"home_team_points" gorm:"type:int"`
	AwayTeamPoints *int       `json:"away_team_points" gorm:"type:int"`
	Stage          MatchStage `json:"stage" gorm:"size:255;not null"`
	Observation    string     `json:"observation" gorm:"type:text"`
	Status         string     `json:"status"`
	Round          int        `json:"round"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`

	// Relationships
	HomeTeam    Team          `json:"home_team,omitempty" gorm:"foreignKey:HomeTeamID"`
	AwayTeam    Team          `json:"away_team,omitempty" gorm:"foreignKey:AwayTeamID"`
	Season      Season        `json:"season,omitempty" gorm:"foreignKey:SeasonID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL,name:fk_matches_season"`
	Stadium     Stadium       `json:"stadium,omitempty" gorm:"foreignKey:StadiumID"`
	PlayerStats []MatchPlayer `json:"player_stats,omitempty" gorm:"foreignKey:MatchID"`
}

// TableName specifies the table name for Match
func (Match) TableName() string {
	return "match"
}
