package main

import (
	"log"
	"net/http"
	"os"
	"timeCardSimple/app/webapp"

	"github.com/go-chi/chi"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Creating your time card application")
	r := chi.NewRouter()

	// e := echo.New()

	// r.File("/swaggerui", "ui/index.html")
	// e.Static("/swaggerui", "ui")

	closer, repos, err := webapp.BuildRepos()
	if err != nil {
		log.Fatal("error building repos:", err)
	}
	defer closer.Close()

	app, err := webapp.BuildRoot(repos)
	if err != nil {
		log.Fatalf("error building application")
	}

	app.RegisterRoutes(r)

	log.Printf("listening on port 9080")
	http.ListenAndServe(":9080", r)

}
