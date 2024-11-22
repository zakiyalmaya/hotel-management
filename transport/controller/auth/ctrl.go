package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/hotel-management/application/auth"
	"github.com/zakiyalmaya/hotel-management/model"
	"github.com/zakiyalmaya/hotel-management/utils"
)

type AuthController struct {
	authSvc auth.AuthService
}

func NewAuthController(authSvc auth.AuthService) *AuthController {
	return &AuthController{
		authSvc: authSvc,
	}
}

func (a *AuthController) Login(ctx *fiber.Ctx) error {
	authReq := model.AuthRequest{}

	if err := ctx.BodyParser(&authReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	if err := utils.Validator(authReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewHttpResponse(fiber.StatusBadRequest, err.Error(), nil))
	}

	authRes, err := a.authSvc.Login(&authReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", authRes))
}

func (a *AuthController) Logout(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)
	if err := a.authSvc.Logout(username); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", nil))
}

func (a *AuthController) Refresh(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)
	refreshRes, err := a.authSvc.RefreshAuthToken(username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewHttpResponse(fiber.StatusInternalServerError, err.Error(), nil))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.NewHttpResponse(fiber.StatusOK, "success", refreshRes))
}
