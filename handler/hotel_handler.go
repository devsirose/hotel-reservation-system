package handler

import (
	"net/http"
	"strconv"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/devsirose/hotel-reservation/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HotelHandler struct {
	hotelService service.HotelService
}

func NewHotelHandler(hotelService service.HotelService) *HotelHandler {
	return &HotelHandler{
		hotelService: hotelService,
	}
}

func (h *HotelHandler) CreateHotel(c *gin.Context) {
	var hotel model.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.hotelService.CreateHotel(c.Request.Context(), &hotel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hotel)
}

func (h *HotelHandler) GetHotel(c *gin.Context) {
	hotelIDStr := c.Param("id")
	hotelID, err := uuid.Parse(hotelIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hotel ID"})
		return
	}

	hotel, err := h.hotelService.GetHotelByID(c.Request.Context(), hotelID)
	if err != nil {
		if err.Error() == "hotel not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotel)
}

func (h *HotelHandler) ListHotels(c *gin.Context) {
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		if parsedPage, err := strconv.Atoi(p); err == nil {
			page = parsedPage
		}
	}

	if ps := c.Query("page_size"); ps != "" {
		if parsedPageSize, err := strconv.Atoi(ps); err == nil {
			pageSize = parsedPageSize
		}
	}

	hotels, err := h.hotelService.ListHotels(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      hotels,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *HotelHandler) UpdateHotel(c *gin.Context) {
	hotelIDStr := c.Param("id")
	hotelID, err := uuid.Parse(hotelIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hotel ID"})
		return
	}

	var hotel model.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotel.HotelID = hotelID

	if err := h.hotelService.UpdateHotel(c.Request.Context(), &hotel); err != nil {
		if err.Error() == "hotel not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "hotel updated successfully"})
}

func (h *HotelHandler) DeleteHotel(c *gin.Context) {
	hotelIDStr := c.Param("id")
	hotelID, err := uuid.Parse(hotelIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hotel ID"})
		return
	}

	if err := h.hotelService.DeleteHotel(c.Request.Context(), hotelID); err != nil {
		if err.Error() == "hotel not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "hotel deleted successfully"})
}