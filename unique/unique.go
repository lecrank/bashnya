package unique

import (
	"github.com/lecrank/bashnya/parse"
)

func FindUnique(input []string, args parse.Options) []string {

	data := make(map[string]int)

	// get index map for the lines
	index_map := fillMap(data, input, args)

	// get strings in the right order according to the mode chosen (-c | -d | -u)
	output := pickLines(data, index_map, args)

	return output
}
