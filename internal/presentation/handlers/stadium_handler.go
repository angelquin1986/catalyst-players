package handlers

import (
	"catalyst-players/internal/application/services"
	"catalyst-players/internal/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StadiumHandler handles HTTP requests for stadium operations
type StadiumHandler struct {
	stadiumService *services.StadiumService
}

// NewStadiumHandler creates a new stadium handler
func NewStadiumHandler(stadiumService *services.StadiumService) *StadiumHandler {
	return &StadiumHandler{
		stadiumService: stadiumService,
	}
}

// CreateStadium handles POST /stadiums
func (h *StadiumHandler) CreateStadium(c *gin.Context) {
	var stadium entities.Stadium
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.stadiumService.CreateStadium(&stadium); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stadium)
}

// GetStadium handles GET /stadiums/:id
func (h *StadiumHandler) GetStadium(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stadium ID"})
		return
	}

	stadium, err := h.stadiumService.GetStadiumByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stadium not found"})
		return
	}

	c.JSON(http.StatusOK, stadium)
}

// GetAllStadiums handles GET /stadiums
func (h *StadiumHandler) GetAllStadiums(c *gin.Context) {
	stadiums, err := h.stadiumService.GetAllStadiums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stadiums)
}

// UpdateStadium handles PUT /stadiums/:id
func (h *StadiumHandler) UpdateStadium(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stadium ID"})
		return
	}

	var stadium entities.Stadium
	if err := c.ShouldBindJSON(&stadium); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stadium.ID = uint(id)
	if err := h.stadiumService.UpdateStadium(&stadium); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stadium)
}

// DeleteStadium handles DELETE /stadiums/:id
func (h *StadiumHandler) DeleteStadium(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stadium ID"})
		return
	}

	if err := h.stadiumService.DeleteStadium(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stadium deleted successfully"})
} 