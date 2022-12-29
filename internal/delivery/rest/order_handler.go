package rest

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/model/constant"
	"go-restaurant-kelas-work/internal/tracing"
	"net/http"
)

func (h *handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()
	var request model.MenuOrderRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

	orderData, err := h.restoUsecase.CreateOrder(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetOrderInfo")
	defer span.End()
	orderID := c.Param("id")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	orderData, err := h.restoUsecase.GetOrderInfo(ctx, model.GetOrderInfoRequest{
		UserID:  userID,
		OrderID: orderID,
	})
	if err != nil {
		//exp result:
		/* SELECT * FROM "orders" WHERE "orders"."order_id" = '0070e4ff-5e53-48e7-bf2a-a20c1736533' ORDER BY "orders"."order_id" LIMIT 1
		time="2022-12-28T14:19:29+07:00" level=error msg="[delivery][rest][order_handler][GetOrderInfo] unable to get order data" error="record not found" */
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("[delivery][rest][order_handler][GetOrderInfo] unable to get order data")

		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
