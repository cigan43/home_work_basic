package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		Address string
		URL     string
		Data    []byte
		want    string
	}{
		{
			desc:    "Post",
			Address: "127.0.0.1:8000",
			URL:     "post",
			Data:    []byte(`{"Name":"John", "LastName": "Pup", "Age":30}`),
			want:    "200 ok",
		},
		{
			desc:    "get",
			Address: "127.0.0.1:8000",
			URL:     "get",
			want:    "200 ok",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.URL == "post" {
				status, _ := clientPost(tC.Address, tC.URL, tC.Data)
				assert.Equal(t, tC.want, status)
			} else {
				status, _ := clientGet(tC.Address, tC.URL)
				assert.Equal(t, tC.want, status)
			}
		})
	}
}
