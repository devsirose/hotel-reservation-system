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

type RoomService interface {
	CreateRoom(ctx context.Context, room *model.Room) error
	GetRoomByID(ctx context.Context, roomID uuid.UUID) (*model.Room, error)
	ListRoomsByHotel(ctx context.Context, hotelID uuid.UUID, page, pageSize int) ([]*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) error
	DeleteRoom(ctx context.Context, roomID uuid.UUID) error
	GetAvailableRooms(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*model.Room, error)
}

type roomService struct {
	roomRepo  repository.RoomRepository
	hotelRepo repository.HotelRepository
}

func NewRoomService(roomRepo repository.RoomRepository, hotelRepo repository.HotelRepository) RoomService {
	return &roomService{
		roomRepo:  roomRepo,
		hotelRepo: hotelRepo,
	}
}

func (s *roomService) CreateRoom(ctx context.Context, room *model.Room) error {
	if room.RoomID == uuid.Nil {
		room.RoomID = uuid.New()
	}
	
	if !room.HotelID.Valid {
		return errors.New("invalid hotel ID")
	}
	
	hotel, err := s.hotelRepo.GetHotelByID(ctx, room.HotelID.UUID)
	if err != nil {
		return err
	}
	if hotel == nil {
		return errors.New("hotel not found")
	}
	
	if room.MaxCapacity.Valid && room.MaxCapacity.Int32 <= 0 {
		return errors.New("max capacity must be greater than 0")
	}
	
	if room.Price.Valid && room.Price.Int32 < 0 {
		return errors.New("price cannot be negative")
	}
	
	if room.Rate.Valid && (room.Rate.Float64 < 0 || room.Rate.Float64 > 5) {
		return errors.New("rate must be between 0 and 5")
	}
	
	room.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	
	return s.roomRepo.CreateRoom(ctx, room)
}

func (s *roomService) GetRoomByID(ctx context.Context, roomID uuid.UUID) (*model.Room, error) {
	room, err := s.roomRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return nil, err
	}
	
	if room == nil {
		return nil, errors.New("room not found")
	}
	
	return room, nil
}

func (s *roomService) ListRoomsByHotel(ctx context.Context, hotelID uuid.UUID, page, pageSize int) ([]*model.Room, error) {
	hotel, err := s.hotelRepo.GetHotelByID(ctx, hotelID)
	if err != nil {
		return nil, err
	}
	if hotel == nil {
		return nil, errors.New("hotel not found")
	}
	
	if page < 1 {
		page = 1
	}
	
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	offset := (page - 1) * pageSize
	return s.roomRepo.ListRoomsByHotel(ctx, hotelID, pageSize, offset)
}

func (s *roomService) UpdateRoom(ctx context.Context, room *model.Room) error {
	existingRoom, err := s.roomRepo.GetRoomByID(ctx, room.RoomID)
	if err != nil {
		return err
	}
	
	if existingRoom == nil {
		return errors.New("room not found")
	}
	
	if room.MaxCapacity.Valid && room.MaxCapacity.Int32 <= 0 {
		return errors.New("max capacity must be greater than 0")
	}
	
	if room.Price.Valid && room.Price.Int32 < 0 {
		return errors.New("price cannot be negative")
	}
	
	if room.Rate.Valid && (room.Rate.Float64 < 0 || room.Rate.Float64 > 5) {
		return errors.New("rate must be between 0 and 5")
	}
	
	room.UpdateAt = sql.NullTime{Time: time.Now(), Valid: true}
	
	return s.roomRepo.UpdateRoom(ctx, room)
}

func (s *roomService) DeleteRoom(ctx context.Context, roomID uuid.UUID) error {
	existingRoom, err := s.roomRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return err
	}
	
	if existingRoom == nil {
		return errors.New("room not found")
	}
	
	return s.roomRepo.DeleteRoom(ctx, roomID)
}

func (s *roomService) GetAvailableRooms(ctx context.Context, hotelID uuid.UUID, checkIn, checkOut time.Time) ([]*model.Room, error) {
	if checkIn.After(checkOut) || checkIn.Equal(checkOut) {
		return nil, errors.New("invalid date range: check-in must be before check-out")
	}
	
	if checkIn.Before(time.Now().Truncate(24 * time.Hour)) {
		return nil, errors.New("check-in date cannot be in the past")
	}
	
	return s.roomRepo.GetAvailableRooms(ctx, hotelID, checkIn.Format(time.RFC3339), checkOut.Format(time.RFC3339))
}