package db

import (
	"time"
)

type LastSeen struct {
	Time   string
	Server string
	Class  string
}

type Pet struct {
	Type int
}

type Outfit struct {
	Skin      int
	Dye1      int
	Dye2      int
	Accessory int
	Clothing  int
}

type ClassQuest struct {
	BestLevel int
	BestFame  int
}

type Character struct {
	Class       int
	Level       int
	Fame        int
	Exp         int
	Rank        int
	Pet         Pet
	Items       []int
	LastSeen    LastSeen
	MaxedStats  int
	Stats       []int
	Outfit      Outfit
	OutfitCount int
	Backpack    bool
}

type Player struct {
	Name            string
	Characters      []*Character
	Pets            []*Pet
	Fame            int
	FameRank        int
	Exp             int
	ExpRank         int
	Stars           int
	AccountFame     int
	AccountFameRank int
	Guild           string
	GuildRank       int
	Created         string
	LastSeen        LastSeen
	Offers          []interface{}
	Description     []string
	Similar         []string
	ClassQuests     map[int]ClassQuest
}

func (p *Player) Len() int {
	return len(p.Characters)
}

func (p *Player) Swap(i, j int) {
	p.Characters[i], p.Characters[j] = p.Characters[j], p.Characters[i]
}

func (p *Player) Less(i, j int) bool {
	if p.Characters[i].LastSeen.Time == "" && p.Characters[j].LastSeen.Time != "" {
		return false
	} else if p.Characters[i].LastSeen.Time != "" && p.Characters[j].LastSeen.Time == "" {
		return true
	} else if p.Characters[i].LastSeen.Time == "" && p.Characters[j].LastSeen.Time == "" {
		return true
	} else {

		t, _ := time.Parse("2006-01-02 15:04:05", p.Characters[i].LastSeen.Time)
		_t, _ := time.Parse("2006-01-02 15:04:05", p.Characters[j].LastSeen.Time)

		return t.After(_t)
	}
}
