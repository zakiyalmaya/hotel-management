package booking

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/hotel-management/model"
)

type bookingRepoImpl struct {
	db *sqlx.DB
}

func NewBookingRepository(db *sqlx.DB) BookingRepository {
	return &bookingRepoImpl{db: db}
}

func (b *bookingRepoImpl) Books(booking *model.BookingEntity) error {
	query := "INSERT INTO bookings (register_number, guest_id, room_name, check_in, check_out, paid_amount, payment_method, payment_status, additional_request) VALUES (:register_number, :guest_id, :room_name, :check_in, :check_out, :paid_amount, :payment_method, :payment_status, :additional_request)"
	_, err := b.db.NamedExec(query, booking)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (b *bookingRepoImpl) GetByRegisterNumber(registerNumber string) (*model.BookingDetail, error) {
	booking := &model.BookingDetail{}
	query := "SELECT b.register_number, b.guest_id, g.first_name, g.last_name, g.identity_number, b.room_name, r.floor, r.type, r.status, b.check_in, b.check_out, b.paid_amount, b.payment_method, b.payment_status, b.additional_request, b.created_at FROM bookings b JOIN guests g ON b.guest_id = g.id JOIN rooms r ON r.name = b.room_name WHERE register_number = ?"

	err := b.db.Get(booking, query, registerNumber)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return booking, nil
}

func (b *bookingRepoImpl) UpdatePayment(booking *model.BookingEntity) error {
	query := "UPDATE bookings SET payment_status = :payment_status WHERE register_number = :register_number"
	_, err := b.db.NamedExec(query, booking)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}

func (b *bookingRepoImpl) Reschedule(booking *model.BookingEntity) error {
	query := "UPDATE bookings SET check_in = :check_in, check_out = :check_out, paid_amount = :paid_amount WHERE register_number = :register_number"
	_, err := b.db.NamedExec(query, booking)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return err
	}

	return nil
}