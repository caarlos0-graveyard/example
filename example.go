package example

import "time"

// Foo is my foo function
func Foo() error {
	time.Sleep(2 * time.Second)
	return nil
}
