package unique

import (
	"sort"
	"strconv"
)

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

// returns the slice which strings contain occurrence number
func countLines(idValues []int, data map[string]int, idMap map[int]string) []string {

	result := []string{}

	for index := range idValues {
		line := idMap[index]
		count := data[line]
		result = append(result, strconv.Itoa(count)+" "+line)
	}

	return result
}

// returns only the lines that were duplicated
func duplicateOnly(idValues []int, data map[string]int, idMap map[int]string) []string {

	result := []string{}

	for index := range idValues {
		line := idMap[index]
		if data[line] > 1 {
			result = append(result, line)
		}
	}
	return result
}

// returns only unique lines
func uniqueOnly(idValues []int, data map[string]int, idMap map[int]string) []string {

	result := []string{}

	for index := range idValues {
		line := idMap[index]
		if data[line] == 1 {
			result = append(result, line)
		}
	}
	return result
}

// returns lines without duplications
func regularMode(idValues []int, data map[string]int, idMap map[int]string) []string {

	result := []string{}

	for index := range idValues {
		line := idMap[index]
		result = append(result, line)
	}
	return result
}

// returns the appropriate lines
func pickLines(data map[string]int, indexMap map[int]string, c, d, u bool) []string {
	indexValues := make([]int, 0, len(data))

	// sort keys of the index map (index of line => line) to access data map (line => count) in the rigth order
	for key := range indexMap {
		indexValues = append(indexValues, key)
	}
	sort.Ints(indexValues)

	result := []string{"error"}
	// if count mode
	if c {
		result = countLines(indexValues, data, indexMap)
	}
	// if non-unique mode
	if d {
		result = countLines(indexValues, data, indexMap)
	}
	// if unique mode
	if u {
		result = uniqueOnly(indexValues, data, indexMap)
	}
	// if regular mode
	if !c && !d && !u {
		result = regularMode(indexValues, data, indexMap)
	}
	return result
}
