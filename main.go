package main

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.org/pong3ds/play-with-echo/logger"
	"github.org/pong3ds/play-with-echo/uuid"
)

func createLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger.CreateLogger(c, uuid.NewUUID(), log.DebugLevel)
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Use(createLoggerMiddleware)
	e.GET("/", func(c echo.Context) error {
		logger.GetLogger().Debug("Debug Message")
		return c.String(http.StatusOK, "Hello, Logging World!")
	})
	e.Start(":8080")
}
