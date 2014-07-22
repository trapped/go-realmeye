package recentchanges

import (
	"github.com/trapped/realmeye/base"
	"github.com/trapped/realmeye/db"
	"net/http"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	type rc struct {
		RecentChanges []db.RecentChange
	}
	b := base.Page{
		Title:       "Recent changes",
		Location:    "/recent-changes",
		Description: "Recent changes and improvements of RealmEye.com",
		Keywords:    "recent changes",
		Specific:    db.Default.RecentChanges(),
	}

	tem := b.Template("recentchanges/index.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
