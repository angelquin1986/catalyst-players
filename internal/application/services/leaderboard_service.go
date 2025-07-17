package services

import (
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/domain/repositories"
	"sort"
)

// LeaderboardService provides services for generating leaderboards.
type LeaderboardService struct {
	matchRepo repositories.MatchRepository
}

// NewLeaderboardService creates a new LeaderboardService.
func NewLeaderboardService(matchRepo repositories.MatchRepository) *LeaderboardService {
	return &LeaderboardService{matchRepo: matchRepo}
}

// GenerateLeaderboard calculates and returns the leaderboard for a given season.
func (s *LeaderboardService) GenerateLeaderboard(seasonID uint) (entities.Leaderboard, error) {
	// 1. Fetch all finished matches for the season
	matches, err := s.matchRepo.GetCompleted(seasonID)
	if err != nil {
		return nil, err
	}

	// 2. Process matches to calculate standings
	standings := make(map[uint]*entities.LeaderboardEntry)

	for _, match := range matches {
		// Ensure teams are loaded
		if match.HomeTeam.ID == 0 || match.AwayTeam.ID == 0 {
			continue // Or handle error, for now skip if team data is missing
		}

		// Get or create entries for home and away teams
		if _, ok := standings[match.HomeTeamID]; !ok {
			standings[match.HomeTeamID] = &entities.LeaderboardEntry{TeamID: match.HomeTeamID, TeamName: match.HomeTeam.Name}
		}
		if _, ok := standings[match.AwayTeamID]; !ok {
			standings[match.AwayTeamID] = &entities.LeaderboardEntry{TeamID: match.AwayTeamID, TeamName: match.AwayTeam.Name}
		}

		homeEntry := standings[match.HomeTeamID]
		awayEntry := standings[match.AwayTeamID]

		// Update stats based on score
		if match.HomeTeamScore != nil && match.AwayTeamScore != nil {
			homeScore := *match.HomeTeamScore
			awayScore := *match.AwayTeamScore

			homeEntry.Played++
			awayEntry.Played++

			homeEntry.GoalsFor += homeScore
			homeEntry.GoalsAgainst += awayScore
			awayEntry.GoalsFor += awayScore
			awayEntry.GoalsAgainst += homeScore

			if homeScore > awayScore { // Home team wins
				homeEntry.Won++
				homeEntry.Points += 3
				awayEntry.Lost++
			} else if awayScore > homeScore { // Away team wins
				awayEntry.Won++
				awayEntry.Points += 3
				homeEntry.Lost++
			} else { // Draw
				homeEntry.Drawn++
				homeEntry.Points++
				awayEntry.Drawn++
				awayEntry.Points++
			}
		}
	}

	// 3. Convert map to slice and calculate goal difference
	leaderboard := make(entities.Leaderboard, 0, len(standings))
	for _, entry := range standings {
		entry.GoalDifference = entry.GoalsFor - entry.GoalsAgainst
		leaderboard = append(leaderboard, *entry)
	}

	// 4. Sort the leaderboard
	sort.Slice(leaderboard, func(i, j int) bool {
		if leaderboard[i].Points != leaderboard[j].Points {
			return leaderboard[i].Points > leaderboard[j].Points // More points is better
		}
		if leaderboard[i].GoalDifference != leaderboard[j].GoalDifference {
			return leaderboard[i].GoalDifference > leaderboard[j].GoalDifference // Higher GD is better
		}
		if leaderboard[i].GoalsFor != leaderboard[j].GoalsFor {
			return leaderboard[i].GoalsFor > leaderboard[j].GoalsFor // More goals scored is better
		}
		return leaderboard[i].TeamName < leaderboard[j].TeamName // Alphabetical as a tie-breaker
	})

	return leaderboard, nil
}
