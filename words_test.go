package numbers

import (
	"strings"
	"testing"
	"strconv"

	"golang.org/x/text/language"
)

var data = `
1
	one
	eins
2
	two
	zwei
23
	twenty-three
	dreiundzwanzig
101€
	one hundred and one
	einhundertein
101
	one hundred and one
	einhunderteins
211
	two hundred and eleven
	zweihundertelf
999
	nine hundred and ninety-nine
	neunhundertneunundneunzig
1013
	one thousand and thirteen
	eintausenddreizehn
`

type testItem struct {
	number int
	hasUnit bool
	text []string
}

func TestWords(t *testing.T) {
	items := parseTestData()
	for _, it := range items {
		testWordsLang(t, it, it.text[0], language.English)
		testWordsLang(t, it, it.text[1], language.German)
	}
}

func testWordsLang(t *testing.T, it *testItem, want string, lang language.Tag) {
	text := Words(it.number, it.hasUnit, lang)
	if text != want {
		t.Errorf("%d: %v %q != %q", it.number, lang, text, want)
	}
}

func parseTestData() []*testItem {
	var items []*testItem

	var it *testItem
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "\t") {
			if it != nil {
				items = append(items, it)
			}
			it = new(testItem)
			if strings.HasSuffix(line, "€") {
				it.hasUnit = true
				line = line[:len(line) - len("€")]
			}
			v, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			it.number = v
			continue
		}
		it.text = append(it.text, line[1:])
	}
	items = append(items, it)
	return items
}
