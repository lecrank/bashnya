package parse

import (
	"errors"
	"flag"
)

type Options struct {
	C bool // count mode
	D bool // duplicate mode
	U bool // unique mode
	I bool // ignore register flag
	F int  // ignored fields count
	S int  // ignored symbols count
}

func GetFlags(flags Options) (Options, error) {
	var c, d, u, i bool
	var f, s int
	//var input_file string //, output_file string

	flag.BoolVar(&c, "c", false, "occurrence count")
	flag.BoolVar(&d, "d", false, "non-unique only")
	flag.BoolVar(&u, "u", false, "unique only")
	flag.IntVar(&f, "f", 0, "amount of ignored fields")
	flag.IntVar(&s, "s", 0, "amount of ignored chars")
	flag.BoolVar(&i, "i", false, "ignore register")
	flag.Parse()

	err := flags.parseOptions(c, d, u, f, s, i)
	if err != nil {
		return flags, err
	}
	return flags, nil
}

func (opt *Options) parseOptions(c, d, u bool, f, s int, i bool) error {

	// check if -c -d -u used simultaneously
	count := 0
	for _, flag := range []bool{c, d, u} {
		if flag {
			count++
		}
	}

	if count > 1 {
		return errors.New("Wrong usage!\n\nCorrect: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
	}

	opt.C = c
	opt.D = d
	opt.U = u
	opt.I = i
	opt.F = f
	opt.S = s

	return nil
}

func OptionsGiven() (Options, error) {
	options := Options{}
	flags, err := GetFlags(options)
	if err != nil {
		return flags, err
	}
	return flags, nil
}
