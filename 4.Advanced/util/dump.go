package util

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/gtldhawalgandhi/go-training/4.Advanced/constants"
)

func Dump(v ...interface{}) {
	spew.Config.Indent = "\t"
	spew.Config.DisableMethods = true
	spew.Config.SortKeys = true
	spew.Config.DisableCapacities = true

	var printableValue interface{}

	if len(v) > 1 {
		typ := reflect.TypeOf(v[0]).Kind().String()

		switch typ {
		case "int":
			spew.Config.MaxDepth = v[0].(int)
			printableValue = v[1:]
		default:
			printableValue = v
		}
	}
	if GetEnv() != constants.EnvRelease {
		spew.Dump(printableValue)
	}
}
