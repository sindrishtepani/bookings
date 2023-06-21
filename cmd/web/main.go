package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sindrishtepani/bookings/pkg/config"
	"github.com/sindrishtepani/bookings/pkg/handlers"
	"github.com/sindrishtepani/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.Session = session
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)
	fmt.Print("Starting application on port", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
