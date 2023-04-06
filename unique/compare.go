package unique

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// returns the result of comparing to strings according to given arguments
func stringsAreEqual(str1, str2 string, f, s int, ignore_reg bool) bool {

	// -i
	if ignore_reg {
		str1 = strings.ToUpper(str1)
		str2 = strings.ToUpper(str2)
	}
	// -f
	if f > 0 {
		if cutFields(f, &str1, &str2) {
			return true
		}
	}
	// -s
	if s > 0 {
		cutSymbols(s, &str1, &str2)
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
func cutFields(amount int, s1, s2 *string) bool {

	var noFields1, noFields2 bool
	var preResult bool

	if len(strings.Split(*s1, " ")) < 2 {
		noFields1 = true
	}
	if len(strings.Split(*s2, " ")) < 2 {
		noFields2 = true
	}
	fmt.Println("here")
	if amount >= len(strings.Split(*s1, " ")) && amount >= len(strings.Split(*s2, " ")) && !noFields1 && !noFields2 {
		return true
	}

	if !noFields1 {
		*s1 = strings.Join(strings.Split(*s1, " ")[amount:], " ")
	}
	fmt.Println("again")
	if !noFields2 {
		*s2 = strings.Join(strings.Split(*s2, " ")[amount:], " ")
	}
	fmt.Println("after s2")
	return preResult
}
