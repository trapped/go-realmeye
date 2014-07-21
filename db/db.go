package db

import (
	"html/template"
)

type Cache struct {
	Initialized bool
	Players     []*Player
	PlayerNames []string
	Outfits     map[string]int
}

type RecentChange struct {
	Date    string
	Changes []template.HTML
}

type Source interface {
	Open()
	Close()
	FindPlayer(name string) (*Player, error)
	RecentChanges() []RecentChange
}

var Default Source
