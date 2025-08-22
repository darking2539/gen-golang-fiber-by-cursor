package http

import (
	"github.com/gofiber/fiber/v2"
	authUC "github.com/user/gen-golang-fiber-by-cursor/internal/usecase/auth"
)

type AuthHandler struct {
	uc authUC.Service
}

func NewAuthHandler(uc authUC.Service) *AuthHandler { return &AuthHandler{uc: uc} }

func (h *AuthHandler) Register(r fiber.Router) {
	r.Post("/login", h.loginHandler)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"accessToken"`
}

func (h *AuthHandler) loginHandler(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	token, err := h.uc.Login(c.Context(), req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}
	return c.JSON(loginResponse{AccessToken: token})
}
