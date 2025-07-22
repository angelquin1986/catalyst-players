package handlers

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/domain/entities"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PlayerHandler handles HTTP requests for player operations
type PlayerHandler struct {
	playerService *services.PlayerService
}

// NewPlayerHandler creates a new player handler
func NewPlayerHandler(playerService *services.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}

// CreatePlayer handles POST /players
func (h *PlayerHandler) CreatePlayer(c *gin.Context) {
	var request struct {
		ID        string `json:"id" binding:"required"` // Identity card (cedula)
		Name      string `json:"name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		TeamID    uint   `json:"team_id" binding:"required"`
		Number    int    `json:"number" binding:"required"`
		Photo     []byte `json:"photo"` // Raw bytes of the photo
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse birth date
	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birth date format (use YYYY-MM-DD)"})
		return
	}

	player := &entities.Player{
		ID:        request.ID,
		Name:      request.Name,
		LastName:  request.LastName,
		BirthDate: birthDate,
		TeamID:    request.TeamID,
		Number:    request.Number,
		Photo:     request.Photo, // Store bytes directly
	}

	if err := h.playerService.CreatePlayer(player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, player)
}

// GetPlayer handles GET /players/:id
func (h *PlayerHandler) GetPlayer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	player, err := h.playerService.GetPlayerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Photo is already []byte, return as is
	c.JSON(http.StatusOK, player)
}

// GetPlayerWithTeam handles GET /players/:id/team
func (h *PlayerHandler) GetPlayerWithTeam(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	player, err := h.playerService.GetPlayerWithTeam(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	// Photo is already []byte, return as is
	c.JSON(http.StatusOK, player)
}

// GetAllPlayers handles GET /players
func (h *PlayerHandler) GetAllPlayers(c *gin.Context) {
	players, err := h.playerService.GetAllPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Photos are already []byte, return as is
	c.JSON(http.StatusOK, players)
}

// GetPlayersByTeamID handles GET /teams/:id/players
func (h *PlayerHandler) GetPlayersByTeamID(c *gin.Context) {
	teamID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	players, err := h.playerService.GetPlayersByTeamID(uint(teamID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Photos are already []byte, return as is
	c.JSON(http.StatusOK, players)
}

// UpdatePlayer handles PUT /players/:id
func (h *PlayerHandler) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	var request struct {
		ID        string `json:"id" binding:"required"` // Identity card (cedula)
		Name      string `json:"name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		TeamID    uint   `json:"team_id" binding:"required"`
		Number    int    `json:"number" binding:"required"`
		Photo     []byte `json:"photo"` // Raw bytes of the photo
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse birth date
	birthDate, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid birth date format (use YYYY-MM-DD)"})
		return
	}

	player := &entities.Player{
		ID:        request.ID,
		Name:      request.Name,
		LastName:  request.LastName,
		BirthDate: birthDate,
		TeamID:    request.TeamID,
		Number:    request.Number,
		Photo:     request.Photo, // Store bytes directly
	}

	if err := h.playerService.UpdatePlayer(player); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, player)
}

// DeletePlayer handles DELETE /players/:id
func (h *PlayerHandler) DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
		return
	}

	if err := h.playerService.DeletePlayer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully"})
}

// GetTopScorers handles GET /seasons/:id/top-scorers
func (h *PlayerHandler) GetTopScorers(c *gin.Context) {
	seasonID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	players, err := h.playerService.GetTopScorers(uint(seasonID), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Photos are already []byte, return as is
	c.JSON(http.StatusOK, players)
}
