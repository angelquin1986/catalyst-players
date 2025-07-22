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

// TeamHandler handles HTTP requests for team operations
type TeamHandler struct {
	teamService *services.TeamService
}

// NewTeamHandler creates a new team handler
func NewTeamHandler(teamService *services.TeamService) *TeamHandler {
	return &TeamHandler{
		teamService: teamService,
	}
}

// CreateTeam handles POST /teams
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var request struct {
		Name      string `json:"name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		Category  string `json:"category" binding:"required"`
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

	team := &entities.Team{
		Name:      request.Name,
		BirthDate: birthDate,
		Category:  request.Category,
		Photo:     request.Photo, // Store bytes directly
	}

	if err := h.teamService.CreateTeam(team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

// GetTeam handles GET /teams/:id
func (h *TeamHandler) GetTeam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	team, err := h.teamService.GetTeamByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	// Photo is already []byte, return as is
	c.JSON(http.StatusOK, team)
}

// GetTeamWithPlayers handles GET /teams/:id/players
func (h *TeamHandler) GetTeamWithPlayers(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	team, err := h.teamService.GetTeamWithPlayers(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	// Photo is already []byte, return as is
	c.JSON(http.StatusOK, team)
}

// GetAllTeams handles GET /teams
func (h *TeamHandler) GetAllTeams(c *gin.Context) {
	teams, err := h.teamService.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Photos are already []byte, return as is
	c.JSON(http.StatusOK, teams)
}

// UpdateTeam handles PUT /teams/:id
func (h *TeamHandler) UpdateTeam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	var request struct {
		Name      string `json:"name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		Category  string `json:"category" binding:"required"`
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

	team := &entities.Team{
		ID:        uint(id),
		Name:      request.Name,
		BirthDate: birthDate,
		Category:  request.Category,
		Photo:     request.Photo, // Store bytes directly
	}

	if err := h.teamService.UpdateTeam(team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

// DeleteTeam handles DELETE /teams/:id
func (h *TeamHandler) DeleteTeam(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid team ID"})
		return
	}

	if err := h.teamService.DeleteTeam(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team deleted successfully"})
}

// GetTeamsBySeasonID handles GET /seasons/:id/teams
func (h *TeamHandler) GetTeamsBySeasonID(c *gin.Context) {
	seasonID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	teams, err := h.teamService.GetTeamsBySeasonID(uint(seasonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Photos are already []byte, return as is
	c.JSON(http.StatusOK, teams)
}

// GetTeamStandings handles GET /seasons/:id/standings
func (h *TeamHandler) GetTeamStandings(c *gin.Context) {
	seasonID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid season ID"})
		return
	}

	teams, err := h.teamService.GetTeamStandings(uint(seasonID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Photos are already []byte, return as is
	c.JSON(http.StatusOK, teams)
}
