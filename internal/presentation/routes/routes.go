package routes

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/infrastructure/repositories"
	"catalyst-players/internal/presentation/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures the application's routes
func SetupRoutes(db *gorm.DB) *gin.Engine {
	// Initialize repositories
	stadiumRepo := repositories.NewStadiumRepositoryImpl(db)
	teamRepo := repositories.NewTeamRepositoryImpl(db)
	tagRepo := repositories.NewTagRepositoryImpl(db)
	matchRepo := repositories.NewMatchRepositoryImpl(db)
	matchPlayerRepo := repositories.NewMatchPlayerRepositoryImpl(db)
	seasonRepo := repositories.NewSeasonRepositoryImpl(db)
	leagueRepo := repositories.NewLeagueRepositoryImpl(db)
	playerRepo := repositories.NewPlayerRepositoryImpl(db)

	// Initialize services
	stadiumService := services.NewStadiumService(stadiumRepo)
	teamService := services.NewTeamService(teamRepo)
	tagService := services.NewTagService(tagRepo)
	matchService := services.NewMatchService(matchRepo)
	matchPlayerService := services.NewMatchPlayerService(matchPlayerRepo)
	seasonService := services.NewSeasonService(seasonRepo)
	leagueService := services.NewLeagueService(leagueRepo)
	playerService := services.NewPlayerService(playerRepo)
	leaderboardService := services.NewLeaderboardService(matchRepo) // Correct dependency

	// Initialize handlers
	stadiumHandler := handlers.NewStadiumHandler(stadiumService)
	teamHandler := handlers.NewTeamHandler(teamService)
	tagHandler := handlers.NewTagHandler(tagService)
	matchHandler := handlers.NewMatchHandler(matchService)
	matchPlayerHandler := handlers.NewMatchPlayerHandler(matchPlayerService)
	seasonHandler := handlers.NewSeasonHandler(seasonService)
	leagueHandler := handlers.NewLeagueHandler(leagueService)
	playerHandler := handlers.NewPlayerHandler(playerService)
	leaderboardHandler := handlers.NewLeaderboardHandler(leaderboardService)

	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	// API v1 routes
	apiV1 := router.Group("/api/v1")
	{
		// Health check
		apiV1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "message": "Catalyst Players API is running"})
		})

		// Tags routes
		tags := apiV1.Group("/tags")
		{
			tags.POST("", tagHandler.CreateTag)
			tags.GET("", tagHandler.GetAllTags)
			tags.GET("/:id", tagHandler.GetTag)
			tags.PUT("/:id", tagHandler.UpdateTag)
			tags.DELETE("/:id", tagHandler.DeleteTag)
		}

		// Stadiums routes
		stadiums := apiV1.Group("/stadiums")
		{
			stadiums.POST("", stadiumHandler.CreateStadium)
			stadiums.GET("", stadiumHandler.GetAllStadiums)
			stadiums.GET("/:id", stadiumHandler.GetStadium)
			stadiums.PUT("/:id", stadiumHandler.UpdateStadium)
			stadiums.DELETE("/:id", stadiumHandler.DeleteStadium)
		}

		// Teams routes
		teams := apiV1.Group("/teams")
		{
			teams.POST("", teamHandler.CreateTeam)
			teams.GET("", teamHandler.GetAllTeams)
			teams.GET("/:id", teamHandler.GetTeam)
			teams.GET("/:id/players", teamHandler.GetTeamWithPlayers)
			teams.PUT("/:id", teamHandler.UpdateTeam)
			teams.DELETE("/:id", teamHandler.DeleteTeam)
			teams.GET("/:id/matches", matchHandler.GetMatchesByTeamID)
			teams.GET("/:id/match-stats", matchPlayerHandler.GetMatchPlayersByTeamID)
		}

		// Players routes
		players := apiV1.Group("/players")
		{
			players.POST("", playerHandler.CreatePlayer)
			players.GET("", playerHandler.GetAllPlayers)
			players.GET("/:id", playerHandler.GetPlayer)
			players.GET("/:id/team", playerHandler.GetPlayerWithTeam)
			players.PUT("/:id", playerHandler.UpdatePlayer)
			players.DELETE("/:id", playerHandler.DeletePlayer)
			players.GET("/:id/match-stats", matchPlayerHandler.GetMatchPlayersByPlayerID)
			players.GET("/:id/stats/:season_id", matchPlayerHandler.GetPlayerStats)
			players.GET("/:id/tags", tagHandler.GetTagsByPlayerID)
		}

		// Leagues routes
		leagues := apiV1.Group("/leagues")
		{
			leagues.POST("", leagueHandler.CreateLeague)
			leagues.GET("", leagueHandler.GetAllLeagues)
			leagues.GET("/:id", leagueHandler.GetLeague)
			leagues.GET("/:id/seasons", seasonHandler.GetSeasonsByLeagueID)
			leagues.PUT("/:id", leagueHandler.UpdateLeague)
			leagues.DELETE("/:id", leagueHandler.DeleteLeague)
		}

		// Seasons routes
		seasonsGroup := apiV1.Group("/seasons")
		{
			seasonsGroup.GET("/:id/matches/completed", matchHandler.GetCompleted)
			seasonsGroup.POST("", seasonHandler.CreateSeason)
			seasonsGroup.GET("", seasonHandler.GetAllSeasons)
			seasonsGroup.GET("/active", seasonHandler.GetActiveSeasons)
			seasonsGroup.GET("/:id", seasonHandler.GetSeason)
			seasonsGroup.GET("/:id/league", seasonHandler.GetSeasonWithLeague)
			seasonsGroup.GET("/:id/teams", seasonHandler.GetSeasonWithTeams)
			seasonsGroup.GET("/:id/matches", matchHandler.GetMatchesBySeasonID)
			seasonsGroup.GET("/:id/standings", teamHandler.GetTeamStandings)
			seasonsGroup.GET("/:id/top-scorers", playerHandler.GetTopScorers)
			seasonsGroup.PUT("/:id", seasonHandler.UpdateSeason)
			seasonsGroup.PUT("/:id/activate", seasonHandler.ActivateSeason)
			seasonsGroup.PUT("/:id/complete", seasonHandler.CompleteSeason)
			seasonsGroup.DELETE("/:id", seasonHandler.DeleteSeason)
		}

		// Matches routes
		matchesGroup := apiV1.Group("/matches")
		{
			matchesGroup.GET("/upcoming", matchHandler.GetUpcoming)
			matchesGroup.POST("", matchHandler.CreateMatch)
			matchesGroup.GET("", matchHandler.GetAllMatches)
			matchesGroup.GET("/date-range", matchHandler.GetMatchesByDateRange)
			matchesGroup.GET("/:id", matchHandler.GetMatch)
			matchesGroup.GET("/:id/details", matchHandler.GetMatchWithDetails)
			matchesGroup.GET("/:id/players", matchPlayerHandler.GetMatchPlayersByMatchID)
			matchesGroup.PUT("/:id", matchHandler.UpdateMatch)
			matchesGroup.PUT("/:id/score", matchHandler.UpdateMatchScore)
			matchesGroup.DELETE("/:id", matchHandler.DeleteMatch)
			// matches.GET("/:season_id/:stage", matchHandler.GetMatchesByStage) // <-- Removed to avoid conflict
			// Now use: /api/v1/matches?season_id=...&stage=...
		}

		// Match players routes
		matchPlayers := apiV1.Group("/match-players")
		{
			matchPlayers.POST("", matchPlayerHandler.CreateMatchPlayer)
			matchPlayers.GET("", matchPlayerHandler.GetAllMatchPlayers)
			matchPlayers.GET("/:id", matchPlayerHandler.GetMatchPlayer)
			matchPlayers.PUT("/:id", matchPlayerHandler.UpdateMatchPlayer)
			matchPlayers.DELETE("/:id", matchPlayerHandler.DeleteMatchPlayer)
		}

		// Team tags routes
		teams.GET("/:id/tags", tagHandler.GetTagsByTeamID)

		// Leaderboard routes
		leaderboardGroup := apiV1.Group("/leaderboards")
		{
			leaderboardGroup.GET("/season/:seasonId", leaderboardHandler.GetLeaderboard)
		}
	}

	return router
}
