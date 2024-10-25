package booking

import (
	"fmt"

	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/infrastructure/repository"
	"github.com/zakiyalmaya/hotel-management/model"
)

type bookingSvcImpl struct {
	repos *repository.Repositories
}

func NewBookingServiceImpl(repos *repository.Repositories) BookingService {
	return &bookingSvcImpl{repos: repos}
}

func (b *bookingSvcImpl) Books(booking *model.BookingEntity) error {
	// Check room availability
	room, err := b.repos.RoomRepo.GetByName(booking.RoomName)
	if err != nil {
		return err
	}

	if room.Status != constant.RoomStatusAvailable {
		return fmt.Errorf("room not available")
	}

	// Calculate total cost based on days
	duration := booking.CheckOut.Sub(booking.CheckIn).Hours() / 24
	totalCost := duration * room.Price
	booking.PaidAmount = totalCost

	err = b.repos.BookingRepo.Books(booking)
	if err != nil {
		return err
	}

	return nil
}

func (b *bookingSvcImpl) GetByRegisterNumber(registerNumber string) (*model.BookingResponse, error) {
	booking, err := b.repos.BookingRepo.GetByRegisterNumber(registerNumber)
	if err != nil {
		return nil, err
	}

	return &model.BookingResponse{
		RegisterNumber:    booking.RegisterNumber,
		GuestID:           booking.GuestID,
		GuestName:         booking.FirstName + " " + booking.LastName,
		GuestIdentity:     booking.Identity,
		RoomName:          booking.RoomName,
		RoomFloor:         booking.Floor,
		RoomType:          booking.Type,
		RoomStatus:        booking.RoomEntity.Status.Enum(),
		CheckIn:           booking.CheckIn.Format("02 January 2006"),
		CheckOut:          booking.CheckOut.Format("02 January 2006"),
		PaidAmount:        booking.PaidAmount,
		PaymentMethod:     booking.PaymentMethod.Enum(),
		PaymentStatus:     booking.PaymentStatus.Enum(),
		AdditionalRequest: booking.AdditionalRequest,
		CreatedAt:         booking.BookingEntity.CreatedAt.Format("02-01-2006 15:04:05"),
	}, nil
}

func (b *bookingSvcImpl) UpdatePayment(booking *model.BookingEntity) error {
	if err := b.repos.BookingRepo.UpdatePayment(booking); err != nil {
		return err
	}

	detail, err := b.repos.BookingRepo.GetByRegisterNumber(booking.RegisterNumber)
	if err != nil {
		return err
	}

	if booking.PaymentStatus == constant.PaymentStatusCompleted {
		if err := b.repos.RoomRepo.Update(&model.RoomEntity{
			Status: constant.RoomStatusBooked,
			Name:   detail.RoomName,
		}); err != nil {
			return err
		}

	} else if booking.PaymentStatus == constant.PaymentStatusRefunded {
		if err := b.repos.RoomRepo.Update(&model.RoomEntity{
			Status: constant.RoomStatusAvailable,
			Name:   detail.RoomName,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (b *bookingSvcImpl) Reschedule(booking *model.BookingEntity) error {
	detail, err := b.repos.BookingRepo.GetByRegisterNumber(booking.RegisterNumber)
	if err != nil {
		return err
	}

	if detail.PaymentStatus != constant.PaymentStatusPending {
		return fmt.Errorf("invalid payment status")
	}

	if detail.RoomEntity.Status != constant.RoomStatusAvailable {
		return fmt.Errorf("invalid room status")
	}

	// Calculate total cost based on days
	room, err := b.repos.RoomRepo.GetByName(detail.RoomName)
	if err != nil {
		return err
	}

	duration := booking.CheckOut.Sub(booking.CheckIn).Hours() / 24
	totalCost := duration * room.Price
	booking.PaidAmount = totalCost

	// Update date time
	if err := b.repos.BookingRepo.Reschedule(booking); err != nil {
		return err
	}

	return nil
}
