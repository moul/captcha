package main

import (
	"bytes"
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func Example() {
	var stdin bytes.Buffer
	stdin.Write([]byte("7\n"))
	err := run([]string{"captcha", "-seed=42"}, &stdin)
	if err != nil {
		panic(err)
	}

	// Output:
	// 3 + 4
	// ->
}
