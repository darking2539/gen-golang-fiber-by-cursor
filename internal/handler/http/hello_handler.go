package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type HelloUsecase interface {
	GetHello(ctx context.Context) (string, error)
}

type HelloHandler struct {
	uc HelloUsecase
}

func NewHelloHandler(uc HelloUsecase) *HelloHandler {
	return &HelloHandler{uc: uc}
}

func (h *HelloHandler) Register(r fiber.Router) {
	r.Get("/hello-world", h.getHello)
}

// getHello godoc
// @Summary      Hello world
// @Description  Returns hello world message
// @Tags         Hello
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /hello-world [get]
func (h *HelloHandler) getHello(c *fiber.Ctx) error {
	msg, err := h.uc.GetHello(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": msg})
}
