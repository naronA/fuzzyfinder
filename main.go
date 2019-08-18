package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/naronA/fuzzyfinder/address"
	"github.com/naronA/fuzzyfinder/score"
)

func findAddress(addresses []string) echo.HandlerFunc {
	return func(c echo.Context) error {
		term := c.QueryParams()["term"]
		finders := make([]score.Finder, 0)
		for _, add := range addresses {
			f := score.Finder{Source: add, Inputs: term}
			finders = append(finders, f)
		}
		output := ""
		for _, f := range finders {
			output += f.String() + "\n"
		}
		return c.String(http.StatusOK, output)
	}
}

func main() {
	addresses := address.Load()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/address", findAddress(addresses))

	err := e.Start(":1234")
	if err != nil {
		panic(err)
	}
}
