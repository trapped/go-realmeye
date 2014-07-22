package player

import (
	"github.com/gorilla/mux"
	"github.com/trapped/realmeye/base"
	"github.com/trapped/realmeye/db"
	"net/http"
	"strings"
	"unicode"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	name := strings.TrimFunc(mux.Vars(req)["name"], func(r rune) bool {
		return unicode.IsSymbol(r) || unicode.IsSpace(r) || unicode.IsMark(r)
	})
	b := base.Page{
		Title:       "Characters of the RotMG Player: " + name,
		Location:    "/player/" + name,
		Description: "Characters of the player " + name + " in Realm of the Mad God the free online mmo rpg game.",
		Keywords:    name + ", player, characters",
	}

	if len(name) == 0 {
		tem := b.Template("player/search.gom")
		err := tem.Execute(w, b)
		if err != nil {
			panic(err)
		}
		return
	}

	p, err := db.Default.FindPlayer(name)

	b.Specific = p

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
