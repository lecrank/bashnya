package unique

/////////////////////////////////////////// DOESN'T WORK :(

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Define a test case struct
type testCase struct {
	name           string  // test case name
	input          string  // function input
	flags          Options // args
	expectedResult bool    // expected outcome
	hasError       bool    // if error is expected
}

func TestFindUnique(t *testing.T) {
	// Define a slice of testCase as test table
	testTable := []testCase{
		{
			name: "without flags",
			input: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Kartik.",
				"I love music of Kartik.",
				"Thanks."},
			flags: Options{},
			expectedResult: []string{"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks."},
			hasError: false,
		}, {
			name: "with -c -f 1 -i flags",
			input: []string{"I love music.",
				"I LoVe muSic.",
				"They love musIc.",
				"",
				"We love music of Kartik.",
				"I love music of Kartik.",
				"Thanks."},
			flags: Options{c: true, f: 1, i: true},
			expectedResult: []string{"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks."},
			hasError: false,
		}, {
			name: "",
			input: []string{"Iq Pove music.",
				"Th Rove music.",
				"LK Gove music.",
				"",
				"Jq Dove music of Kartik.",
				"ql Zove music of Kartik.",
				"Thanks."},
			flags: Options{d: true, f: 1, s: 1},
			expectedResult: []string{"Iq Pove music.",
				"Jq Dove music of Kartik."},
			hasError: false,
		},
	}
	// Begin test
	for _, test := range testTable {
		actual, err := FindUnique(test.input, test.flags)
		assert.Equal(t, test.expectedResult, actual, test.name)

		// assert.Nil and assert.NotNil to assert that the current test case does not or does expect an error respectively.
		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}
