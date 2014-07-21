package home

import (
	"fmt"
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
	fmt.Println(db.Default.RecentChanges())
	b := base.Page{
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
