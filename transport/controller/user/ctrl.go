package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application/user"
	"github.com/zakiyalmaya/hotel-management/model"
	"github.com/zakiyalmaya/hotel-management/utils"
)

type UserController struct {
	userSvc user.UserService
}

func NewUserController(userSvc user.UserService) *UserController {
	return &UserController{userSvc: userSvc}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	createUserReq := model.CreateUserRequest{}

	if err := ctx.BodyParser(&createUserReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := utils.Validator(createUserReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := c.userSvc.Create(&createUserReq); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewHttpResponse(fiber.StatusCreated, "success", nil))
}

func (c *UserController) ChangePassword(ctx *fiber.Ctx) error {
	changePasswordReq := model.ChangePasswordRequest{}
	if err := ctx.BodyParser(&changePasswordReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	changePasswordReq.Username= ctx.Locals("username").(string)
	if err := utils.Validator(changePasswordReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	err := c.userSvc.ChangePassword(&changePasswordReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", nil))
}
	