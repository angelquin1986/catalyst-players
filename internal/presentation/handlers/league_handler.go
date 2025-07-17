package handlers

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// LeagueHandler handles HTTP requests for league operations
type LeagueHandler struct {
	leagueService *services.LeagueService
}

// NewLeagueHandler creates a new league handler
func NewLeagueHandler(leagueService *services.LeagueService) *LeagueHandler {
	return &LeagueHandler{
		leagueService: leagueService,
	}
}

// CreateLeague handles POST /leagues
func (h *LeagueHandler) CreateLeague(c *gin.Context) {
	var league entities.League
	if err := c.ShouldBindJSON(&league); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.leagueService.CreateLeague(&league); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, league)
}

// GetLeague handles GET /leagues/:id
func (h *LeagueHandler) GetLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	league, err := h.leagueService.GetLeagueByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}

	c.JSON(http.StatusOK, league)
}

// GetLeagueWithSeasons handles GET /leagues/:id/seasons
func (h *LeagueHandler) GetLeagueWithSeasons(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	league, err := h.leagueService.GetLeagueWithSeasons(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "League not found"})
		return
	}

	c.JSON(http.StatusOK, league)
}

// GetAllLeagues handles GET /leagues
func (h *LeagueHandler) GetAllLeagues(c *gin.Context) {
	leagues, err := h.leagueService.GetAllLeagues()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, leagues)
}

// UpdateLeague handles PUT /leagues/:id
func (h *LeagueHandler) UpdateLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	var league entities.League
	if err := c.ShouldBindJSON(&league); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	league.ID = uint(id)
	if err := h.leagueService.UpdateLeague(&league); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, league)
}

// DeleteLeague handles DELETE /leagues/:id
func (h *LeagueHandler) DeleteLeague(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid league ID"})
		return
	}

	if err := h.leagueService.DeleteLeague(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "League deleted successfully"})
} 