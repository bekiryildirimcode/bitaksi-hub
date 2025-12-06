package handlers

import (
	"net/http"
	"strconv"

	"github.com/bekiryildirimcode/driver-service/model"
	"github.com/labstack/echo/v4"
)

type DriverHandler struct {
	service model.DriverService
}

func NewDriverHandler(service model.DriverService) *DriverHandler {
	return &DriverHandler{service: service}
}

func (h *DriverHandler) GetAll(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.ParseInt(c.QueryParam("pageSize"), 10, 64)
	if err != nil {
		pageSize = 10
	}
	drivers, err := h.service.ListDrivers(int(page), int(pageSize))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, drivers)

}

func (h *DriverHandler) Create(c echo.Context) error {
	var req model.CreateDriverRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(400, echo.Map{"validationError": err.Error()})
	}
	id, err := h.service.CreateDriver(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]string{"id": id})
}

func (h *DriverHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var req model.CreateDriverRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}
	if err := c.Validate(&req); err != nil {
		return c.JSON(400, echo.Map{"validationError": err.Error()})
	}
	err := h.service.UpdateDriver(id, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"id": id})
}

func (h *DriverHandler) GetNearby(c echo.Context) error {
	var req model.NearbyDriverQuery

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"validationError": err.Error()})
	}

	drivers, err := h.service.FindNearby(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, drivers)
}
