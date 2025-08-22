package http

import (
	"github.com/gofiber/fiber/v2"
	profileUC "github.com/user/gen-golang-fiber-by-cursor/internal/usecase/profile"
)

type ProfileHandler struct {
	uc profileUC.Service
}

func NewProfileHandler(uc profileUC.Service) *ProfileHandler { return &ProfileHandler{uc: uc} }

func (h *ProfileHandler) Register(r fiber.Router) {
	r.Get("/profile", h.getProfile)
}

func (h *ProfileHandler) getProfile(c *fiber.Ctx) error {
	usernameVal := c.Locals("username")
	username, _ := usernameVal.(string)
	if username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	prof, err := h.uc.GetByUsername(c.Context(), username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(prof)
}
