package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zibbp/eros/ent"
	"github.com/zibbp/eros/internal/script"
)

type ScriptService interface {
	CreateScript(c echo.Context, scriptDto script.Script) (*ent.Script, error)
	GetScripts(c echo.Context, limit int, offset int) (script.PaginationResponse, error)
}

type CreateScriptRequest struct {
	Name     string `json:"name" validate:"required"`
	Hostname string `json:"hostname" validate:"required"`
}

type PaginationRequest struct {
	Limit  int `query:"limit"`
	Offset int `query:"offset"`
}

func (h *Handler) CreateScript(c echo.Context) error {
	csr := new(CreateScriptRequest)
	if err := c.Bind(csr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(csr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	scriptDto := script.Script{
		Name:     csr.Name,
		Hostname: csr.Hostname,
	}

	script, err := h.Service.ScriptService.CreateScript(c, scriptDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, script)
}

func (h *Handler) GetScripts(c echo.Context) error {
	pr := new(PaginationRequest)
	if err := c.Bind(pr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if pr.Limit == 0 {
		pr.Limit = 25
	}
	if pr.Offset == 0 {
		pr.Offset = 0
	}
	if err := c.Validate(pr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	paginationResponse, err := h.Service.ScriptService.GetScripts(c, pr.Limit, pr.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, paginationResponse)
}
