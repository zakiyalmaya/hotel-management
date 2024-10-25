package booking

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zakiyalmaya/hotel-management/application/booking"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/model"
	"github.com/zakiyalmaya/hotel-management/utils"
)

type BookingController struct {
	bookingSvc booking.BookingService
}

func NewBookingController(bookingSvc booking.BookingService) *BookingController {
	return &BookingController{bookingSvc: bookingSvc}
}

func (c *BookingController) Books(ctx *fiber.Ctx) error {
	bookingReq := model.BookingRequest{}
	if err := ctx.BodyParser(&bookingReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := utils.Validator(bookingReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if isMethodValid := constant.PaymentMethod(bookingReq.PaymentMethod).Validation(); !isMethodValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, "invalid payment method", nil))
	}

	parseCheckIn, err := time.Parse("02-01-2006", bookingReq.CheckIn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	parseCheckOut, err := time.Parse("02-01-2006", bookingReq.CheckOut)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if parseCheckIn.After(parseCheckOut) {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, "check out date must be greater than check in date", nil))
	}

	registerNumber := uuid.New().String()
	if err := c.bookingSvc.Books(&model.BookingEntity{
		RegisterNumber:    registerNumber,
		GuestID:           bookingReq.GuestID,
		RoomName:          bookingReq.RoomName,
		CheckIn:           parseCheckIn,
		CheckOut:          parseCheckOut,
		PaymentMethod:     constant.PaymentMethod(bookingReq.PaymentMethod),
		PaymentStatus:     constant.PaymentStatusPending,
		AdditionalRequest: &bookingReq.AdditionRequest,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewHttpResponse(fiber.StatusCreated, "success", map[string]string{"register_number": registerNumber}))
}

func (c *BookingController) GetByRegisterNumber(ctx *fiber.Ctx) error {
	registerNumber := ctx.Query("register_number")
	if registerNumber == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, "register number is required", nil))
	}

	resp, err := c.bookingSvc.GetByRegisterNumber(registerNumber)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", resp))
}

func (c *BookingController) UpdatePayment(ctx *fiber.Ctx) error {
	updatePaymentReq := model.UpdatePaymentRequest{}
	if err := ctx.BodyParser(&updatePaymentReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := utils.Validator(updatePaymentReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if isStatusValid := constant.PaymentMethod(updatePaymentReq.PaymentStatus).Validation(); !isStatusValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, "invalid payment status", nil))
	}

	if err := c.bookingSvc.UpdatePayment(&model.BookingEntity{
		RegisterNumber: updatePaymentReq.RegisterNumber,
		PaymentStatus:  constant.PaymentStatus(updatePaymentReq.PaymentStatus),
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", nil))
}

func (c *BookingController) Reschedule(ctx *fiber.Ctx) error {
	rescheduleReq := model.ResceduleRequest{}
	if err := ctx.BodyParser(&rescheduleReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := utils.Validator(rescheduleReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	parseCheckIn, err := time.Parse("02-01-2006", rescheduleReq.CheckIn)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	parseCheckOut, err := time.Parse("02-01-2006", rescheduleReq.CheckOut)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := c.bookingSvc.Reschedule(&model.BookingEntity{
		RegisterNumber: rescheduleReq.RegisterNumber,
		CheckIn:        parseCheckIn,
		CheckOut:       parseCheckOut,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", nil))
}
