package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/trapped/realmeye/base"
	"html/template"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

//updates done through the web interface must both update the data in the cache AND queue
//a db update

type MySQL struct {
	Host       string
	Port       string
	User       string
	Password   string
	Database   string
	Cached     bool
	Cache      Cache
	Connection *sql.DB
}

func (m *MySQL) Open() {
	if len(m.Host) == 0 {
		m.Host = "127.0.0.1"
	}
	if len(m.Port) == 0 {
		m.Port = "3306"
	}
	if len(m.User) == 0 {
		m.User = "root"
	}
	if len(m.Database) == 0 {
		m.Database = "realmeye"
	}
	conn, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", m.User, m.Password, m.Host, m.Port, m.Database))
	if err != nil {
		panic(err)
	}
	m.Connection = conn
	go func() {
		for {
			m.cache_players()
			time.Sleep(time.Minute)
		}
	}()
}

func (m *MySQL) Close() {
	err := m.Connection.Close()
	if err != nil {
		panic(err)
	}
}

func (m *MySQL) cache_players() {
	fmt.Println("[DBCACHE] Caching players...")

	//prevent/catch crashes
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[DBCACHE] Unable to finish caching; recovered from/crashed in", r)
			m.Cache.Initialized = false
			m.Cache = Cache{}
		}
	}()

	//temporary storage
	temp := make(map[int]*Player)
	temp_outfits := make(map[string]int)

	//accounts
	rows, err := m.Connection.Query("SELECT id, name, guildRank, regTime FROM accounts")
	if err != nil {
		panic(err)
	}
	names := []string{}
	for rows.Next() {
		id := -1
		p := Player{Id: id}

		err := rows.Scan(&id, &p.Name, &p.GuildRank, &p.Created)
		if err != nil {
			panic(err)
		}

		temp[id] = &p
		names = append(names, p.Name)
	}

	//guilds
	rows, err = m.Connection.Query("SELECT name, members FROM guilds")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		members, name := "", ""

		err := rows.Scan(&name, &members)
		if err != nil {
			panic(err)
		}

		for _, member := range strings.Split(members, ",") {
			_member := strings.TrimSpace(member)
			if len(_member) > 0 {
				__member, err := strconv.Atoi(_member)

				if err != nil {
					panic(err)
				}

				if temp[__member] == nil {
					fmt.Printf("[DBCACHE] Cannot assign guild membership to account #%v (not in 'accounts')\n", __member)
				} else {
					temp[__member].Guild = name
				}
			}
		}
	}

	//class quests
	rows, err = m.Connection.Query("SELECT accId, objType, bestLv, bestFame FROM classstats")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		id, class, bestlevel, bestfame := -1, -1, 0, 0

		err := rows.Scan(&id, &class, &bestlevel, &bestfame)
		if err != nil {
			panic(err)
		}

		if temp[id] == nil {
			fmt.Printf("[DBCACHE] Cannot assign class stats to account #%v (not in 'accounts')\n", id)
		} else {
			if temp[id].ClassQuests == nil {
				temp[id].ClassQuests = make(map[int]ClassQuest)
			}

			c := ClassQuest{
				BestLevel: bestlevel,
				BestFame:  bestfame,
			}

			temp[id].ClassQuests[class] = c
		}
	}

	//account stats
	rows, err = m.Connection.Query("SELECT accId, fame FROM stats")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		id := -1
		accountfame := 0

		err := rows.Scan(&id, &accountfame)
		if err != nil {
			panic(err)
		}

		if temp[id] == nil {
			fmt.Printf("[DBCACHE] Cannot assign account stats to account #%v (not in 'accounts')\n", id)
		} else {
			temp[id].AccountFame = accountfame
		}
	}

	//characters
	rows, err = m.Connection.Query("SELECT accId, dead, lastSeen, charType, level, exp, fame, items, stats, tex1, tex2, petId, hasBackpack, skin FROM characters")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		accId, charType, level, exp, fame, tex1, tex2, petId, skin := -1, -1, -1, -1, -1, -1, -1, 0, -1
		items, stats, lastSeen := "", "", ""
		dead, hasBackpack := false, false

		err := rows.Scan(&accId, &dead, &lastSeen, &charType, &level, &exp, &fame, &items, &stats, &tex1, &tex2, &petId, &hasBackpack, &skin)
		if err != nil {
			panic(err)
		}

		if temp[accId] == nil {
			fmt.Printf("[DBCACHE] Cannot assign character stats to account #%v (not in 'accounts')\n", accId)
		} else {
			if dead {
				continue
			}

			_lastSeen := LastSeen{
				Time: lastSeen,
			}

			outfit := Outfit{
				Skin:      skin,
				Accessory: tex2,
				Clothing:  tex1,
			}

			_outfit := strings.Join(base.Aitoa([]int{outfit.Skin, outfit.Accessory, outfit.Clothing}), ",")

			temp_outfits[_outfit]++
			_stats := base.Aatoi(stats, ", ")
			_items := base.Aatoi(items, ", ")

			temp[accId].Characters = append(temp[accId].Characters, &Character{
				Class:    charType,
				Level:    level,
				Exp:      exp,
				Fame:     fame,
				Pet:      petId,
				Stats:    _stats,
				Items:    _items,
				Outfit:   outfit,
				Backpack: hasBackpack,
				LastSeen: _lastSeen,
			})
		}
	}

	//pets
	rows, err = m.Connection.Query("SELECT accId, petId, objType FROM pets")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		accId, petId, objType := -1, -1, -1

		err := rows.Scan(&accId, &petId, &objType)
		if err != nil {
			panic(err)
		}

		if temp[accId] == nil {
			fmt.Printf("[DBCACHE] Cannot assign pet stats to account #%v (not in 'accounts')\n", accId)
		} else {
			if temp[accId].Pets == nil {
				temp[accId].Pets = make(map[int]Pet)
			}
			temp[accId].Pets[petId] = Pet{
				Id:   petId,
				Type: objType,
			}
		}
	}

	//apply outfit count and transform map into array
	newcache := []*Player{}
	for _, p := range temp {
		sort.Sort(p)
		if len(p.Characters) > 0 {
			p.LastSeen = p.Characters[0].LastSeen
			p.LastSeen.Class = base.Capitalize(base.ClassString(p.Characters[0].Class))
		}
		for _, c := range p.Characters {
			c.OutfitCount = temp_outfits[strings.Join(base.Aitoa([]int{c.Outfit.Skin, c.Outfit.Accessory, c.Outfit.Clothing}), ",")]
			p.Fame += c.Fame
			p.Exp += c.Exp
		}
		for _, cq := range p.ClassQuests {
			p.Stars += base.FameGoals(cq.BestFame)
		}
		newcache = append(newcache, p)
	}

	//sort players and calculate ranks
	if len(newcache) > 1 {
		conv_sorted_accfame := []*Player{}
		conv_sorted_fame := []*Player{}
		conv_sorted_exp := []*Player{}

		exp_cmp := func(i interface{}, j interface{}) bool {
			return (i != nil && j != nil) && (i.(*Player).Exp >= j.(*Player).Exp)
		}
		fame_cmp := func(i interface{}, j interface{}) bool {
			return (i != nil && j != nil) && (i.(*Player).Fame >= j.(*Player).Fame)
		}
		accfame_cmp := func(i interface{}, j interface{}) bool {
			return (i != nil && j != nil) && (i.(*Player).AccountFame >= j.(*Player).AccountFame)
		}
		converted := make([]interface{}, len(newcache))
		for _, elem := range newcache {
			converted = append(converted, interface{}(elem))
		}

		sorted_accfame, sorted_fame, sorted_exp := base.MergeSort(converted, accfame_cmp), base.MergeSort(converted, fame_cmp), base.MergeSort(converted, exp_cmp)

		for i, p := range sorted_accfame {
			if p != nil {
				p.(*Player).AccountFameRank = i + 1
				conv_sorted_accfame = append(conv_sorted_accfame, p.(*Player))
			}
		}
		for i, p := range sorted_fame {
			if p != nil {
				p.(*Player).FameRank = i + 1
				conv_sorted_fame = append(conv_sorted_fame, p.(*Player))
			}
		}
		for i, p := range sorted_exp {
			if p != nil {
				p.(*Player).ExpRank = i + 1
				conv_sorted_exp = append(conv_sorted_exp, p.(*Player))
			}
		}

		m.Cache.PlayersByAccFame = conv_sorted_accfame
		m.Cache.PlayersByFame = conv_sorted_fame
		m.Cache.PlayersByExp = conv_sorted_exp
	}

	//swap arrays for no downtime
	m.Cache.Players = newcache
	m.Cache.PlayerNames = names
	m.Cache.Outfits = temp_outfits

	m.Cache.Initialized = true

	fmt.Printf("[DBCACHE] Cached %v players\n", len(newcache))
}

