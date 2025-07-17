package handlers

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MatchPlayerHandler handles HTTP requests for match player statistics operations
type MatchPlayerHandler struct {
	matchPlayerService *services.MatchPlayerService
}

// NewMatchPlayerHandler creates a new match player handler
func NewMatchPlayerHandler(matchPlayerService *services.MatchPlayerService) *MatchPlayerHandler {
	return &MatchPlayerHandler{
		matchPlayerService: matchPlayerService,
	}
}

// CreateMatchPlayer handles POST /match-players
func (h *MatchPlayerHandler) CreateMatchPlayer(c *gin.Context) {
	var matchPlayer entities.MatchPlayer
	if err := c.ShouldBindJSON(&matchPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.matchPlayerService.CreateMatchPlayer(&matchPlayer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, matchPlayer)
}

// GetMatchPlayer handles GET /match-players/:id
func (h *MatchPlayerHandler) GetMatchPlayer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match player ID"})
		return
	}

	matchPlayer, err := h.matchPlayerService.GetMatchPlayerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Match player not found"})
		return
	}

	c.JSON(http.StatusOK, matchPlayer)
}

// GetAllMatchPlayers handles GET /match-players
func (h *MatchPlayerHandler) GetAllMatchPlayers(c *gin.Context) {
	matchPlayers, err := h.matchPlayerService.GetAllMatchPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matchPlayers)
}

// GetMatchPlayersByMatchID handles GET /matches/:id/players
func (h *MatchPlayerHandler) GetMatchPlayersByMatchID(c *gin.Context) {
	matchID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match ID"})
		return
	}

	matchPlayers, err := h.matchPlayerService.GetMatchPlayersByMatchID(uint(matchID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matchPlayers)
}

// GetMatchPlayersByPlayerID handles GET /players/:id/match-stats
func (h *MatchPlayerHandler) GetMatchPlayersByPlayerID(c *gin.Context) {
	playerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	matchPlayers, err := h.matchPlayerService.GetMatchPlayersByPlayerID(uint(playerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matchPlayers)
}

// GetMatchPlayersByTeamID handles GET /teams/:id/match-stats
func (h *MatchPlayerHandler) GetMatchPlayersByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	matchPlayers, err := h.matchPlayerService.GetMatchPlayersByTeamID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matchPlayers)
}

// GetPlayerStats handles GET /players/:id/stats/:season_id
func (h *MatchPlayerHandler) GetPlayerStats(c *gin.Context) {
	playerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	seasonID, err := strconv.ParseUint(c.Param("season_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	matchPlayers, err := h.matchPlayerService.GetPlayerStats(uint(playerID), uint(seasonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matchPlayers)
}

// UpdateMatchPlayer handles PUT /match-players/:id
func (h *MatchPlayerHandler) UpdateMatchPlayer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match player ID"})
		return
	}

	var matchPlayer entities.MatchPlayer
	if err := c.ShouldBindJSON(&matchPlayer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	matchPlayer.ID = uint(id)
	if err := h.matchPlayerService.UpdateMatchPlayer(&matchPlayer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, matchPlayer)
}

// DeleteMatchPlayer handles DELETE /match-players/:id
func (h *MatchPlayerHandler) DeleteMatchPlayer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid match player ID"})
		return
	}

	if err := h.matchPlayerService.DeleteMatchPlayer(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Match player deleted successfully"})
} 