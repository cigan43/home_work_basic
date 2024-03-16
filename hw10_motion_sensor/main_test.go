package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenData(t *testing.T) {
	testCases := []struct {
		desc  string
		limit int64
		want  bool
	}{
		{
			desc:  "GOOD",
			limit: 20.0,
			want:  true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			testChan := make(chan int)
			go genData(testChan, tC.limit)
			_, ok := <-testChan
			if !ok {
				assert.Equal(t, tC.want, false)
			}
			assert.Equal(t, tC.want, ok)
		})
	}
}
