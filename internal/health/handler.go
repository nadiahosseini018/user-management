package health

import "github.com/labstack/echo/v4"

type Handler struct {
}

func NewHandler(h Handler) *Handler {
	return &h
}
func (h Handler) SetRoute(group *echo.Group) {
	group.GET("/", h.health)
}

func (h Handler) health(ctx echo.Context) error {
	return ctx.JSON(200, map[string]string{
		"health": "OK",
	})
}
