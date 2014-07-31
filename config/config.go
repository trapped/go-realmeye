package config

import (
	"encoding/json"
	"github.com/trapped/realmeye/db"
	"io/ioutil"
	"strings"
)

type config struct {
	Type     string
	Host     string
	Schema   string
	User     string
	Password string
	Cached   bool
}

var DB db.Source = nil

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	c := config{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	switch c.Type {
	case "mysql":
		DB = db.Source(&db.MySQL{
			Host:     strings.Split(c.Host, ":")[0],
			Port:     strings.Split(c.Host, ":")[1],
			Database: c.Schema,
			User:     c.User,
			Password: c.Password,
		})
		break
	case "bogus":
		DB = db.Source(&db.Bogus{})
		break
	default:
		panic(nil)
	}
}
