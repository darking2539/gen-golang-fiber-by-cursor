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

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

// loginHandler godoc
// @Summary      Login
// @Description  Authenticate and get access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      LoginRequest  true  "Login request"
// @Success      200   {object}  LoginResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /login [post]
func (h *AuthHandler) loginHandler(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	token, err := h.uc.Login(c.Context(), req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}
	return c.JSON(LoginResponse{AccessToken: token})
}
