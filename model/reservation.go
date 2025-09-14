package model

import (
	"database/sql"
	db "github.com/devsirose/hotel-reservation/db/sqlc"
	"github.com/google/uuid"
)

type Reservation struct {
	ReservationID uuid.UUID      `json:"reservation_id"`
	RoomID        uuid.NullUUID  `json:"room_id"`
	UserID        sql.NullString `json:"user_id"`
	StartDate     sql.NullTime   `json:"start_date"`
	EndDate       sql.NullTime   `json:"end_date"`
	Status        sql.NullString `json:"status"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	CreatedBy     uuid.NullUUID  `json:"created_by"`
	UpdateAt      sql.NullTime   `json:"update_at"`
	UpdateBy      uuid.NullUUID  `json:"update_by"`
}

// ToDBModel converts model.Reservation to db.Reservation
func (r *Reservation) ToDBModel() *db.Reservation {
	return &db.Reservation{
		ReservationID: r.ReservationID,
		RoomID:        r.RoomID,
		UserID:        r.UserID,
		StartDate:     r.StartDate,
		EndDate:       r.EndDate,
		Status:        r.Status,
		CreatedAt:     r.CreatedAt,
		CreatedBy:     r.CreatedBy,
		UpdateAt:      r.UpdateAt,
		UpdateBy:      r.UpdateBy,
	}
}

// FromDBReservation converts db.Reservation to model.Reservation
func FromDBReservation(dbReservation *db.Reservation) *Reservation {
	return &Reservation{
		ReservationID: dbReservation.ReservationID,
		RoomID:        dbReservation.RoomID,
		UserID:        dbReservation.UserID,
		StartDate:     dbReservation.StartDate,
		EndDate:       dbReservation.EndDate,
		Status:        dbReservation.Status,
		CreatedAt:     dbReservation.CreatedAt,
		CreatedBy:     dbReservation.CreatedBy,
		UpdateAt:      dbReservation.UpdateAt,
		UpdateBy:      dbReservation.UpdateBy,
	}
}