package chessboard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc, want             string
		countColumn, firstStep int
	}{
		{
			desc:        "OneBlack",
			countColumn: 8,
			firstStep:   0,
			want:        "# # # # ",
		},
		{
			desc:        "OneWhite",
			countColumn: 9,
			firstStep:   1,
			want:        " # # # # ",
		},
		{
			desc:        "ErrorOneWhite",
			countColumn: 5,
			firstStep:   1,
			want:        "# # #",
		},
		{
			desc:        "ErrorString",
			countColumn: 6,
			firstStep:   2,
			want:        "# # #",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := CreateString(tC.countColumn, tC.firstStep)
			assert.Equal(t, got, tC.want)
		})
	}
}
