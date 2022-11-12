package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	config2 "github.com/prashanth-gajula/bookings/internal/config"
	handlers2 "github.com/prashanth-gajula/bookings/internal/handlers"
	models2 "github.com/prashanth-gajula/bookings/internal/models"
	renders2 "github.com/prashanth-gajula/bookings/internal/renders"
	"net/http"
	"time"

	//"github/prashanth-gajula/go-course/pkg/config"
	//"github/prashanth-gajula/go-course/pkg/handlers"
	//"github/prashanth-gajula/go-course/pkg/renders"
	"log"
)

const PortNumber = ":8888"

var app config2.AppConfig
var session *scs.SessionManager

func main() {

	//code to tell the go that we are storing the values in the session
	gob.Register(models2.Reservation{})

	//change this to true in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	tc, err := renders2.CreateTemplateCache()
	//log.Println(tc)
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers2.NewRepo(&app)
	handlers2.NewHandlers(repo)
	renders2.NewTemplates(&app)
	//http.HandleFunc("/", handlers2.Repo.Home)
	//http.HandleFunc("/about", handlers2.Repo.About)

	fmt.Println(fmt.Sprintf("Starting Page At PortNumber %s", PortNumber))

	//_ = http.ListenAndServe(PortNumber, nil)
	serv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}
	err = serv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
