package auth

import (
	"encoding/json"
	"strconv"
	auth_models "testovoe/internal/auth/models"
	"testovoe/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

type handlers struct {
	usecase Usecase
	logger  *logger.Logger
}

func NewHandlers(usecase Usecase, logger *logger.Logger) *handlers {
	return &handlers{
		usecase: usecase,
		logger:  logger,
	}
}

func (h *handlers) MapRoutes(router fiber.Router) {
	router.Post("/signin", h.Signin())
	router.Get("/confirm_email", h.ConfirmEmail())
}

func (h *handlers) ConfirmEmail() fiber.Handler {
	return func(c fiber.Ctx) error {
		mail := c.Query("mail", "")
		if mail == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		queryCode := c.Query("code", "")
		if queryCode == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		code, err := strconv.Atoi(queryCode)
		if err != nil {
			return err
		}

		if err := h.usecase.ConfirmEmail(mail, code); err != nil {
			return err
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func (h *handlers) Signin() fiber.Handler {
	return func(c fiber.Ctx) error {
		var params auth_models.SigninParams

		if err := json.Unmarshal(c.Body(), &params); err != nil {
			return err
		}

		if params.Mail == "" || params.Password == "" || params.Nickname == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		resp, err := h.usecase.Signin(&params)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(resp)
	}
}
