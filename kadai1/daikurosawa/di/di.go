// Package di is dependency injection to image convert cil tool.
package di

import (
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/cli"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/convert"
	"github.com/gopherdojo/dojo3/kadai1/daikurosawa/option"
)

// InjectCli is dependency injection to Cli interface.
func InjectCli(convert convert.Convert, option *option.Option) cli.Cil {
	return cli.NewCli(convert, option)
}

// InjectConvert is dependency injection to Convert interface.
func InjectConvert(option *option.Option) convert.Convert {
	return convert.NewConvert(option)
}
