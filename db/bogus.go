package db

import (
	"errors"
	"github.com/trapped/realmeye/base"
	"html/template"
)

type Bogus struct{}

func (b *Bogus) Open() {
}

func (b *Bogus) Close() {
}

func (b *Bogus) RecentChanges() []RecentChange {
	return []RecentChange{
		RecentChange{
			Date: "31 July 2014",
			Changes: []template.HTML{
				"This is a bogus news item.",
				"This is <b>another</b> bogus news item.",
			},
		},
	}
}

func (b *Bogus) FindPlayer(name string) (*Player, error) {
	if name != "bogus" {
		p := Player{
			Name: name,
			Similar: base.Similars(name, []string{
				"tset",
				"teste",
				"tester",
				"dadad",
			}, 10, false),
		}
		return &p, errors.New("Player not found")
	}
	pets := make(map[int]Pet)
	pets[1] = Pet{
		Id:   1,
		Type: 6000,
	}
	p := Player{
		Name:        name,
		Fame:        767,
		FameRank:    1,
		Exp:         15989,
		ExpRank:     1,
		Stars:       57,
		AccountFame: 21,
		Guild:       "MAFIA",
		GuildRank:   40,
		Created:     "2014-07-15 08:07:23",
		LastSeen: LastSeen{
			Time:   "2014-07-15 08:08:57",
			Server: "EUWest",
			Class:  "Paladin",
		},
		Description: []string{
			"This is a bogus player",
			"with a bogus description",
			"and bogus stats.",
		},
		Pets: pets,
		Characters: []*Character{
			&Character{
				Class:      782,
				Level:      20,
				Fame:       1523,
				Exp:        20000,
				Rank:       133,
				Pet:        1,
				Items:      []int{},
				LastSeen:   LastSeen{},
				MaxedStats: 3,
				Stats:      []int{},
				Outfit:     Outfit{},
			},
		},
	}
	return &p, nil
}
