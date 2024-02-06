package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc       string
		sl         []int
		want       int
		er         error
		searchItem int
	}{
		{
			desc:       "Good",
			sl:         []int{100, 101, 105, 200, 202, 107, 5, 40},
			want:       4,
			searchItem: 105,
			er:         nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := search(0, len(tC.sl), tC.searchItem, sortSlice(tC.sl))
			if err != nil {
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}