func (m *MySQL) RecentChanges() []RecentChange {
	return []RecentChange{
		RecentChange{
			Date: "21 July 2014",
			Changes: []template.HTML{
				template.HTML(`Now showing stats from <a href="http://c453.pw">c453.pw</a>.`),
				template.HTML(`Updated player profiles with fame/experience ranks.`),
			},
		},
	}
}

func (m *MySQL) FindPlayer(name string) (*Player, error) {
	for _, player := range m.Cache.Players {
		if strings.ToUpper(name) == strings.ToUpper(player.Name) {
			return player, nil
		}
	}

	p := Player{
		Name:    name,
		Similar: base.Similars(name, m.Cache.PlayerNames, 10, false),
	}
	return &p, errors.New("Player not found")
}

func (m *MySQL) SortPlayers(sorting string, offset int, num int) ([]*Player, error) {
	if offset > len(m.Cache.Players) {
		return []*Player{}, errors.New(fmt.Sprintf("invalid offset %d: only %d players", offset, num, len(m.Cache.Players)))
	}
	if num < 1 {
		return []*Player{}, nil
	}
	cached := []*Player{}
	switch sorting {
	case SortFame:
		cached = m.Cache.PlayersByFame
	case SortAccFame:
		cached = m.Cache.PlayersByAccFame
	case SortExp:
		cached = m.Cache.PlayersByExp
	default:
		return []*Player{}, errors.New("unknown sorting \"" + sorting + "\"")
	}
	return cached[int(math.Min(float64(len(cached)), float64(offset))):int(math.Min(float64(len(cached)), float64(offset+num)))], nil
}
