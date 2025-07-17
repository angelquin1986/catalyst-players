package entities

// LeaderboardEntry represents a single team's standing in the leaderboard.
type LeaderboardEntry struct {
	TeamID         uint   `json:"teamId"`
	TeamName       string `json:"teamName"`
	Played         int    `json:"played"`
	Won            int    `json:"won"`
	Drawn          int    `json:"drawn"`
	Lost           int    `json:"lost"`
	GoalsFor       int    `json:"goalsFor"`
	GoalsAgainst   int    `json:"goalsAgainst"`
	GoalDifference int    `json:"goalDifference"`
	Points         int    `json:"points"`
}

// Leaderboard represents the entire table of standings for a season.
type Leaderboard []LeaderboardEntry
