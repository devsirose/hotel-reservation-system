package service

import (
	"context"
	"errors"

	"github.com/devsirose/hotel-reservation/model"
	"github.com/devsirose/hotel-reservation/repository"
	"github.com/google/uuid"
)

type HotelService interface {
	CreateHotel(ctx context.Context, hotel *model.Hotel) error
	GetHotelByID(ctx context.Context, hotelID uuid.UUID) (*model.Hotel, error)
	ListHotels(ctx context.Context, page, pageSize int) ([]*model.Hotel, error)
	UpdateHotel(ctx context.Context, hotel *model.Hotel) error
	DeleteHotel(ctx context.Context, hotelID uuid.UUID) error
}

type hotelService struct {
	hotelRepo repository.HotelRepository
}

func NewHotelService(hotelRepo repository.HotelRepository) HotelService {
	return &hotelService{
		hotelRepo: hotelRepo,
	}
}

func (s *hotelService) CreateHotel(ctx context.Context, hotel *model.Hotel) error {
	if hotel.HotelID == uuid.Nil {
		hotel.HotelID = uuid.New()
	}
	
	if hotel.TotalRoom.Valid && hotel.TotalRoom.Int32 < 0 {
		return errors.New("total room cannot be negative")
	}
	
	if hotel.Rating.Valid && (hotel.Rating.Float64 < 0 || hotel.Rating.Float64 > 5) {
		return errors.New("rating must be between 0 and 5")
	}
	
	return s.hotelRepo.CreateHotel(ctx, hotel)
}

func (s *hotelService) GetHotelByID(ctx context.Context, hotelID uuid.UUID) (*model.Hotel, error) {
	hotel, err := s.hotelRepo.GetHotelByID(ctx, hotelID)
	if err != nil {
		return nil, err
	}
	
	if hotel == nil {
		return nil, errors.New("hotel not found")
	}
	
	return hotel, nil
}

func (s *hotelService) ListHotels(ctx context.Context, page, pageSize int) ([]*model.Hotel, error) {
	if page < 1 {
		page = 1
	}
	
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	offset := (page - 1) * pageSize
	return s.hotelRepo.ListHotels(ctx, pageSize, offset)
}

func (s *hotelService) UpdateHotel(ctx context.Context, hotel *model.Hotel) error {
	existingHotel, err := s.hotelRepo.GetHotelByID(ctx, hotel.HotelID)
	if err != nil {
		return err
	}
	
	if existingHotel == nil {
		return errors.New("hotel not found")
	}
	
	if hotel.TotalRoom.Valid && hotel.TotalRoom.Int32 < 0 {
		return errors.New("total room cannot be negative")
	}
	
	if hotel.Rating.Valid && (hotel.Rating.Float64 < 0 || hotel.Rating.Float64 > 5) {
		return errors.New("rating must be between 0 and 5")
	}
	
	return s.hotelRepo.UpdateHotel(ctx, hotel)
}

func (s *hotelService) DeleteHotel(ctx context.Context, hotelID uuid.UUID) error {
	existingHotel, err := s.hotelRepo.GetHotelByID(ctx, hotelID)
	if err != nil {
		return err
	}
	
	if existingHotel == nil {
		return errors.New("hotel not found")
	}
	
	return s.hotelRepo.DeleteHotel(ctx, hotelID)
}