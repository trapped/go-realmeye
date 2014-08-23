package topplayers

import (
	"github.com/gorilla/mux"
	"github.com/trapped/realmeye/base"
	"github.com/trapped/realmeye/db"
	"net/http"
	"strconv"
)

func Serve(w http.ResponseWriter, req *http.Request) {
	sorting := mux.Vars(req)["sorting"]
	offset := 0
	if n, err := strconv.Atoi(mux.Vars(req)["offset"]); err == nil {
		offset = n - 1
	}
	b := base.Page{
		Title:    "Top RotMG Players by " + base.Capitalize(sorting) + " 1-100",
		Location: "/top-players-by-" + sorting,
		//Description: "Characters of the player " + name + " in Realm of the Mad God the free online mmo rpg game.",
		//Keywords:    name + ", player, characters",
	}

	tem := b.Template("topplayers/bystat.gom")

	pl, err := db.Default.SortPlayers(sorting, 0, 1000)
	if err != nil {
		panic(err)
	}

	num := len(pl)

	pgs := []int{}
	for i := 1; i <= len(pl)/100; i++ {
		pgs = append(pgs, i*100)
	}
	if len(pl)%100 > 0 {
		lt := 0
		if len(pgs) > 0 {
			lt = pgs[len(pgs)-1]
		}
		pgs = append(pgs, lt+len(pl)%100)
	}

	if len(pl) < offset {
		offset = 0
	}
	pl = pl[offset:]
	if len(pl) > 100 {
		pl = pl[:100]
	}

	b.Specific = struct {
		Offset  int
		Pages   []int
		Players []*db.Player
		Number  int
		Sorting string
	}{
		Offset:  offset + 1,
		Players: pl,
		Pages:   pgs,
		Number:  num,
		Sorting: sorting,
	}

	err = tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
