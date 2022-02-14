package main

import (
	"fmt"
	"log"
	"github.com/atharva-art/bookings-project/pkg/config"
	"github.com/atharva-art/bookings-project/pkg/handlers"
	"github.com/atharva-art/bookings-project/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"
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
	if err != nil{
		log.Fatal(err)
	}

	app.Session = session
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Server listening on localhost%s\n", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
