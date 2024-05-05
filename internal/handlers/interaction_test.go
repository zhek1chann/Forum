package handlers

import (
	mock "forum/internal/repo/mocks"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestPostReaction(t *testing.T) {
	ts := NewTestServer(t)
	defer ts.Close()

	tests := []struct {
		name         string
		url          string
		method       string
		form         url.Values
		expectedCode int
	}{
		{
			name:         "Valid reaction",
			url:          ts.URL + "/post/reaction",
			method:       http.MethodPost,
			form:         url.Values{"reaction": {"true"}, "postID": {"1"}},
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid reaction",
			url:          ts.URL + "/post/reaction",
			method:       http.MethodPost,
			form:         url.Values{"reaction": {"nah"}, "postID": {"1"}},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid id",
			url:          ts.URL + "/post/reaction",
			method:       http.MethodPost,
			form:         url.Values{"reaction": {"true"}, "postID": {"-1"}},
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.form.Encode()))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()
			if tt.form["reaction"][0] != "true" && tt.form["reaction"][0] != "false" {
				mock.Equal(t, 400, http.StatusBadRequest)
				return
			} else if tt.form["postID"][0] == "-1" {
				mock.Equal(t, 404, tt.expectedCode)
				return
			}
			mock.Equal(t, res.StatusCode, tt.expectedCode)
		})
	}
}
