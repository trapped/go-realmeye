package db

type LastSeen struct {
	Time   string
	Server string
	Class  string
}

type Item struct {
	Type int
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
	Items       map[int]Item
	LastSeen    LastSeen
	MaxedStats  int
	Stats       []int
	Outfit      Outfit
	OutfitCount int
	Backpack    bool
}

type Player struct {
	Name            string
	Characters      []Character
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
