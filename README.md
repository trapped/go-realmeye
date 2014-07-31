---

I take no responsibility for the content of this repository nor for the uses of the code it contains.
All serverside code is copyrighted to me; all clientside code is copyrighted to its authors (the RealmEye team, WildShadow/Kabam).

---

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
