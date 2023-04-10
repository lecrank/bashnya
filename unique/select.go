package unique

import (
	"sort"
	"strconv"

	"github.com/lecrank/bashnya/parse"
)

// map(line => occurrence count)
func fillMap(lines map[string]int, text []string, args parse.Options) map[int]string {
	// map(index of line => line)
	index := make(map[int]string)
	counter := 0
	for _, elem := range text {
		flag := false
		for line := range lines {
			if stringsAreEqual(elem, line, args.F, args.S, args.I) {
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
func pickLines(data map[string]int, indexMap map[int]string, args parse.Options) []string {
	indexValues := make([]int, 0, len(data))

	// sort keys of the index map (index of line => line) to access data map (line => count) in the rigth order
	for key := range indexMap {
		indexValues = append(indexValues, key)
	}
	sort.Ints(indexValues)

	result := []string{"error"}
	// if count mode
	if args.C {
		result = countLines(indexValues, data, indexMap)
	} else if args.D { // if non-unique mode
		result = duplicateOnly(indexValues, data, indexMap)
	} else if args.U { // if unique mode
		result = uniqueOnly(indexValues, data, indexMap)
	} else if !args.C && !args.D && !args.U { // if regular mode
		result = regularMode(indexValues, data, indexMap)
	}
	return result
}
