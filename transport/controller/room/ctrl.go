package room

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application/room"
	"github.com/zakiyalmaya/hotel-management/constant"
	"github.com/zakiyalmaya/hotel-management/model"
	"github.com/zakiyalmaya/hotel-management/utils"
)

type RoomController struct {
	roomSvc room.RoomService
}

func NewRoomController(roomSvc room.RoomService) *RoomController {
	return &RoomController{roomSvc: roomSvc}
}

func (c *RoomController) Create(ctx *fiber.Ctx) error {
	roomReq := model.CreateRoomRequest{}
	if err := ctx.BodyParser(&roomReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := utils.Validator(roomReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := c.roomSvc.Create(&model.RoomEntity{
		Name:        roomReq.Name,
		Type:        roomReq.Type,
		Price:       roomReq.Price,
		Floor:       roomReq.Floor,
		Status:      constant.RoomStatus(roomReq.Status),
		Description: roomReq.Description,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewHttpResponse(fiber.StatusCreated, "success", nil))
}

func (c *RoomController) GetByName(ctx *fiber.Ctx) error {
	name := ctx.Query("name")

	resp, err := c.roomSvc.GetByName(name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", resp))
}

func (c *RoomController) GetAll(ctx *fiber.Ctx) error {
	floor := ctx.Query("floor")
	status := ctx.Query("status")

	var request model.GetAllRoomRequest
	if floor != "" {
		request.Floor = &floor
	}

	if status != "" {
		request.Status = &status
	}

	resp, err := c.roomSvc.GetAll(&request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", resp))
}

func (c *RoomController) Update(ctx *fiber.Ctx) error {
	roomReq := model.UpdateRoomRequest{}
	if err := ctx.BodyParser(&roomReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	name := ctx.Params("name")
	roomReq.Name = name
	if err := utils.Validator(roomReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := c.roomSvc.Update(&roomReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", nil))
}