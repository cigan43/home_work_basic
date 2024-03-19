package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	testCases := []struct {
		desc string
		cnt  *Counter
		// finish chan struct{}
		want int
	}{
		{
			desc: "Good",
			cnt:  &Counter{num: 0},
			// finish: make(chan struct{}),
			want: 20,
		},
	}
	finish := make(chan struct{})
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			go Do(tC.cnt, finish)
			<-finish
			assert.Equal(t, tC.want, tC.cnt.Value())
		})
	}
}
