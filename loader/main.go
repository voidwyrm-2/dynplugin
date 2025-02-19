package main

import (
	"fmt"
	"plugin"
	"reflect"
)

func CallFunction(fnSyb plugin.Symbol, args ...any) {
	fn := reflect.ValueOf(fnSyb)

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	fn.Call(in)
}

const pluginPath = "../dynlib/dynlib.so"

func main() {
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	printFunc, err := plug.Lookup("Print")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	CallFunction(printFunc, "Hello!")
}
