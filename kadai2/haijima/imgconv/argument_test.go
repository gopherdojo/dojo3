package imgconv

import (
	"bytes"
	"reflect"
	"testing"
	"fmt"
	"os"
)

func TestArgument_Validate(t *testing.T) {
	type fields struct {
		Option Option
		Dir    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "#1 normal case", fields: fields{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "png"}, Dir: "."}, wantErr: false},
		{name: "#2 quiet dry run", fields: fields{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "png", DryRun: true, Quiet: true}, Dir: "."}, wantErr: true},
		{name: "#3 input extension and output extension are same", fields: fields{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "jpg", DryRun: true, Quiet: true}, Dir: "."}, wantErr: true},
		{name: "#4 invalid input extension", fields: fields{Option: Option{Input: []string{"jpg", "svg"}, Output: "png", DryRun: true, Quiet: true}, Dir: "."}, wantErr: true},
		{name: "#5 invalid output extension", fields: fields{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "jpeg", DryRun: true, Quiet: true}, Dir: "."}, wantErr: true},
		{name: "#6 validate input extension always case insensitively", fields: fields{Option: Option{Input: []string{"JPG", "JpEg"}, Output: "png", CaseSensitive: true}, Dir: "."}, wantErr: false},
		{name: "#7 validate output extension always case sensitively", fields: fields{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "`PNG", CaseSensitive: true}, Dir: "."}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Argument{
				Option: tt.fields.Option,
				Dir:    tt.fields.Dir,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Argument.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

const helpMsg = "Usage of conv:\nconv [-i srcExts] [-o destExt] [-w] [--dry-run|-q] [-s] [directory]\n  -dry-run\n    \tDry run mode\n  -i string\n    \tInput extension. (default \"jpg|jpeg\")\n  -o string\n    \tOutput extension. (default \"png\")\n  -q\tQuiet mode. Suppress print\n  -s\tMatches file extension case-sensitively\n  -w\tIf converted file has already existed, Overwrite old files.\n"

func TestCreateArgument(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    Argument
		wantOut string
		wantErr bool
	}{
		{name: "#1 normal case", args: args{args: []string{"conv"}}, want: Argument{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "png"}, Dir: "."}, wantOut: "", wantErr: false},
		{name: "#2 specify all options", args: args{args: []string{"conv", "-i", "png", "-o", "jpg", "-w", "-q", "--dry-run", "someDir"}}, want: Argument{Option: Option{Input: []string{"png"}, Output: "jpg", Overwrite: true, DryRun: true, Quiet: true}, Dir: "someDir"}, wantOut: "", wantErr: false},
		{name: "#3 too many Directories", args: args{args: []string{"conv", "someDir", "otherDirectory"}}, want: Argument{}, wantOut: "", wantErr: true},
		{name: "#4 invalid option", args: args{args: []string{"conv", "-hoge"}}, want: Argument{}, wantOut: "flag provided but not defined: -hoge\n" + helpMsg, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			got, err := CreateArgument(tt.args.args, out)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateArgument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateArgument() = %v, want %v", got, tt.want)
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("CreateArgument() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func ExampleCreateArgument() {
	argument, _ := CreateArgument([]string{""}, os.Stdout)
	fmt.Println(&argument)
	// Output:
	// Argument{ Option: Option{ Input: [jpg jpeg], Output: png, Overwrite: false, DryRun: false, Quiet: false, CaseSensitive: false }, Dir: . }
}

func Test_createDefaultArg(t *testing.T) {
	tests := []struct {
		name string
		want Argument
	}{
		{name: "#1 normal case", want: Argument{Dir: ".", Option: Option{Input: []string{"jpg", "jpeg"}, Output: "png", Overwrite: false, DryRun: false}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createDefaultArg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleArgument_String() {
	fmt.Println(&Argument{Option: Option{Input: []string{"jpg", "jpeg"}, Output: "png"}, Dir: "."})
	// Output:
	// Argument{ Option: Option{ Input: [jpg jpeg], Output: png, Overwrite: false, DryRun: false, Quiet: false, CaseSensitive: false }, Dir: . }
}

func ExampleOption_String() {
	fmt.Println(&Option{Input: []string{"jpg", "jpeg"}, Output: "png"})
	// Output:
	// Option{ Input: [jpg jpeg], Output: png, Overwrite: false, DryRun: false, Quiet: false, CaseSensitive: false }
}
