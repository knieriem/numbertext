package numbers

import (
	"golang.org/x/text/language"
)

type table struct {
	zero                string
	oneOnly             string
	mult                []string
	tens                []string
	ones                []string
	teens               []string
	wordsSep            string
	tensSep             string
	onesBeforeTens      bool
	hundredsConjunction string
	tensConjunction     string
	teenConjunction     string
}

var wordsMap = map[language.Tag]*table{
	language.English: {
		wordsSep:            " ",
		tensSep:             "-",
		hundredsConjunction: " and ",
		zero:                "zero",
		oneOnly:             "one",
		onesBeforeTens:      false,
		mult: []string{
			"",
			"hundred",
			"thousand",
		},
		tens: []string{
			"",
			"ten",
			"twenty",
			"thirty",
			"fourty",
			"fifty",
			"sixty",
			"seventy",
			"eighty",
			"ninety",
		},
		teens: []string{
			0: "",
			1: "eleven",
			2: "twelve",
			3: "thirteen",
			4: "fourteen",
			5: "fifteen",
			6: "sixteen",
			7: "seventeen",
			8: "eighteen",
			9: "nineteen",
		},
		ones: []string{
			"",
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
		},
	},
	language.German: {
		zero:            "null",
		oneOnly:         "eins",
		tensConjunction: "und",
		onesBeforeTens:  true,
		mult: []string{
			"",
			"hundert",
			"tausend",
		},
		tens: []string{
			"",
			"zehn",
			"zwanzig",
			"dreißig",
			"vierzig",
			"fünfzig",
			"sechzig",
			"siebzig",
			"achzig",
			"neunzig",
		},
		teens: []string{
			0: "",
			1: "elf",
			2: "zwölf",
			6: "sechzehn",
			7: "siebzehn",
		},
		ones: []string{
			"",
			"ein",
			"zwei",
			"drei",
			"vier",
			"fünf",
			"sechs",
			"sieben",
			"acht",
			"neun",
		},
	},
}

func Words(n int, hasUnit bool, lang language.Tag) string {
	t, ok := wordsMap[lang]
	if !ok {
		t = wordsMap[language.English]
	}

	if n == 0 {
		return t.zero
	}
	digit := func() int {
		d := n % 10
		n -= d
		n /= 10
		return d
	}
	var s string
	ones := digit()
	tens := digit()
	hundreds := digit()
	thousands := digit()

	if tens == 0 {
		if ones == 1 && !hasUnit {
			s = t.oneOnly
		} else {
			s = t.ones[ones]
		}
	} else if tens == 1 {
		if ones < len(t.teens) {
			s = t.teens[ones]
			if s != "" {
				goto out
			}
		}
		s = t.tens[1]
		if t.onesBeforeTens {
			s = t.ones[ones] + s
		} else {
			s += t.teenConjunction + t.ones[ones]
		}
	} else {
		s = t.tens[tens]
		if t.onesBeforeTens {
			if ones != 0 {
				s = t.ones[ones] + t.tensConjunction + s
			}
		} else {
			s += t.tensConjunction + t.tensSep + t.ones[ones]
		}
	}
out:
	if hundreds != 0 {
		s = t.ones[hundreds] + t.wordsSep + t.mult[1] + t.hundredsConjunction + s
	}
	if thousands != 0 {
		s = t.ones[thousands] + t.wordsSep + t.mult[2] + t.wordsSep + s
	}
	return s
}
