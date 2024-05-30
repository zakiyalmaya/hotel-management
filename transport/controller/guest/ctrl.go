package guest

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application/guest"
	"github.com/zakiyalmaya/hotel-management/model"
	"github.com/zakiyalmaya/hotel-management/utils"
)

type GuestController struct {
	guestSvc guest.Service
}

func NewGuestController(guestSvc guest.Service) *GuestController {
	return &GuestController{guestSvc: guestSvc}
}

func (c *GuestController) Create(ctx *fiber.Ctx) error {
	guestReq := model.CreateGuestRequest{}
	if err := ctx.BodyParser(&guestReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := utils.Validator(guestReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	parsedDateOfBirth, err := time.Parse("02-01-2006", guestReq.DateOfBirth)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := c.guestSvc.Create(&model.GuestEntity{
		FirstName:   guestReq.FirstName,
		LastName:    guestReq.LastName,
		Identity:    guestReq.Identity,
		DateOfBirth: parsedDateOfBirth,
		Email:       guestReq.Email,
		PhoneNumber: guestReq.PhoneNumber,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}

func (c *GuestController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	resp, err := c.guestSvc.GetByID(idInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": resp})
}
