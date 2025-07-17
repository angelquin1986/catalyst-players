package handlers

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/domain/entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// MatchHandler handles HTTP requests for match operations
type MatchHandler struct {
	matchService *services.MatchService
}

// NewMatchHandler creates a new match handler
func NewMatchHandler(matchService *services.MatchService) *MatchHandler {
	return &MatchHandler{
		matchService: matchService,
	}
}

// CreateMatch handles POST /matches
func (h *MatchHandler) CreateMatch(c *gin.Context) {
	var match entities.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.matchService.CreateMatch(&match); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, match)
}

// GetMatch handles GET /matches/:id
func (h *MatchHandler) GetMatch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
		return
	}

	match, err := h.matchService.GetMatchByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// GetMatchWithDetails handles GET /matches/:id/details
func (h *MatchHandler) GetMatchWithDetails(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
		return
	}

	match, err := h.matchService.GetMatchWithDetails(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match not found"})
		return
	}

	c.JSON(http.StatusOK, match)
}

// GetAllMatches handles GET /matches
func (h *MatchHandler) GetAllMatches(c *gin.Context) {
	matches, err := h.matchService.GetAllMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchesBySeasonID handles GET /seasons/:id/matches
func (h *MatchHandler) GetMatchesBySeasonID(c *gin.Context) {
	seasonID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	matches, err := h.matchService.GetMatchesBySeasonID(uint(seasonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchesByStage handles GET /matches?season_id=...&stage=...
func (h *MatchHandler) GetMatchesByStage(c *gin.Context) {
	seasonIDStr := c.Query("season_id")
	stageStr := c.Query("stage")

	if seasonIDStr == "" || stageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "season_id and stage are required as query parameters"})
		return
	}

	seasonID, err := strconv.ParseUint(seasonIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season_id"})
		return
	}

	stage := entities.MatchStage(stageStr)
	matches, err := h.matchService.GetMatchesByStage(uint(seasonID), stage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchesByDateRange handles GET /matches/date-range
func (h *MatchHandler) GetMatchesByDateRange(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
		return
	}

	matches, err := h.matchService.GetMatchesByDateRange(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetMatchesByTeamID handles GET /teams/:id/matches
func (h *MatchHandler) GetMatchesByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	matches, err := h.matchService.GetMatchesByTeamID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetUpcoming handles GET /matches/upcoming
func (h *MatchHandler) GetUpcoming(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	matches, err := h.matchService.GetUpcoming(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// GetCompleted handles GET /seasons/:id/matches/completed
func (h *MatchHandler) GetCompleted(c *gin.Context) {
	seasonID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	matches, err := h.matchService.GetCompleted(uint(seasonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matches)
}

// UpdateMatch handles PUT /matches/:id
func (h *MatchHandler) UpdateMatch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
		return
	}

	var match entities.Match
	if err := c.ShouldBindJSON(&match); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match.ID = uint(id)
	if err := h.matchService.UpdateMatch(&match); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, match)
}

// UpdateMatchScore handles PUT /matches/:id/score
func (h *MatchHandler) UpdateMatchScore(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
		return
	}

	var request struct {
		HomeScore int `json:"home_score" binding:"required"`
		AwayScore int `json:"away_score" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.matchService.UpdateMatchScore(uint(id), request.HomeScore, request.AwayScore); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match score updated successfully"})
}

// DeleteMatch handles DELETE /matches/:id
func (h *MatchHandler) DeleteMatch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
		return
	}

	if err := h.matchService.DeleteMatch(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match deleted successfully"})
}
