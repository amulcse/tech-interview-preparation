package main

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		data   []int
		value  int
		expect bool
	}{
		{[]int{1, 2, 3, 4, 5}, 6, false},
		{[]int{}, 1, false},
	}

	for _, tc := range testCases {
		t.Run("Testing values", func(t *testing.T) {
			result := binarySearch(tc.data, tc.value)
			if result != tc.expect {
				t.Errorf("For data %v and value %d, expected %v but got %v", tc.data, tc.value, tc.expect, result)
			}
		})
	}
}
