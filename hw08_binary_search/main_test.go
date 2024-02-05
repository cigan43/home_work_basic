package main

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc string
		sl   []int
		wont int
		er   error
	}{
		{
			desc: "Good",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
