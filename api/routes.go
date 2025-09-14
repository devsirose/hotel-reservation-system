package api

import (
	"net/http"
	
	"github.com/devsirose/hotel-reservation/middleware"
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRoutes() *gin.Engine {
	router := gin.Default()
	
	// Add middleware
	router.Use(middleware.RecoveryWithLogger)
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())
	
	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Hotel routes
		hotels := v1.Group("/hotels")
		{
			hotels.POST("", server.hotelHandler.CreateHotel)
			hotels.GET("/:id", server.hotelHandler.GetHotel)
			hotels.GET("", server.hotelHandler.ListHotels)
			hotels.PUT("/:id", server.hotelHandler.UpdateHotel)
			hotels.DELETE("/:id", server.hotelHandler.DeleteHotel)
		}
		
		// Room routes
		rooms := v1.Group("/rooms")
		{
			rooms.POST("", server.roomHandler.CreateRoom)
			rooms.GET("/:id", server.roomHandler.GetRoom)
			rooms.GET("", server.roomHandler.ListRooms)
			rooms.GET("/hotel/:hotel_id", server.roomHandler.ListRoomsByHotel)
			rooms.GET("/available", server.roomHandler.GetAvailableRooms)
			rooms.PUT("/:id", server.roomHandler.UpdateRoom)
			rooms.DELETE("/:id", server.roomHandler.DeleteRoom)
		}
		
		// Reservation routes
		reservations := v1.Group("/reservations")
		{
			reservations.POST("", server.reservHandler.CreateReservation)
			reservations.GET("/:id", server.reservHandler.GetReservation)
			reservations.GET("", server.reservHandler.ListReservations)
			reservations.GET("/user/:user_id", server.reservHandler.ListReservationsByUser)
			reservations.GET("/room/:room_id", server.reservHandler.ListReservationsByRoom)
			reservations.PUT("/:id", server.reservHandler.UpdateReservation)
			reservations.PUT("/:id/status", server.reservHandler.UpdateReservationStatus)
			reservations.DELETE("/:id", server.reservHandler.DeleteReservation)
		}
	}
	
	return router
}

// Placeholder handler for endpoints not yet implemented
func (server *Server) placeholderHandler(handlerName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{
			"error": "endpoint not implemented yet",
			"handler": handlerName,
		})
	}
}