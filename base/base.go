package base

import (
	humanize "github.com/dustin/go-humanize"
	"html/template"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type Page struct {
	Title       string
	Location    string
	Keywords    string
	Description string
	Canonical   string
	Visit       string
	Specific    interface{}
}

func ordinal(num int) string {
	s := strconv.Itoa(num)
	d := s[len(s)-1]
	switch d {
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

func starstring(num int) string {
	switch {
	case num >= 0 && num <= 13:
		return "light-blue"
	case num >= 14 && num <= 27:
		return "blue"
	case num >= 28 && num <= 41:
		return "red"
	case num >= 42 && num <= 55:
		return "orange"
	case num >= 56 && num <= 69:
		return "yellow"
	case num == 70:
		return "white"
	default:
		return ""
	}
}

func guildrankstring(num int) string {
	switch num {
	case 0:
		return "Initiate"
	case 10:
		return "Member"
	case 20:
		return "Officer"
	case 30:
		return "Leader"
	case 40:
		return "Founder"
	default:
		return ""
	}
}

func classstring(num int) string {
	switch num {
	case 768:
		return "Rogue"
	case 775:
		return "Archer"
	case 782:
		return "Wizard"
	case 784:
		return "Priest"
	case 797:
		return "Warrior"
	case 798:
		return "Knight"
	case 800:
		return "Assassin"
	case 801:
		return "Necromancer"
	case 802:
		return "Huntress"
	case 803:
		return "Mystic"
	case 804:
		return "Trickster"
	case 805:
		return "Sorcerer"
	case 806:
		return "Ninja"
	default:
		return ""
	}
}

func plural(s string) string {
	if strings.HasSuffix(s, "ess") {
		return s + "es"
	} else {
		return s + "s"
	}
}

func humantime(t string) string {
	if len(t) == 0 {
		return "never"
	}
	parts := strings.Split(t, " ")
	date := strings.Split(parts[0], "-")
	year, err := strconv.Atoi(date[0])
	if err != nil {
		panic(err)
	}
	month, err := strconv.Atoi(date[1])
	if err != nil {
		panic(err)
	}
	day, err := strconv.Atoi(date[2])
	if err != nil {
		panic(err)
	}
	timep := strings.Split(parts[1], ":")
	hour, err := strconv.Atoi(timep[0])
	if err != nil {
		panic(err)
	}
	minute, err := strconv.Atoi(timep[1])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(timep[2])
	if err != nil {
		panic(err)
	}

	datetime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)

	return humanize.Time(datetime)
}

func split(t string, s string) []string {
	return strings.Split(t, s)
}

func last(a []string) string {
	return a[len(a)-1]
}

func hasext(s string) bool {
	return len(strings.Split(s, ".")) > 1
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

//Calculates the *edit distance* between two strings:
//the required number of edits required to make them identical.
//Can be both case sensitive and insensitive.
func edist(a string, b string, casesens bool, distance bool) float32 {
	if a == b {
		return 0
	}
	//difference in length
	diff := int(math.Abs(float64(len(a) - len(b))))
	shortest := a
	longest := b
	if len(b) < len(a) {
		shortest = b
		longest = a
	}
	for i := 0; i < len(shortest); i++ {
		if shortest[i] != longest[i] {
			if !casesens {
				if unicode.ToUpper(rune(shortest[i])) != unicode.ToUpper(rune(longest[i])) {
					diff++
				}
			} else {
				diff++
			}
		}
	}
	return (float32(diff) / 3) / float32(len(longest))
}

//Whether the `jaccard` function should recurse when calculating the Jaccard Distance.
var JACCARD_RECURSIVE bool = false

//Calculates the Jaccard Index (Similarity coefficient) between two sets;
//formula: `J(A,B) = |A n B| / |A u B|`
//given that `J(A,B) = 1` if `A` and `B` are both empty (100% similarity).
//Respectively, the Jaccard Distance (Dissimilarity coefficient) between two
//sets has the following formula: `Jd(A,B) = 1 - J(A,B) = (|A u B| - |A n B|) / |A u B|`
//given that `J(A,B) = 0` if `A` and `B` are both empty (0% dissimilarity).
func jaccard(a string, b string, casesens bool, distance bool) float32 {
	intersection := ""
	union := ""

	//calculate intersection size
	for _, c := range a {
		if casesens {
			c = rune(strings.ToUpper(string(c))[0])
		}
		if strings.ContainsRune(b, c) && !strings.ContainsRune(intersection, c) {
			intersection += string(c)
		}
	}

	//calculate union size
	for _, c := range a + b {
		if casesens {
			c = rune(strings.ToUpper(string(c))[0])
		}
		if !strings.ContainsRune(union, c) {
			union += string(c)
		}
	}

	if !distance {
		return float32(len(intersection)) / float32(len(union))
	} else {
		if JACCARD_RECURSIVE {
			return 1 - jaccard(a, b, casesens, !distance)
		} else {
			return (float32(len(union)) - float32(len(intersection))) / float32(len(union))
		}
	}

	return -1
}

//Which function the `Similars` should use.
var SIMILARITY_FUNC func(a string, b string, casesens bool, distance bool) float32 = func(a string, b string, casesens bool, distance bool) float32 {
	return float32((edist(a, b, casesens, distance) / 2) + (jaccard(a, b, casesens, distance))/2.5)
}

func Similars(needle string, haystack []string, max int, casesens bool) (r []string) {
	tolerance := float32(1) / (float32(len(needle)) / 2) //float32(len(needle)) / 100 * 30
	for i := 0; i < len(haystack) && len(r) <= max; i++ {
		dist := SIMILARITY_FUNC(needle, haystack[i], casesens, true)
		if dist <= tolerance {
			r = append(r, haystack[i])
		}
	}
	return r
}

func (p *Page) Template(file string) *template.Template {
	cwd, _ := os.Getwd()
	tem := template.New(filepath.Base(file))
	tem = template.Must(tem.Funcs(template.FuncMap{
		"ordinal":         ordinal,
		"starstring":      starstring,
		"guildrankstring": guildrankstring,
		"classstring":     classstring,
		"humantime":       humantime,
		"split":           split,
		"last":            last,
		"hasext":          hasext,
		"add":             add,
		"sub":             sub,
		"plural":          plural,
		"edist":           edist,
	}).ParseFiles(
		filepath.Join(cwd, "./base/index.gom"),
		filepath.Join(cwd, "./"+file),
	))
	return tem
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	b := Page{
		Title:    "Not found | RealmEye",
		Location: req.URL.String(),
	}

	tem := b.Template("base/notfound.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
