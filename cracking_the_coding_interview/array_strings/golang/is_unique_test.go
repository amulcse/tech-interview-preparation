package main

import "testing"

func TestIsUnique(t *testing.T) {

	testCases := []struct {
		text   string
		result bool
	}{
		{"amul", true},
		{"aa", false},
		{"aA", true},
	}

	for _, c := range testCases {
		t.Run("unique string test", func(t *testing.T) {
			r := isUnique(c.text)
			if r != c.result {
				t.Errorf("%s result should be %t but got %t", c.text, c.result, r)
			}
		})
	}
}
