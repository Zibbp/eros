package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eros/internal/utils"
)

type Services struct {
	ScriptService ScriptService
	ReportService ReportService
}

type Handler struct {
	Server  *echo.Echo
	Service Services
}

func NewHandler(scriptService ScriptService, reportService ReportService) *Handler {
	log.Debug().Msg("creating new handler")

	h := &Handler{
		Server: echo.New(),
		Service: Services{
			ScriptService: scriptService,
			ReportService: reportService,
		},
	}

	// Middleware
	h.Server.Validator = &utils.CustomValidator{Validator: validator.New()}

	h.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	h.mapRoutes()

	return h
}

func (h *Handler) mapRoutes() {
	log.Debug().Msg("mapping routes")

	h.Server.GET("/", func(c echo.Context) error {
		return c.String(200, "Eros API")
	})

	v1 := h.Server.Group("/api/v1")
	groupV1Routes(v1, h)
}

func groupV1Routes(e *echo.Group, h *Handler) {

	// Script
	scriptGroup := e.Group("/script")
	scriptGroup.POST("", h.CreateScript)
	scriptGroup.GET("", h.GetScripts)
	scriptGroup.GET("/:id", h.GetScript)
	scriptGroup.GET("/:id/reports", h.GetScriptReports)

	// Report
	reportGroup := e.Group("/report")
	reportGroup.POST("", h.CreateReport)
	reportGroup.GET("/:id", h.GetReport)

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.Start(":4000"); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to start server")
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := h.Server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to shutdown server")
	}

	return nil
}
