package template_test

import (
	templateadapter "logur.dev/adapter/template"
)

func ExampleNew() {
	var l interface{}
	logger := templateadapter.New(l)

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := templateadapter.New(nil)

	// Output:
	_ = logger
}
