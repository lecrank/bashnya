package unique

import (
	"testing"

	"github.com/lecrank/bashnya/parse"
	"github.com/stretchr/testify/assert"
)

// Define a test case struct
type testCase struct {
	name           string        // test case name
	input          []string      // function input
	flags          parse.Options // args
	expectedResult []string      // expected outcome
}

func TestFindUnique(t *testing.T) {
	// Define a slice of testCase as test table
	testTable := []testCase{
		{
			name:  "without flags",
			input: []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
			flags: parse.Options{},
			expectedResult: []string{"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks."},
		}, {
			name: "with -c -f 1 -i flags",
			input: []string{"I love music.",
				"I LoVe muSic.",
				"They love musIc.",
				"",
				"We love music of Kartik.",
				"I love music of Kartik.",
				"Thanks."},
			flags: parse.Options{C: true, F: 1, I: true},
			expectedResult: []string{"3 I love music.",
				"1 ",
				"2 We love music of Kartik.",
				"1 Thanks."},
		}, {
			name: "with -d -f 1 -s 1",
			input: []string{"Iq Pove music.",
				"Th Rove music.",
				"LK Gove music.",
				"",
				"Jq Dove music of Kartik.",
				"ql Zove music of Kartik.",
				"Thanks."},
			flags: parse.Options{D: true, F: 1, S: 1},
			expectedResult: []string{"Iq Pove music.",
				"Jq Dove music of Kartik."},
		}, {
			name: "with -u -f 5 -s 1",
			input: []string{"Iq Pove music.",
				"Th Rove music.",
				"LK Gove music.",
				"",
				"Jq Dove music of Kartik.",
				"ql Zove music of Kartik.",
				"Thanks."},
			flags:          parse.Options{D: true, F: 1, S: 1},
			expectedResult: []string{"", "Thanks."},
		},
	}
	// Begin test
	for _, test := range testTable {
		actual := FindUnique(test.input, test.flags)
		res := assert.Equal(t, test.expectedResult, actual, test.name)

		//assert.Nil and assert.NotNil to assert that the current test case does not or does expect an error respectively.

		assert.True(t, res)
	}
}
