package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/3dsinteractive/govalidator"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.org/pong3ds/play-with-echo/logger"
	"github.org/pong3ds/play-with-echo/uuid"
)

type user struct {
	Username string `json:"user_name" valid:"required"`
	Password string `json:"password" valid:"alphanum,required"`
	Age      int    `json:"age" valid:"int,required"`
}

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

		model := new(user)
		if err := c.Bind(model); err != nil {
			logger.GetLogger().Error(err.Error())
			return err
		}

		isValid, err := govalidator.ValidateStruct(model)
		if err != nil {
			for _, value := range err.(govalidator.Errors) {
				e := value.(govalidator.Error)
				msg := fmt.Sprintf("%s:%s:%s", strings.ToUpper(e.Validator), e.Name, e.Err.Error())
				logger.GetLogger().Error(msg)
			}
			return err
		}

		if !isValid {
			return nil
		}

		res, err := json.Marshal(model)
		if err != nil {
			logger.GetLogger().Error(err.Error())
			return err
		}

		return c.String(http.StatusOK, string(res))
	})
	e.Start(":8080")
}
