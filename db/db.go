package db

import (
	"html/template"
)

type Cache struct {
	Initialized      bool
	Players          []*Player
	PlayersByFame    []*Player
	PlayersByAccFame []*Player
	PlayersByExp     []*Player
	PlayerNames      []string
	Outfits          map[string]int
}

type RecentChange struct {
	Date    string
	Changes []template.HTML
}

const (
	SortFame    string = "fame"
	SortAccFame string = "account-fame"
	SortExp     string = "exp"
)

type Source interface {
	Open()
	Close()
	FindPlayer(name string) (*Player, error)
	SortPlayers(sorting string, offset int, num int) ([]*Player, error)
	RecentChanges() []RecentChange
}

var Default Source
