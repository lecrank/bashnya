package unique

import (
	"sort"
	"strconv"
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
		var zero_fields1, zero_fields2 bool

		if len(strings.Split(str1, " ")) < 2 {
			zero_fields1 = true
		}
		if len(strings.Split(str2, " ")) < 2 {
			zero_fields2 = true
		}

		if f >= len(strings.Split(str1, " ")) && f >= len(strings.Split(str2, " ")) && !zero_fields1 && !zero_fields2 {
			return true
		}
		if !zero_fields1 {
			str1 = strings.Join(strings.Split(str1, " ")[f:], " ")
		}
		if !zero_fields2 {
			str2 = strings.Join(strings.Split(str2, " ")[f:], " ")
		}
	}
	// -s
	if s > 0 {
		flag := false
		if s >= utf8.RuneCountInString(str1) {
			str1 = ""
			flag = true
		}
		if s >= utf8.RuneCountInString(str2) {
			str2 = ""
			flag = true
		}
		if !flag {
			rune_str1 := []rune(str1)
			str1 = string(rune_str1[s:])

			rune_str2 := []rune(str2)
			str2 = string(rune_str2[s:])
		}
	}
	// comparison after changes
	if str1 == str2 {
		return true
	} else {
		return false
	}
}

// map(line => occurrence count)
func fillMap(lines map[string]int, text []string, f, s int, i bool) map[int]string {
	// map(index of line => line)
	index := make(map[int]string)
	counter := 0
	for _, elem := range text {
		flag := false
		for line := range lines {
			if stringsAreEqual(elem, line, f, s, i) {
				flag = true
				lines[line]++
			}
		}
		// add to the both maps if not in lines map
		if !flag {
			lines[elem] = 1
			index[counter] = elem
			counter++
		}
	}
	return index
}

func pickLines(data map[string]int, index_map map[int]string, c, d, u bool) []string {
	index_values := make([]int, 0, len(data))

	// sort keys of the index map (index of line => line) to access data map (line => count) in the rigth order
	for key := range index_map {
		index_values = append(index_values, key)
	}
	sort.Ints(index_values)

	result := []string{}
	// if count mode
	if c {
		for index := range index_values {
			line := index_map[index]
			count := data[line]
			result = append(result, strconv.Itoa(count)+" "+line)
		}
		return result
	}
	// if non-unique mode
	if d {
		for index := range index_values {
			line := index_map[index]
			if data[line] > 1 {
				result = append(result, line)
			}
		}
		return result
	}
	// if unique mode
	if u {
		for index := range index_values {
			line := index_map[index]
			if data[line] == 1 {
				result = append(result, line)
			}
		}
	}
	// if regular mode
	if !c && !d && !u {
		for index := range index_values {
			line := index_map[index]
			result = append(result, line)
		}
		return result
	}
	return result
}

func GetOutput(input []string, c, d, u bool, f, s int, i bool) []string {

	data := make(map[string]int)
	// get index map for the lines
	index_map := fillMap(data, input, f, s, i)
	// get strings in the right order according to the mode chosen (-c | -d | -u)
	output := pickLines(data, index_map, c, d, u)
	return output
}
