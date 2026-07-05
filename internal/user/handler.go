package user

import "github.com/labstack/echo/v4"

type Handler struct {
	Service *Service
}

func NewHandler(h Handler) *Handler {
	return &h
}
func (h Handler) SetRoute(group *echo.Group) {
	group.POST("/", h.save)
	group.DELETE("/", h.delete)
}

func (h Handler) save(ctx echo.Context) error {
	var user User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(echo.ErrBadRequest.Code, map[string]string{
			"error": "Invalid body: " + err.Error(),
		})
	}

	saveUser, err := h.Service.Save(ctx.Request().Context(), user)
	if err != nil {
		return ctx.JSON(echo.ErrBadRequest.Code, map[string]string{
			"error": "Cant save: " + err.Error(),
		})
	}

	return ctx.JSON(200, saveUser)
}
func (h Handler) delete(ctx echo.Context) error {
	return nil
}
