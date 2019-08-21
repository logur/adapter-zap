package zap_test

import (
	zapadapter "logur.dev/adapter/zap"
)

func ExampleNew() {
	var l interface{}
	logger := zapadapter.New(l)

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := zapadapter.New(nil)

	// Output:
	_ = logger
}
