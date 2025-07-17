package handlers

import (
	"catalyst-players/internal/application/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LeaderboardHandler handles requests related to leaderboards.
type LeaderboardHandler struct {
	leaderboardService *services.LeaderboardService
}

// NewLeaderboardHandler creates a new LeaderboardHandler.
func NewLeaderboardHandler(leaderboardService *services.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{leaderboardService: leaderboardService}
}

// GetLeaderboard retrieves and returns the leaderboard for a specific season.
// @Summary Get season leaderboard
// @Description Get the calculated leaderboard for a given season ID.
// @Tags Leaderboards
// @Accept json
// @Produce json
// @Param seasonId path int true "Season ID"
// @Success 200 {object} entities.Leaderboard
// @Failure 400 {object} map[string]string "Invalid season ID"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /leaderboards/season/{seasonId} [get]
func (h *LeaderboardHandler) GetLeaderboard(c *gin.Context) {
	seasonIDStr := c.Param("seasonId")
	seasonID, err := strconv.ParseUint(seasonIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	leaderboard, err := h.leaderboardService.GenerateLeaderboard(uint(seasonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate leaderboard"})
		return
	}

	c.JSON(http.StatusOK, leaderboard)
}
