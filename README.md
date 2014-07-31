---

I take no responsibility for the content of this repository nor for the uses of the code it contains.
All serverside code is copyrighted to me; all clientside code is copyrighted to its authors (the RealmEye team, WildShadow/Kabam).

---

####Configuration

Currently this only supports two types of database: a bogus/example one (for easy testing/debugging purposes) and a MySQL one which currently runs in a "private server DB"-compatibility mode with complete data caching (refreshing every minute; I've only implemented this at the moment as I don't want to run my own DB and plugging someone else's was the laziest choice).

You can easily configure the DB settings in the config.json file:

```json
{
	"Type": "mysql",
	"Host": "localhost:3306",
	"User": "realmeye",
	"Password": "",
	"Schema": "rotmg",
	"Cached": true
}
```

####Status/todo

- [x] Recent changes
- [x] Home page
- [x] Player account stats
- [x] Player characters
- [ ] Characters' last seen server
- [ ] Player pets
- [ ] Player graveyard
- [ ] Top players (by fame/etc.)
- [ ] Guilds
- [ ] Top guilds (by fame/etc.)
- [ ] Wiki
- [ ] MrEyeBall
- [ ] ...

####String similarity equations/algorithms used for the "player not found" page

I've implemented both the **Edit Distance** (number of edits needed to transform a string into another) and the **Jaccard Index** (also called **Similarity Coefficient**) algorithms/equations - though the latter includes both the *similarity* and the *dissimilarity* parts: [Gist](https://gist.github.com/trapped/d1e62dd3b05e00bfd904)
