package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/devsirose/hotel-reservation/repository"
	"github.com/google/uuid"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, reservation *model.Reservation) error
	GetReservationByID(ctx context.Context, reservationID uuid.UUID) (*model.Reservation, error)
	ListReservationsByUser(ctx context.Context, userID string, page, pageSize int) ([]*model.Reservation, error)
	ListReservationsByRoom(ctx context.Context, roomID uuid.UUID, page, pageSize int) ([]*model.Reservation, error)
	UpdateReservation(ctx context.Context, reservation *model.Reservation) error
	CancelReservation(ctx context.Context, reservationID uuid.UUID) error
	ConfirmReservation(ctx context.Context, reservationID uuid.UUID) error
}

type reservationService struct {
	reservationRepo repository.ReservationRepository
	roomRepo        repository.RoomRepository
}

func NewReservationService(reservationRepo repository.ReservationRepository, roomRepo repository.RoomRepository) ReservationService {
	return &reservationService{
		reservationRepo: reservationRepo,
		roomRepo:        roomRepo,
	}
}

func (s *reservationService) CreateReservation(ctx context.Context, reservation *model.Reservation) error {
	if reservation.ReservationID == uuid.Nil {
		reservation.ReservationID = uuid.New()
	}

	roomID := reservation.RoomID.UUID
	if !reservation.RoomID.Valid {
		return errors.New("invalid room ID")
	}

	room, err := s.roomRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return err
	}
	if room == nil {
		return errors.New("room not found")
	}

	// Fix: room.HotelID is likely a uuid.NullUUID, so use .UUID and check .Valid
	if !room.HotelID.Valid {
		return errors.New("invalid hotel ID")
	}
	// Fix: reservation.StartDate and EndDate are sql.NullTime, so use .Time and check .Valid
	if !reservation.StartDate.Valid || !reservation.EndDate.Valid {
		return errors.New("invalid reservation dates")
	}

	availableRooms, err := s.roomRepo.GetAvailableRooms(
		ctx,
		room.HotelID.UUID,
		reservation.StartDate.Time.Format(time.RFC3339),
		reservation.EndDate.Time.Format(time.RFC3339),
	)
	if err != nil {
		return err
	}
	
	isAvailable := false
	for _, availableRoom := range availableRooms {
		// availableRoom.RoomID is uuid.UUID, reservation.RoomID is uuid.NullUUID
		if reservation.RoomID.Valid && availableRoom.RoomID == reservation.RoomID.UUID {
			isAvailable = true
			break
		}
	}
	
	if !isAvailable {
		return errors.New("room is not available for the selected dates")
	}

	reservation.Status = sql.NullString{String: "PENDING", Valid: true}
	reservation.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	return s.reservationRepo.CreateReservation(ctx, reservation)
}

func (s *reservationService) GetReservationByID(ctx context.Context, reservationID uuid.UUID) (*model.Reservation, error) {
	reservation, err := s.reservationRepo.GetReservationByID(ctx, reservationID)
	if err != nil {
		return nil, err
	}
	
	if reservation == nil {
		return nil, errors.New("reservation not found")
	}
	
	return reservation, nil
}

func (s *reservationService) ListReservationsByUser(ctx context.Context, userID string, page, pageSize int) ([]*model.Reservation, error) {
	if page < 1 {
		page = 1
	}
	
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	offset := (page - 1) * pageSize
	return s.reservationRepo.ListReservationsByUser(ctx, userID, pageSize, offset)
}

func (s *reservationService) ListReservationsByRoom(ctx context.Context, roomID uuid.UUID, page, pageSize int) ([]*model.Reservation, error) {
	room, err := s.roomRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return nil, err
	}
	if room == nil {
		return nil, errors.New("room not found")
	}
	
	if page < 1 {
		page = 1
	}
	
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	offset := (page - 1) * pageSize
	return s.reservationRepo.ListReservationsByRoom(ctx, roomID, pageSize, offset)
}

func (s *reservationService) UpdateReservation(ctx context.Context, reservation *model.Reservation) error {
	existingReservation, err := s.reservationRepo.GetReservationByID(ctx, reservation.ReservationID)
	if err != nil {
		return err
	}
	
	if existingReservation == nil {
		return errors.New("reservation not found")
	}
	
	if existingReservation.Status.Valid && existingReservation.Status.String == "CANCELLED" {
		return errors.New("cannot update cancelled reservation")
	}
	
	if !reservation.StartDate.Valid || !reservation.EndDate.Valid {
		return errors.New("invalid reservation dates")
	}
	
	if reservation.StartDate.Time.After(reservation.EndDate.Time) || reservation.StartDate.Time.Equal(reservation.EndDate.Time) {
		return errors.New("invalid date range: start date must be before end date")
	}
	
	if reservation.StartDate.Time.Before(time.Now().Truncate(24 * time.Hour)) {
		return errors.New("start date cannot be in the past")
	}
	
	if !reservation.RoomID.Valid {
		return errors.New("invalid room ID")
	}
	
	room, err := s.roomRepo.GetRoomByID(ctx, reservation.RoomID.UUID)
	if err != nil {
		return err
	}
	if room == nil {
		return errors.New("room not found")
	}
	
	reservation.UpdateAt = sql.NullTime{Time: time.Now(), Valid: true}
	
	return s.reservationRepo.UpdateReservation(ctx, reservation)
}

func (s *reservationService) CancelReservation(ctx context.Context, reservationID uuid.UUID) error {
	reservation, err := s.reservationRepo.GetReservationByID(ctx, reservationID)
	if err != nil {
		return err
	}
	
	if reservation == nil {
		return errors.New("reservation not found")
	}
	
	if reservation.Status.Valid && reservation.Status.String == "CANCELLED" {
		return errors.New("reservation is already cancelled")
	}
	
	if reservation.Status.Valid && reservation.Status.String == "COMPLETED" {
		return errors.New("cannot cancel completed reservation")
	}
	
	return s.reservationRepo.UpdateReservationStatus(ctx, reservationID, "CANCELLED")
}

func (s *reservationService) ConfirmReservation(ctx context.Context, reservationID uuid.UUID) error {
	reservation, err := s.reservationRepo.GetReservationByID(ctx, reservationID)
	if err != nil {
		return err
	}
	
	if reservation == nil {
		return errors.New("reservation not found")
	}
	
	if !reservation.Status.Valid || reservation.Status.String != "PENDING" {
		return errors.New("only pending reservations can be confirmed")
	}
	
	return s.reservationRepo.UpdateReservationStatus(ctx, reservationID, "CONFIRMED")
}