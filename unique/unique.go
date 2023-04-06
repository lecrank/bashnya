package unique

import (
	"github.com/lecrank/bashnya/parse"
)

func FindUnique(input []string, args parse.Options) []string {

	data := make(map[string]int)

	// get index map for the lines
	index_map := fillMap(data, input, args.F, args.S, args.I)

	// get strings in the right order according to the mode chosen (-c | -d | -u)
	output := pickLines(data, index_map, args.C, args.D, args.U)

	return output
}
