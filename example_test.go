package zap_test

import (
	"go.uber.org/zap"

	zapadapter "logur.dev/adapter/zap"
)

func ExampleNew() {
	logger := zapadapter.New(zap.NewNop())

	// Output:
	_ = logger
}

// If logger is nil, a default instance is created.
func ExampleNew_default() {
	logger := zapadapter.New(nil)

	// Output:
	_ = logger
}
