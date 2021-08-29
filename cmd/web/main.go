package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fournelf2/myproject/pkg/config"
	"github.com/fournelf2/myproject/pkg/handlers"
	"github.com/fournelf2/myproject/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHanglers(repo)
	render.NewTemplates(&app)

	test1 := fmt.Sprintf("Starting application on port %s", portNumber)
	fmt.Print(test1)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
