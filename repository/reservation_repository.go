package repository

import (
	"context"
	"database/sql"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/google/uuid"
)

type ReservationRepository interface {
	CreateReservation(ctx context.Context, reservation *model.Reservation) error
	GetReservationByID(ctx context.Context, reservationID uuid.UUID) (*model.Reservation, error)
	ListReservationsByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Reservation, error)
	ListReservationsByRoom(ctx context.Context, roomID uuid.UUID, limit, offset int) ([]*model.Reservation, error)
	UpdateReservation(ctx context.Context, reservation *model.Reservation) error
	DeleteReservation(ctx context.Context, reservationID uuid.UUID) error
	UpdateReservationStatus(ctx context.Context, reservationID uuid.UUID, status string) error
}

type reservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) CreateReservation(ctx context.Context, reservation *model.Reservation) error {
	query := `
		INSERT INTO reservation (reservation_id, room_id, user_id, start_date, end_date, status, created_at, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		reservation.ReservationID,
		reservation.RoomID,
		reservation.UserID,
		reservation.StartDate,
		reservation.EndDate,
		reservation.Status,
		reservation.CreatedAt,
		reservation.CreatedBy,
	)
	return err
}

func (r *reservationRepository) GetReservationByID(ctx context.Context, reservationID uuid.UUID) (*model.Reservation, error) {
	var reservation model.Reservation
	query := `
		SELECT reservation_id, room_id, user_id, start_date, end_date, status, 
		       created_at, created_by, update_at, update_by
		FROM reservation
		WHERE reservation_id = $1
	`
	err := r.db.QueryRowContext(ctx, query, reservationID).Scan(
		&reservation.ReservationID,
		&reservation.RoomID,
		&reservation.UserID,
		&reservation.StartDate,
		&reservation.EndDate,
		&reservation.Status,
		&reservation.CreatedAt,
		&reservation.CreatedBy,
		&reservation.UpdateAt,
		&reservation.UpdateBy,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &reservation, nil
}

func (r *reservationRepository) ListReservationsByUser(ctx context.Context, userID string, limit, offset int) ([]*model.Reservation, error) {
	query := `
		SELECT reservation_id, room_id, user_id, start_date, end_date, status,
		       created_at, created_by, update_at, update_by
		FROM reservation
		WHERE user_id = $1
		ORDER BY start_date DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []*model.Reservation
	for rows.Next() {
		var reservation model.Reservation
		err := rows.Scan(
			&reservation.ReservationID,
			&reservation.RoomID,
			&reservation.UserID,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.Status,
			&reservation.CreatedAt,
			&reservation.CreatedBy,
			&reservation.UpdateAt,
			&reservation.UpdateBy,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, &reservation)
	}
	return reservations, nil
}

func (r *reservationRepository) ListReservationsByRoom(ctx context.Context, roomID uuid.UUID, limit, offset int) ([]*model.Reservation, error) {
	query := `
		SELECT reservation_id, room_id, user_id, start_date, end_date, status,
		       created_at, created_by, update_at, update_by
		FROM reservation
		WHERE room_id = $1
		ORDER BY start_date DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, roomID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []*model.Reservation
	for rows.Next() {
		var reservation model.Reservation
		err := rows.Scan(
			&reservation.ReservationID,
			&reservation.RoomID,
			&reservation.UserID,
			&reservation.StartDate,
			&reservation.EndDate,
			&reservation.Status,
			&reservation.CreatedAt,
			&reservation.CreatedBy,
			&reservation.UpdateAt,
			&reservation.UpdateBy,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, &reservation)
	}
	return reservations, nil
}

func (r *reservationRepository) UpdateReservation(ctx context.Context, reservation *model.Reservation) error {
	query := `
		UPDATE reservation
		SET room_id = $2, user_id = $3, start_date = $4, end_date = $5, 
		    status = $6, update_at = $7, update_by = $8
		WHERE reservation_id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		reservation.ReservationID,
		reservation.RoomID,
		reservation.UserID,
		reservation.StartDate,
		reservation.EndDate,
		reservation.Status,
		reservation.UpdateAt,
		reservation.UpdateBy,
	)
	return err
}

func (r *reservationRepository) DeleteReservation(ctx context.Context, reservationID uuid.UUID) error {
	query := `DELETE FROM reservation WHERE reservation_id = $1`
	_, err := r.db.ExecContext(ctx, query, reservationID)
	return err
}

func (r *reservationRepository) UpdateReservationStatus(ctx context.Context, reservationID uuid.UUID, status string) error {
	query := `
		UPDATE reservation
		SET status = $2, update_at = NOW()
		WHERE reservation_id = $1
	`
	_, err := r.db.ExecContext(ctx, query, reservationID, status)
	return err
}