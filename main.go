package main

import (
	"log"
	"timeCardSimple/app/webapp"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Printf("Creating your time card application")
	e := echo.New()

	e.File("/swaggerui", "ui/index.html")
	e.Static("/swaggerui", "ui")

	closer, repos, err := webapp.BuildRepos()
	if err != nil {
		log.Fatal("error building repos:", err)
	}
	defer closer.Close()

	app, err := webapp.BuildRoot(repos)
	if err != nil {
		log.Fatalf("error building application")
	}

	app.RegisterRoutes(e)

	log.Printf("listening on port 9080")
	e.Logger.Fatal((e.Start(":9080")))

}
