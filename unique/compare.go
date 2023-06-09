package unique

import (
	"strings"
	"unicode/utf8"

	"github.com/lecrank/bashnya/parse"
)

// returns the result of comparing to strings according to given arguments
func stringsAreEqual(str1, str2 string, opt parse.Options) bool {

	// -i
	if opt.I {
		str1 = strings.ToUpper(str1)
		str2 = strings.ToUpper(str2)
	}
	// -f
	if opt.F > 0 {
		cutFields(opt.F, &str1)
		cutFields(opt.F, &str2)
	}
	// -s
	if opt.S > 0 {
		cutSymbols(opt.S, &str1, &str2)
	}

	return str1 == str2
}

// cuts given strings by <amount> symbols
func cutSymbols(amount int, s1, s2 *string) {
	flag := false
	if amount >= utf8.RuneCountInString(*s1) {
		*s1 = ""
		flag = true
	}
	if amount >= utf8.RuneCountInString(*s2) {
		*s2 = ""
		flag = true
	}
	if !flag {
		rune_str1 := []rune(*s1)
		*s1 = string(rune_str1[amount:])

		rune_str2 := []rune(*s2)
		*s2 = string(rune_str2[amount:])
	}
}

// cuts given strings by <amount> fields and returns true if the both strings got empty (== strings are equal)
func cutFields(amount int, str *string) {

	if len(strings.Split(*str, " ")) > 1 {
		if amount < len(strings.Split(*str, " ")) {
			*str = strings.Join(strings.Split(*str, " ")[amount:], " ")
		} else {
			*str = ""
		}
	}
}
