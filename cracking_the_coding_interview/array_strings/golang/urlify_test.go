package main

import "testing"

func testMain(t *testing.T) {
	t.Run("test encoding", func(t *testing.T) {
		r := encodeURL("Hey there!")
		if r != "Hey%20there" {
			t.Errorf("%s is not encoded url", r)
		}
	})
}
