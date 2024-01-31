package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordSplit(t *testing.T) {
	testCases := []struct {
		desc string
		row  string
		er   error
		want []string
	}{
		{
			desc: "Good test",
			row:  "Шла шла пирожок нашла",
			er:   nil,
			want: []string{"Шла", "шла", "пирожок", "нашла"},
		},
		{
			desc: "Error test",
			row:  "",
			er:   errors.New("передали пустую строку"),
			want: []string{""},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := WordSplit(tC.row)
			if err != nil {
				assert.Equal(t, tC.er, err)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestCountWorld(t *testing.T) {
	testCases := []struct {
		desc string
		row  []string
		want map[string]int
	}{
		{
			desc: "Good test",
			row:  []string{"шла", "шла", "пирожок", "нашла"},
			want: map[string]int{
				"нашла":   1,
				"пирожок": 1,
				"шла":     2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := countWorld(tC.row)
			assert.Equal(t, tC.want, got)
		})
	}
}
