package repository

import (
	"context"
	"database/sql"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/google/uuid"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	GetRoomByID(ctx context.Context, roomID uuid.UUID) (*model.Room, error)
	ListRoomsByHotel(ctx context.Context, hotelID uuid.UUID, limit, offset int) ([]*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
	DeleteRoom(ctx context.Context, roomID uuid.UUID) error
	GetAvailableRooms(ctx context.Context, hotelID uuid.UUID, startDate, endDate string) ([]*model.Room, error)
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) CreateRoom(ctx context.Context, room *model.Room) error {
	query := `
		INSERT INTO room (room_id, room_name, hotel_id, floor, type_id, max_capacity, rate, description, price, created_at, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.ExecContext(ctx, query,
		room.RoomID,
		room.RoomName,
		room.HotelID,
		room.Floor,
		room.TypeID,
		room.MaxCapacity,
		room.Rate,
		room.Description,
		room.Price,
		room.CreatedAt,
		room.CreatedBy,
	)
	return err
}

func (r *roomRepository) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*model.Room, error) {
	var room model.Room
	query := `
		SELECT room_id, room_name, hotel_id, floor, type_id, max_capacity, rate, description, price, 
		       created_at, created_by, update_at, update_by
		FROM room
		WHERE room_id = $1
	`
	err := r.db.QueryRowContext(ctx, query, roomID).Scan(
		&room.RoomID,
		&room.RoomName,
		&room.HotelID,
		&room.Floor,
		&room.TypeID,
		&room.MaxCapacity,
		&room.Rate,
		&room.Description,
		&room.Price,
		&room.CreatedAt,
		&room.CreatedBy,
		&room.UpdateAt,
		&room.UpdateBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) ListRoomsByHotel(ctx context.Context, hotelID uuid.UUID, limit, offset int) ([]*model.Room, error) {
	query := `
		SELECT room_id, room_name, hotel_id, floor, type_id, max_capacity, rate, description, price,
		       created_at, created_by, update_at, update_by
		FROM room
		WHERE hotel_id = $1
		ORDER BY room_id
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, hotelID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(
			&room.RoomID,
			&room.RoomName,
			&room.HotelID,
			&room.Floor,
			&room.TypeID,
			&room.MaxCapacity,
			&room.Rate,
			&room.Description,
			&room.Price,
			&room.CreatedAt,
			&room.CreatedBy,
			&room.UpdateAt,
			&room.UpdateBy,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func (r *roomRepository) UpdateRoom(ctx context.Context, room *model.Room) error {
	query := `
		UPDATE room
		SET room_name = $2, floor = $3, type_id = $4, max_capacity = $5, 
		    rate = $6, description = $7, price = $8, update_at = $9, update_by = $10
		WHERE room_id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		room.RoomID,
		room.RoomName,
		room.Floor,
		room.TypeID,
		room.MaxCapacity,
		room.Rate,
		room.Description,
		room.Price,
		room.UpdateAt,
		room.UpdateBy,
	)
	return err
}

func (r *roomRepository) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	query := `DELETE FROM room WHERE room_id = $1`
	_, err := r.db.ExecContext(ctx, query, roomID)
	return err
}

func (r *roomRepository) GetAvailableRooms(ctx context.Context, hotelID uuid.UUID, startDate, endDate string) ([]*model.Room, error) {
	query := `
		SELECT r.room_id, r.room_name, r.hotel_id, r.floor, r.type_id, r.max_capacity, 
		       r.rate, r.description, r.price, r.created_at, r.created_by, r.update_at, r.update_by
		FROM room r
		WHERE r.hotel_id = $1
		AND r.room_id NOT IN (
			SELECT res.room_id
			FROM reservation res
			WHERE res.status != 'CANCELLED'
			AND (
				(res.start_date <= $2::TIMESTAMPTZ AND res.end_date >= $2::TIMESTAMPTZ) OR
				(res.start_date <= $3::TIMESTAMPTZ AND res.end_date >= $3::TIMESTAMPTZ) OR
				(res.start_date >= $2::TIMESTAMPTZ AND res.end_date <= $3::TIMESTAMPTZ)
			)
		)
		ORDER BY r.room_id
	`
	rows, err := r.db.QueryContext(ctx, query, hotelID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*model.Room
	for rows.Next() {
		var room model.Room
		err := rows.Scan(
			&room.RoomID,
			&room.RoomName,
			&room.HotelID,
			&room.Floor,
			&room.TypeID,
			&room.MaxCapacity,
			&room.Rate,
			&room.Description,
			&room.Price,
			&room.CreatedAt,
			&room.CreatedBy,
			&room.UpdateAt,
			&room.UpdateBy,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}