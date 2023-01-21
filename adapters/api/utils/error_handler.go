package utils

import (
	"errors"
	"net/http"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/validator"
	"github.com/labstack/echo/v4"
)

func ListenAppError(e *echo.Echo) {
	e.HTTPErrorHandler = customErrorHandler
}

func customErrorHandler(err error, c echo.Context) {
	var validationErr *validator.ValidationError

	if err == nil {
		return
	}

	report, ok := err.(*echo.HTTPError)

	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, NewResponse(nil, err.Error()))
	}

	if errors.As(err, &validationErr) {
		c.JSON(http.StatusBadRequest, NewErrValidation(validationErr.ErrData, validationErr.Error()))
		return
	}

	switch err {
	case ErrBadRequest, ErrInvalidID:
		c.JSON(http.StatusBadRequest, NewResponse(nil, err.Error()))
	case ErrNotFound:
		c.JSON(http.StatusNotFound, NewResponse(nil, err.Error()))
	default:
		c.JSON(report.Code, report)
	}
}
