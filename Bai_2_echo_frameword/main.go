package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main()  {
	e:=echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK,"hell0 word")
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func getUser(c echo.Context) error  {
	id:=c.Param("id")
	return c.String(http.StatusOK,id)
}

