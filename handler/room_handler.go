package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/devsirose/hotel-reservation/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoomHandler struct {
	roomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var room model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.roomService.CreateRoom(c.Request.Context(), &room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, room)
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	roomIDStr := c.Param("id")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
		return
	}

	room, err := h.roomService.GetRoomByID(c.Request.Context(), roomID)
	if err != nil {
		if err.Error() == "room not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) ListRoomsByHotel(c *gin.Context) {
	hotelIDStr := c.Param("hotel_id")
	hotelID, err := uuid.Parse(hotelIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hotel ID"})
		return
	}

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

	rooms, err := h.roomService.ListRoomsByHotel(c.Request.Context(), hotelID, page, pageSize)
	if err != nil {
		if err.Error() == "hotel not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      rooms,
		"hotel_id":  hotelID,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *RoomHandler) UpdateRoom(c *gin.Context) {
	roomIDStr := c.Param("id")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
		return
	}

	var room model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.RoomID = roomID

	if err := h.roomService.UpdateRoom(c.Request.Context(), &room); err != nil {
		if err.Error() == "room not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "room updated successfully"})
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	roomIDStr := c.Param("id")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
		return
	}

	if err := h.roomService.DeleteRoom(c.Request.Context(), roomID); err != nil {
		if err.Error() == "room not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "room deleted successfully"})
}

func (h *RoomHandler) GetAvailableRooms(c *gin.Context) {
	hotelIDStr := c.Query("hotel_id")
	if hotelIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hotel_id query parameter is required"})
		return
	}

	hotelID, err := uuid.Parse(hotelIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hotel ID"})
		return
	}

	checkInStr := c.Query("check_in")
	checkOutStr := c.Query("check_out")

	if checkInStr == "" || checkOutStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "check_in and check_out dates are required"})
		return
	}

	checkIn, err := time.Parse("2006-01-02", checkInStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid check_in date format (use YYYY-MM-DD)"})
		return
	}

	checkOut, err := time.Parse("2006-01-02", checkOutStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid check_out date format (use YYYY-MM-DD)"})
		return
	}

	rooms, err := h.roomService.GetAvailableRooms(c.Request.Context(), hotelID, checkIn, checkOut)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      rooms,
		"hotel_id":  hotelID,
		"check_in":  checkInStr,
		"check_out": checkOutStr,
	})
}

func (h *RoomHandler) ListRooms(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "ListRooms not yet implemented",
		"message": "Please use /hotels/:hotel_id/rooms to list rooms by hotel",
	})
}