package main

import (
	"fmt"
	"plugin"
	"reflect"
)

type Interpreter struct {
	functions map[string]reflect.Value
}

func NewInterpreter() *Interpreter {
	return &Interpreter{functions: make(map[string]reflect.Value)}
}

// LoadPlugin loads a Go plugin and registers its functions
func (interp *Interpreter) LoadPlugin(pluginPath string) error {
	// Load the plugin
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		return err
	}

	// Lookup the exported function
	printFunc, err := plug.Lookup("Print")
	if err != nil {
		return err
	}

	// Register the functions in the interpreter
	interp.functions["Print"] = reflect.ValueOf(printFunc)

	return nil
}

func (interp *Interpreter) CallFunction(name string, args ...any) {
	if fn, ok := interp.functions[name]; ok {
		// Convert arguments to reflect.Values
		in := make([]reflect.Value, len(args))
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}

		// Call the function
		fn.Call(in)
	} else {
		fmt.Printf("Error: function %s not found\n", name)
	}
}

func main() {
	// Initialize the interpreter
	interp := NewInterpreter()

	// Load the plugin
	err := interp.LoadPlugin("../dynlib/dynlib.so")
	if err != nil {
		fmt.Println("Error loading plugin:", err)
		return
	}

	// Call functions from the plugin
	interp.CallFunction("Add", 5, 3)       // Outputs: 8
	interp.CallFunction("Print", "Hello!") // Outputs: Hello!
}
