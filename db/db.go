package db

type Cache struct {
	Initialized bool
	Players     []*Player
	PlayerNames []string
	Outfits     map[string]int
}

type Source interface {
	Open()
	Close()
	FindPlayer(name string) (*Player, error)
}

var Default Source
