package home

import (
	"github.com/trapped/realmeye/base"
	"github.com/trapped/realmeye/db"
	"net/http"
)

var Watching string = ""

type home struct {
	Watching      string
	RecentChanges []db.RecentChange
}

func Serve(w http.ResponseWriter, req *http.Request) {
	b := base.Page{
		Title:       "Home",
		Location:    "/",
		Description: "Rankings, statistics, in-game trading, player and guild profiles, and more for Realm of the Mad God - the free online MMO RPG game.",
		Keywords:    "home",
		Specific: home{
			Watching:      Watching,
			RecentChanges: db.Default.RecentChanges(),
		},
	}

	tem := b.Template("home/index.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
