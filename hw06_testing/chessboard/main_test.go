package chessboard

import (
	"testing"
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
			desc:        "OneBlack1",
			countColumn: 5,
			firstStep:   0,
			want:        "# # #",
		},
		{
			desc:        "OneWhite2",
			countColumn: 6,
			firstStep:   3,
			want:        " # # #",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := CreateString(tC.countColumn, tC.firstStep)
			// assert.Equal(t, tC.want, got)
			if got != tC.want {
				t.Errorf("got %s, want %s", got, tC.want)
			}
		})
	}
}
