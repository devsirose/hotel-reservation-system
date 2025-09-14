package api

import (
	"database/sql"

	db "github.com/devsirose/hotel-reservation/db/sqlc"
	"github.com/devsirose/hotel-reservation/handler"
	"github.com/devsirose/hotel-reservation/repository"
	"github.com/devsirose/hotel-reservation/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store         db.Store
	router        *gin.Engine
	hotelHandler  *handler.HotelHandler
	roomHandler   *handler.RoomHandler
	reservHandler *handler.ReservationHandler
}

func NewServer(store db.Store, sqlDB *sql.DB) *Server {
	// Initialize repositories
	hotelRepo := repository.NewHotelRepository(sqlDB)
	roomRepo := repository.NewRoomRepository(sqlDB)
	reservationRepo := repository.NewReservationRepository(sqlDB)
	
	// Initialize services
	hotelService := service.NewHotelService(hotelRepo)
	roomService := service.NewRoomService(roomRepo, hotelRepo)
	reservationService := service.NewReservationService(reservationRepo, roomRepo)
	
	// Initialize handlers
	hotelHandler := handler.NewHotelHandler(hotelService)
	roomHandler := handler.NewRoomHandler(roomService)
	reservHandler := handler.NewReservationHandler(reservationService)

	server := &Server{
		store:         store,
		hotelHandler:  hotelHandler,
		roomHandler:   roomHandler,
		reservHandler: reservHandler,
	}

	// Setup routes
	router := server.setupRoutes()
	
	//custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("currency", validCurrency)
	}

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Close() error {
	return server.Close()
}

func errorResponse(err error) gin.H {
	return gin.H{
		//custom err response here
		"error": err.Error(),
	}
}

func notFoundResponse(obj any) gin.H {
	return gin.H{
		"error": obj.(string) + " not found",
	}
}

func validCurrency(fl validator.FieldLevel) bool {
	// Placeholder validator - implement currency validation logic here
	currency := fl.Field().String()
	return currency != ""
}
