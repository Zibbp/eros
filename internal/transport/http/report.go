package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zibbp/eros/internal/report"
	"github.com/zibbp/eros/internal/script"
	"github.com/zibbp/eros/internal/utils"
)

type ReportService interface {
	CreateReport(c echo.Context, reportDto report.Report) (*report.CreateReportResponse, error)
	GetReport(c echo.Context, id uuid.UUID) (*report.ReportResponse, error)
	GetScriptReports(c echo.Context, scriptId uuid.UUID, limit int, offset int) (report.ScriptReportsPaginationResponse, error)
}

type CreateReportRequest struct {
	Name     string             `form:"name" validate:"required"`
	Hostname string             `form:"hostname" validate:"required"`
	Status   utils.ReportStatus `form:"status" validate:"required,oneof=success failed"`
}

func (h *Handler) CreateReport(c echo.Context) error {
	crr := new(CreateReportRequest)
	if err := c.Bind(crr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(crr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	fileSrc, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer fileSrc.Close()

	reportDto := report.Report{
		Name:   crr.Name,
		Status: crr.Status,
		Script: script.Script{
			Name:     crr.Name,
			Hostname: crr.Hostname,
		},
		File:    file,
		FileSrc: fileSrc,
	}

	report, err := h.Service.ReportService.CreateReport(c, reportDto)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, report)
}

func (h *Handler) GetReport(c echo.Context) error {
	id := c.Param("id")

	uuid, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	report, err := h.Service.ReportService.GetReport(c, uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, report)
}

func (h *Handler) GetScriptReports(c echo.Context) error {
	scriptId := c.Param("id")

	uuid, err := uuid.Parse(scriptId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

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

	reports, err := h.Service.ReportService.GetScriptReports(c, uuid, pr.Limit, pr.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, reports)
}
