package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Printf("Creating your time card application")
	e := echo.New()

	e.File("/swaggerui", "ui/index.html")
	e.Static("/swaggerui", "ui")

	log.Printf("listening on port 9080")
	e.Logger.Fatal((e.Start(":9080")))

}
