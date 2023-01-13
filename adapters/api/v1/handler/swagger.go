package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/achmadAlli/go-simple-boilerplate/adapters/api/utils"
	"github.com/achmadAlli/go-simple-boilerplate/internal/constants"
	"github.com/labstack/echo/v4"
)

type SwaggerHandler struct {
}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (h SwaggerHandler) Swagger(c echo.Context) error {

	reader, err := os.Open("./docs/swagger/app.v1.swagger.json")

	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewResponse(nil, "api docs not found"))
	}

	defer reader.Close()

	bytesDocs, err := ioutil.ReadAll(reader)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(nil, constants.ErrInternalServiceError.Error()))
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(bytesDocs), &result)

	return c.JSON(http.StatusOK, result)
}
