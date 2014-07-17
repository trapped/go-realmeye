package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/trapped/realmeye/base"
	"strconv"
	"strings"
	"time"
)

//updates done through the web interface must both update the data in the cache AND queue
//a db update
type Cache struct {
	Initialized bool
	Players     []*Player
	PlayerNames []string
}

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
	temp := make(map[int]*Player)
	//accounts
	rows, err := m.Connection.Query("SELECT id, name, guildRank, regTime FROM accounts")
	if err != nil {
		panic(err)
	}
	names := []string{}
	for rows.Next() {
		id := -1
		p := Player{}
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
		members := ""
		name := ""
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
					fmt.Printf("[DBCACHE] Cannot assign guild membership to account #%v (not in the 'accounts' table)\n", __member)
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
		id := -1
		class := -1
		bestlevel := 0
		bestfame := 0
		err := rows.Scan(&id, &class, &bestlevel, &bestfame)
		if err != nil {
			panic(err)
		}
		if temp[id] == nil {
			fmt.Printf("[DBCACHE] Cannot assign class stats to account #%v (not in the 'accounts' table)\n", id)
		} else {
			temp[id].ClassQuests = make(map[int]ClassQuest)
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
			fmt.Printf("[DBCACHE] Cannot assign account stats to account #%v (not in the 'accounts' table)\n", id)
		} else {
			temp[id].AccountFame = accountfame
		}
	}
	//swap arrays for no downtime
	newcache := []*Player{}
	for _, p := range temp {
		newcache = append(newcache, p)
	}
	m.Cache.Players = newcache
	m.Cache.PlayerNames = names
	fmt.Printf("[DBCACHE] Cached %v players\n", len(newcache))
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
