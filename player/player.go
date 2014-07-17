package player

import (
	"github.com/gorilla/mux"
	"github.com/trapped/realmeye/base"
	"github.com/trapped/realmeye/db"
	"net/http"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	p, err := db.Default.FindPlayer(mux.Vars(req)["name"])

	b := base.Page{
		Specific: p,
	}

	if err != nil {
		tem := b.Template("player/notfound.gom")
		err = tem.Execute(w, b)
		if err != nil {
			panic(err)
		}
		return
	}

	tem := b.Template("player/index.gom")

	err = tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
