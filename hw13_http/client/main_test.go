package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientGet(t *testing.T) {
	testCases := []struct {
		desc    string
		Address string
		URL     string
		Data    []byte
		want    string
	}{
		// {
		// 	desc:    "Post",
		// 	Address: "127.0.0.1:8000",
		// 	URL:     "post",
		// 	Data:    []byte(`{"Name":"John", "LastName": "Pup", "Age":30}`),
		// 	want:    "200 OK",
		// },
		{
			desc:    "get",
			Address: "127.0.0.1:8000",
			URL:     "get",
			want:    "200 OK",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// if tC.URL == "post" {
			// 	status, _ := clientPost(tC.Address, tC.URL, tC.Data)
			// 	assert.Equal(t, tC.want, status)
			// } else {

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				// Test request parameters
				assert.Equal(t, req.URL.String(), "/get")
				// Send response to be tested
				rw.WriteHeader(http.StatusOK)
				rw.Write([]byte(`200 OK`))

			}))
			// Close the server when test finishes
			defer server.Close()
			status, _ := clientGet(server.URL, tC.URL)
			fmt.Println(status, "UUUUUU")
			// assert.Equal(t, []byte("200 OK"), []byte(status))
		})
	}
}
