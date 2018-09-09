package opt_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo3/kadai3-2/nguyengiabk/opt"
)

func Example() {
	params, err := opt.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(params.ProcNum)
}

var testParseFixtures = map[string]struct {
	args   []string
	err    error
	output *opt.Parameter
}{
	"Test invalid input type": {
		[]string{"-p", "0", "https://file.to.download"},
		errors.New("Invalid number of parallel processes"),
		nil,
	},
	"Test no input specified": {
		[]string{},
		errors.New("URL was not specified"),
		nil,
	},
	"Test default parameters": {
		[]string{"https://file.to.download"},
		nil,
		&opt.Parameter{ProcNum: 4, URL: "https://file.to.download"},
	},
	"Test normal case": {
		[]string{"-p", "6", "https://file.to.download"},
		nil,
		&opt.Parameter{ProcNum: 6, URL: "https://file.to.download"},
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
