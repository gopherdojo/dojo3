package opt_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai3-1/nguyengiabk/opt"
)

func Example() {
	params, err := opt.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(params.Timeout)
}

var testParseFixtures = map[string]struct {
	args   []string
	err    error
	output *opt.Parameter
}{
	"Test invalid input type": {
		[]string{"-t", "0"},
		errors.New("Invalid timeout"),
		nil,
	},
	"Test default parameters": {
		[]string{},
		nil,
		&opt.Parameter{Timeout: 15},
	},
	"Test normal case": {
		[]string{"-t", "40"},
		nil,
		&opt.Parameter{Timeout: 40},
	},
}

func TestParse(t *testing.T) {
	for name, tc := range testParseFixtures {
		t.Run(name, func(t *testing.T) {
			result, err := opt.Parse(tc.args)
			if !reflect.DeepEqual(result, tc.output) {
				t.Errorf("Parse return wrong result, input = %v, actual = %v, expected = %v", tc.args, result, tc.output)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("Parse return wrong error, actual error = %v, expected error = %v", err, tc.err)
			}
		})
	}
}
