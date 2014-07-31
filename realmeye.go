package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/trapped/realmeye/base"
	"github.com/trapped/realmeye/config"
	"github.com/trapped/realmeye/db"
	"github.com/trapped/realmeye/home"
	"github.com/trapped/realmeye/player"
	"github.com/trapped/realmeye/recentchanges"
	"net/http"
	"os"
	"path/filepath"
)

func registerHandlers(m *mux.Router) {
	m.NotFoundHandler = http.HandlerFunc(base.NotFound)
	m.HandleFunc("/", home.Serve)
	//Player search
	m.HandleFunc("/player", player.Serve)
	//Player profile
	m.HandleFunc("/player/{name}", player.Serve)
	//m.HandleFunc("/pets-of/{name}", petsOf)
	//m.HandleFunc("/graveyard-of-player/{name}", graveyardOf)
	//Guilds
	//m.HandleFunc("/guild/{name}", guild)
	//m.HandleFunc("/top-characters-of-guild/{name}", topCharsOf)
	//m.HandleFunc("/top-pets-of-guild/{name}", topPetsOf)
	//Recent changes
	m.HandleFunc("/recent-changes", recentchanges.Serve)
}

func main() {
	cwd, _ := os.Getwd()
	config.Load(filepath.Join(cwd, "./config.json"))
	db.Default = config.DB
	db.Default.Open()
	defer db.Default.Close()

	m := mux.NewRouter()

	registerHandlers(m)

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.Use(negroni.NewStatic(http.Dir("static")))
	n.UseHandler(m)
	n.Run(":3001")
}
