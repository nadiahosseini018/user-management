package user

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *Service
}

func NewHandler(h Handler) *Handler {
	return &h
}
func (h Handler) SetRoute(group *echo.Group) {
	group.POST("/", h.save)
	group.DELETE("/", h.delete)
	group.GET("/login/", h.login)

}

func (h Handler) save(ctx echo.Context) error {
	var user UserRegisterRequest
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

func (h Handler) login(ctx echo.Context) error {
	var req UserLoginRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(echo.ErrBadRequest.Code, map[string]string{
			"error": "user not found" + err.Error(),
		})
	}
	user, err := h.Service.Find(ctx.Request().Context(), req)

	if err != nil || user.Password != req.Password {
		return ctx.JSON(echo.ErrBadRequest.Code, map[string]string{
			"error": "username or password mismatch",
		})
	}

	return ctx.JSON(200, user)
}
func (h Handler) delete(ctx echo.Context) error {
	return nil
}
