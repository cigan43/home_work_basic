package comparator

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		book1   Book
		book2   Book
		compare int
		want    bool
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
