package home

import (
	"github.com/trapped/realmeye/base"
	"net/http"
)

var Watching string = ""

type home struct {
	Watching string
}

func Serve(w http.ResponseWriter, req *http.Request) {
	b := base.Page{
		Specific: home{Watching: Watching},
	}

	tem := b.Template("home/index.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
