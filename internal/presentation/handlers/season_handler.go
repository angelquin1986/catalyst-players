package handlers

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/domain/entities"
	"catalyst-players/internal/infrastructure/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SeasonHandler handles HTTP requests for season operations
type SeasonHandler struct {
	logger        logger.Logger
	seasonService *services.SeasonService
}

// NewSeasonHandler creates a new season handler
func NewSeasonHandler(seasonService *services.SeasonService) *SeasonHandler {
	return &SeasonHandler{
		logger:        logger.NewLogger(),
		seasonService: seasonService,
	}
}

// CreateSeason handles POST /seasons
func (h *SeasonHandler) CreateSeason(c *gin.Context) {
	h.logger.Info("Creating new season")
	var season entities.Season
	if err := c.ShouldBindJSON(&season); err != nil {
		h.logger.Error("Failed to bind JSON for season creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.seasonService.CreateSeason(&season); err != nil {
		h.logger.Error("Failed to create season: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.logger.Info("Successfully created season with ID: %d", season.ID)
	c.JSON(http.StatusCreated, season)
}

// GetSeason handles GET /seasons/:id
func (h *SeasonHandler) GetSeason(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.logger.Error("Invalid season ID format: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	h.logger.Info("Getting season with ID: %d", uint(id))
	season, err := h.seasonService.GetSeasonByID(uint(id))
	if err != nil {
		h.logger.Error("Failed to get season with ID %d: %v", uint(id), err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Season not found"})
		return
	}

	h.logger.Info("Successfully retrieved season with ID: %d", uint(id))
	c.JSON(http.StatusOK, season)
}

// GetSeasonWithLeague handles GET /seasons/:id/league
func (h *SeasonHandler) GetSeasonWithLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	season, err := h.seasonService.GetSeasonWithLeague(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Season not found"})
		return
	}

	c.JSON(http.StatusOK, season)
}

// GetSeasonWithTeams handles GET /seasons/:id/teams
func (h *SeasonHandler) GetSeasonWithTeams(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	season, err := h.seasonService.GetSeasonWithTeams(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Season not found"})
		return
	}

	c.JSON(http.StatusOK, season)
}

// GetAllSeasons handles GET /seasons
func (h *SeasonHandler) GetAllSeasons(c *gin.Context) {
	h.logger.Info("Getting all seasons")
	seasons, err := h.seasonService.GetAllSeasons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seasons)
}

// GetAllSeasonsWithTeams handles GET /seasons/with-teams
func (h *SeasonHandler) GetAllSeasonsWithTeams(c *gin.Context) {
	h.logger.Info("Getting all seasons with teams")
	seasons, err := h.seasonService.GetAllSeasonsWithTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seasons)
}

// GetActiveSeasons handles GET /seasons/active
func (h *SeasonHandler) GetActiveSeasons(c *gin.Context) {
	seasons, err := h.seasonService.GetActiveSeasons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seasons)
}

// GetSeasonsByLeagueID handles GET /leagues/:id/seasons
func (h *SeasonHandler) GetSeasonsByLeagueID(c *gin.Context) {
	leagueID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	seasons, err := h.seasonService.GetSeasonsByLeagueID(uint(leagueID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, seasons)
}

// UpdateSeason handles PUT /seasons/:id
func (h *SeasonHandler) UpdateSeason(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	var season entities.Season
	if err := c.ShouldBindJSON(&season); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	season.ID = uint(id)
	if err := h.seasonService.UpdateSeason(&season); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, season)
}

// DeleteSeason handles DELETE /seasons/:id
func (h *SeasonHandler) DeleteSeason(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	if err := h.seasonService.DeleteSeason(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Season deleted successfully"})
}

// ActivateSeason handles PUT /seasons/:id/activate
func (h *SeasonHandler) ActivateSeason(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	if err := h.seasonService.ActivateSeason(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Season activated successfully"})
}

// CompleteSeason handles PUT /seasons/:id/complete
func (h *SeasonHandler) CompleteSeason(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	if err := h.seasonService.CompleteSeason(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Season completed successfully"})
}
