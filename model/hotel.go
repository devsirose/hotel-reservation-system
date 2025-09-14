package model

import (
	"database/sql"
	db "github.com/devsirose/hotel-reservation/db/sqlc"
	"github.com/google/uuid"
)

type Hotel struct {
	HotelID       uuid.UUID       `json:"hotel_id"`
	DestinationID uuid.NullUUID   `json:"destination_id"`
	TypeID        sql.NullString  `json:"type_id"`
	TotalRoom     sql.NullInt32   `json:"total_room"`
	Rating        sql.NullFloat64 `json:"rating"`
}

// ToDBModel converts model.Hotel to db.Hotel
func (h *Hotel) ToDBModel() *db.Hotel {
	return &db.Hotel{
		HotelID:       h.HotelID,
		DestinationID: h.DestinationID,
		TypeID:        h.TypeID,
		TotalRoom:     h.TotalRoom,
		Rating:        h.Rating,
	}
}

// FromDBModel converts db.Hotel to model.Hotel
func FromDBHotel(dbHotel *db.Hotel) *Hotel {
	return &Hotel{
		HotelID:       dbHotel.HotelID,
		DestinationID: dbHotel.DestinationID,
		TypeID:        dbHotel.TypeID,
		TotalRoom:     dbHotel.TotalRoom,
		Rating:        dbHotel.Rating,
	}
}