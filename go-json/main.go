package main

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
)

type JSONSerializer struct{}

func (j *JSONSerializer) Serialize(c echo.Context, i interface{}, _ string) error {
	return json.NewEncoder(c.Response()).Encode(i)
}

func (j *JSONSerializer) Deserialize(c echo.Context, i interface{}) error {
	return json.NewDecoder(c.Request().Body).Decode(i)
}

func main() {
	e := echo.New()
	e.JSONSerializer = &JSONSerializer{}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok!")
	})
	e.Start(":8080")
}
