package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/LexusEgorov/todo/internal/models/dto"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	Access(string) error
	GetId(string) (string, error)
}

type Middleware struct {
	logger *slog.Logger
	auth   AuthService
}

func New(logger *slog.Logger /*, authService *AuthService*/) *Middleware {
	return &Middleware{
		logger: logger,
		// auth:   authService,
	}
}

func (m Middleware) WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwt := c.Request().Header.Get(echo.HeaderAuthorization)
		if err := m.auth.Access(jwt); err != nil {
			return c.JSON(echo.ErrUnauthorized.Code, dto.BadResponse{
				Message: err.Error(),
			})
		}

		return next(c)
	}
}

func (m Middleware) WithLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		timeStart := time.Now()
		err := next(c)
		if err != nil {
			m.logger.Error(err.Error())
		}

		innerLogger := m.logger.With(
			"method", c.Request().Method,
			"url", c.Request().URL,
			"duration", time.Since(timeStart))

		code := c.Response().Status
		if code >= http.StatusBadRequest && code <= http.StatusNetworkAuthenticationRequired {
			innerLogger.Error("request result", "code", code)
		} else {
			innerLogger.Info("request result", "code", code)
		}

		return err
	}
}

func (m Middleware) WithRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				m.logger.Error("recovered: %v", r)
				err := c.JSON(echo.ErrInternalServerError.Code, dto.BadResponse{
					Message: http.StatusText(echo.ErrInternalServerError.Code),
				})
				if err != nil {
					m.logger.Error("recover error", "error", err.Error())
				}
			}
		}()

		return next(c)
	}
}

func (m Middleware) WithCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwt := c.Request().Header.Get(echo.HeaderAuthorization)
		uId, err := m.auth.GetId(jwt)
		if err != nil {
			return c.JSON(echo.ErrInternalServerError.Code, dto.BadResponse{
				Message: err.Error(),
			})
		}

		if uId != c.Param("id") {
			return c.JSON(echo.ErrForbidden.Code, dto.BadResponse{
				Message: "forbidden",
			})
		}

		return next(c)
	}
}
