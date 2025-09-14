package model

import (
	"database/sql"
	db "github.com/devsirose/hotel-reservation/db/sqlc"
	"github.com/google/uuid"
)

type Room struct {
	RoomID      uuid.UUID       `json:"room_id"`
	RoomName    sql.NullString  `json:"room_name"`
	HotelID     uuid.NullUUID   `json:"hotel_id"`
	Floor       sql.NullInt32   `json:"floor"`
	TypeID      sql.NullString  `json:"type_id"`
	MaxCapacity sql.NullInt32   `json:"max_capacity"`
	Rate        sql.NullFloat64 `json:"rate"`
	Description sql.NullString  `json:"description"`
	Price       sql.NullInt32   `json:"price"`
	CreatedAt   sql.NullTime    `json:"created_at"`
	CreatedBy   uuid.NullUUID   `json:"created_by"`
	UpdateAt    sql.NullTime    `json:"update_at"`
	UpdateBy    uuid.NullUUID   `json:"update_by"`
}

// ToDBModel converts model.Room to db.Room
func (r *Room) ToDBModel() *db.Room {
	return &db.Room{
		RoomID:      r.RoomID,
		RoomName:    r.RoomName,
		HotelID:     r.HotelID,
		Floor:       r.Floor,
		TypeID:      r.TypeID,
		MaxCapacity: r.MaxCapacity,
		Rate:        r.Rate,
		Description: r.Description,
		Price:       r.Price,
		CreatedAt:   r.CreatedAt,
		CreatedBy:   r.CreatedBy,
		UpdateAt:    r.UpdateAt,
		UpdateBy:    r.UpdateBy,
	}
}

// FromDBRoom converts db.Room to model.Room
func FromDBRoom(dbRoom *db.Room) *Room {
	return &Room{
		RoomID:      dbRoom.RoomID,
		RoomName:    dbRoom.RoomName,
		HotelID:     dbRoom.HotelID,
		Floor:       dbRoom.Floor,
		TypeID:      dbRoom.TypeID,
		MaxCapacity: dbRoom.MaxCapacity,
		Rate:        dbRoom.Rate,
		Description: dbRoom.Description,
		Price:       dbRoom.Price,
		CreatedAt:   dbRoom.CreatedAt,
		CreatedBy:   dbRoom.CreatedBy,
		UpdateAt:    dbRoom.UpdateAt,
		UpdateBy:    dbRoom.UpdateBy,
	}
}