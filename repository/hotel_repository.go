package repository

import (
	"context"
	"database/sql"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/google/uuid"
)

type HotelRepository interface {
	CreateHotel(ctx context.Context, hotel *model.Hotel) error
	GetHotelByID(ctx context.Context, hotelID uuid.UUID) (*model.Hotel, error)
	ListHotels(ctx context.Context, limit, offset int) ([]*model.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *model.Hotel) error
	DeleteHotel(ctx context.Context, hotelID uuid.UUID) error
}

type hotelRepository struct {
	db *sql.DB
}

func NewHotelRepository(db *sql.DB) HotelRepository {
	return &hotelRepository{db: db}
}

func (r *hotelRepository) CreateHotel(ctx context.Context, hotel *model.Hotel) error {
	query := `
		INSERT INTO hotel (hotel_id, destination_id, type_id, total_room, rating)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, query,
		hotel.HotelID,
		hotel.DestinationID,
		hotel.TypeID,
		hotel.TotalRoom,
		hotel.Rating,
	)
	return err
}

func (r *hotelRepository) GetHotelByID(ctx context.Context, hotelID uuid.UUID) (*model.Hotel, error) {
	var hotel model.Hotel
	query := `
		SELECT hotel_id, destination_id, type_id, total_room, rating
		FROM hotel
		WHERE hotel_id = $1
	`
	err := r.db.QueryRowContext(ctx, query, hotelID).Scan(
		&hotel.HotelID,
		&hotel.DestinationID,
		&hotel.TypeID,
		&hotel.TotalRoom,
		&hotel.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &hotel, nil
}

func (r *hotelRepository) ListHotels(ctx context.Context, limit, offset int) ([]*model.Hotel, error) {
	query := `
		SELECT hotel_id, destination_id, type_id, total_room, rating
		FROM hotel
		ORDER BY hotel_id
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []*model.Hotel
	for rows.Next() {
		var hotel model.Hotel
		err := rows.Scan(
			&hotel.HotelID,
			&hotel.DestinationID,
			&hotel.TypeID,
			&hotel.TotalRoom,
			&hotel.Rating,
		)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, &hotel)
	}
	return hotels, nil
}

func (r *hotelRepository) UpdateHotel(ctx context.Context, hotel *model.Hotel) error {
	query := `
		UPDATE hotel
		SET destination_id = $2, type_id = $3, total_room = $4, rating = $5
		WHERE hotel_id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		hotel.HotelID,
		hotel.DestinationID,
		hotel.TypeID,
		hotel.TotalRoom,
		hotel.Rating,
	)
	return err
}

func (r *hotelRepository) DeleteHotel(ctx context.Context, hotelID uuid.UUID) error {
	query := `DELETE FROM hotel WHERE hotel_id = $1`
	_, err := r.db.ExecContext(ctx, query, hotelID)
	return err
}