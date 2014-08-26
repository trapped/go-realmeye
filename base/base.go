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

func Ordinal(num int) string {
    switch (num % 100) {
        case 11:
        case 12:
        case 13:
            return "th";
        default:
            break;
    }
    switch (num % 10) {
        case 1:
            return "st";
        case 2:
            return "nd";
        case 3:
            return "rd";
        default:
            return "th";
    }
}

func StarString(num int) string {
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
	case num >= 70:
		return "white"
	default:
		return ""
	}
}

func GuildRankString(num int) string {
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

func ClassString(num int) string {
	switch num {
	case 768:
		return "rogue"
	case 775:
		return "archer"
	case 782:
		return "wizard"
	case 784:
		return "priest"
	case 797:
		return "warrior"
	case 798:
		return "knight"
	case 799:
		return "paladin"
	case 800:
		return "assassin"
	case 801:
		return "necromancer"
	case 802:
		return "huntress"
	case 803:
		return "mystic"
	case 804:
		return "trickster"
	case 805:
		return "sorcerer"
	case 806:
		return "ninja"
	default:
		return ""
	}
}

func Capitalize(s string) (r string) {
	if len(s) > 0 {
		ss := strings.Split(strings.Replace(s, "-", " ", -1), " ")
		for i := range ss {
			if len(ss[i]) > 0 {
				ss[i] = string(unicode.ToUpper(rune(ss[i][0]))) + ss[i][1:]
			}
		}
		return strings.Join(ss, " ")
	}
	return ""
}

func Plural(s string) string {
	if strings.HasSuffix(s, "ess") {
		return s + "es"
	} else {
		return s + "s"
	}
}

func HumanTime(t string) string {
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

	datetime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)

	return humanize.Time(datetime)
}

func FameGoals(fame int) int {
	switch {
	case fame < 20:
		return 0
	case fame >= 20 && fame < 150:
		return 1
	case fame >= 150 && fame < 400:
		return 2
	case fame >= 400 && fame < 800:
		return 3
	case fame >= 800 && fame < 2000:
		return 4
	case fame >= 2000:
		return 5
	default:
		return 0
	}
}

func Last(a []string) string {
	return a[len(a)-1]
}

func HasExt(s string) bool {
	return len(strings.Split(s, ".")) > 1
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func div(a int, b int) int {
	return a / b
}

func mul(a int, b int) int {
	return a * b
}

func Aitoa(a []int) (result []string) {
	for i := 0; i < len(a); i++ {
		result = append(result, strconv.Itoa(a[i]))
	}
	return result
}

func Aatoi(s string, sep string) (result []int) {
	temp := strings.Split(s, sep)
	for i := 0; i < len(temp); i++ {
		n, err := strconv.Atoi(temp[i])
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}

func join(a []string, s string) (result string) {
	for i := 0; i < len(a); i++ {
		result += a[i]
		if i < len(a)-1 {
			result += s
		}
	}
	return result
}

func striplowercase(a string) (result string) {
	for _, r := range a {
		if unicode.IsUpper(r) {
			result += string(r)
		}
	}
	return result
}

//Calculates the *edit distance* between two strings:
//the required number of edits required to make them identical.
//Can be both case sensitive and insensitive.
func EDist(a string, b string, casesens bool, distance bool) float32 {
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
func Jaccard(a string, b string, casesens bool, distance bool) float32 {
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
			return 1 - Jaccard(a, b, casesens, !distance)
		} else {
			return (float32(len(union)) - float32(len(intersection))) / float32(len(union))
		}
	}

	return -1
}

//Which function the `Similars` should use.
var SIMILARITY_FUNC func(a string, b string, casesens bool, distance bool) float32 = func(a string, b string, casesens bool, distance bool) float32 {
	return float32((EDist(a, b, casesens, distance) / 2) + (Jaccard(a, b, casesens, distance))/2.5)
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

func Merge(l, r []interface{}, less func(interface{}, interface{}) bool) []interface{} {
	ret := make([]interface{}, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if less(l[0], r[0]) {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []interface{}, less func(interface{}, interface{}) bool) []interface{} {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n], less)
	r := MergeSort(s[n:], less)
	return Merge(l, r, less)
}

func (p *Page) Template(file string) *template.Template {
	p.Keywords = "free online mmo rpg game, realm of the mad god, rotmg, statistics, stats, rankings, ranks, " + p.Keywords
	cwd, _ := os.Getwd()
	tem := template.New(filepath.Base(file))
	tem = template.Must(tem.Funcs(template.FuncMap{
		"ordinal":         Ordinal,
		"starstring":      StarString,
		"guildrankstring": GuildRankString,
		"classstring":     ClassString,
		"humantime":       HumanTime,
		"split":           strings.Split,
		"last":            Last,
		"hasext":          HasExt,
		"add":             add,
		"sub":             sub,
		"div":             div,
		"mul":             mul,
		"plural":          Plural,
		"edist":           EDist,
		"famegoals":       FameGoals,
		"join":            join,
		"aitoa":           Aitoa,
		"capitalize":      Capitalize,
		"striplowercase":  striplowercase,
	}).ParseFiles(
		filepath.Join(cwd, "./base/index.gom"),
		filepath.Join(cwd, "./"+file),
	))
	return tem
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	b := Page{
		Title:    "Not found",
		Location: req.URL.String(),
	}

	tem := b.Template("base/notfound.gom")

	err := tem.Execute(w, b)
	if err != nil {
		panic(err)
	}
}
