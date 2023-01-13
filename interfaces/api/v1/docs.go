package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/achmadAlli/go-simple-boilerplate/interfaces/api/utils"
	"github.com/achmadAlli/go-simple-boilerplate/internal/constants"
	"github.com/labstack/echo/v4"
)

func RegisterV1(g *echo.Group) {
	g.GET("/swagger", func(c echo.Context) error {
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
	})
}
