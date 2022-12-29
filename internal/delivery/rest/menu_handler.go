package rest

import (
	"github.com/labstack/echo/v4"
	"go-restaurant-kelas-work/internal/tracing"
	"net/http"
)

func (h *handler) GetMenuList(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetMenuList")
	defer span.End()
	typeMenu := c.FormValue("type")
	menu, err := h.restoUsecase.GetMenuList(ctx, typeMenu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menu,
	})
}
