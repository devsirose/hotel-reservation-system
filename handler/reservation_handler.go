package handler

import (
	"net/http"
	"strconv"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/devsirose/hotel-reservation/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationHandler struct {
	reservationService service.ReservationService
}

func NewReservationHandler(reservationService service.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		reservationService: reservationService,
	}
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	var reservation model.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.reservationService.CreateReservation(c.Request.Context(), &reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, reservation)
}

func (h *ReservationHandler) GetReservation(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	reservation, err := h.reservationService.GetReservationByID(c.Request.Context(), reservationID)
	if err != nil {
		if err.Error() == "reservation not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservation)
}

func (h *ReservationHandler) ListReservationsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	
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

	reservations, err := h.reservationService.ListReservationsByUser(c.Request.Context(), userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      reservations,
		"user_id":   userID,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *ReservationHandler) ListReservationsByRoom(c *gin.Context) {
	roomIDStr := c.Param("room_id")
	roomID, err := uuid.Parse(roomIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid room ID"})
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

	reservations, err := h.reservationService.ListReservationsByRoom(c.Request.Context(), roomID, page, pageSize)
	if err != nil {
		if err.Error() == "room not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      reservations,
		"room_id":   roomID,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *ReservationHandler) UpdateReservation(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	var reservation model.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reservation.ReservationID = reservationID

	if err := h.reservationService.UpdateReservation(c.Request.Context(), &reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation updated successfully"})
}

func (h *ReservationHandler) CancelReservation(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	if err := h.reservationService.CancelReservation(c.Request.Context(), reservationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation cancelled successfully"})
}

func (h *ReservationHandler) ConfirmReservation(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	if err := h.reservationService.ConfirmReservation(c.Request.Context(), reservationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation confirmed successfully"})
}

func (h *ReservationHandler) ListReservations(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": "ListReservations not yet implemented",
		"message": "Please use /users/:user_id/reservations or /rooms/:room_id/reservations",
	})
}

func (h *ReservationHandler) UpdateReservationStatus(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if statusUpdate.Status == "CANCELLED" {
		if err := h.reservationService.CancelReservation(c.Request.Context(), reservationID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else if statusUpdate.Status == "CONFIRMED" {
		if err := h.reservationService.ConfirmReservation(c.Request.Context(), reservationID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status. Use CANCELLED or CONFIRMED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation status updated successfully"})
}

func (h *ReservationHandler) DeleteReservation(c *gin.Context) {
	reservationIDStr := c.Param("id")
	reservationID, err := uuid.Parse(reservationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	if err := h.reservationService.CancelReservation(c.Request.Context(), reservationID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "reservation deleted (cancelled) successfully"})
}