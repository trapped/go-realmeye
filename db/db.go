package db

type Source interface {
	Open()
	Close()
	FindPlayer(name string) (*Player, error)
}

var Default Source
