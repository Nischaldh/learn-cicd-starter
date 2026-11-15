package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyt(t *testing.T) {

	type res struct {
		auth string
		err  error
	}
	type test struct {
		name  string
		input http.Header
		want  res
	}
	tests := []test{
		{
			name: "Empty Header",
			input: http.Header{
				"Authorization": []string{""},
			},
			want: res{
				auth: "",
				err:  ErrNoAuthHeaderIncluded,
			},
		},
		{
			name: "Malformed Api key with len 1",
			input: http.Header{
				"Authorization": []string{"ApiKeywqer"},
			},
			want: res{
				auth: "",
				err:  errors.New("malformed authorization header"),
			},
		},
		{
			name: "Malformed Api key with len 2",
			input: http.Header{
				"Authorization": []string{"ALSJSDF sdlkfjs"},
			},
			want: res{
				auth: "",
				err:  errors.New("malformed authorization header"),
			},
		},
		{
			name: "Correct Api key",
			input: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			want: res{
				auth: "abc123",
				err:  nil,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotAuth, gotErr := GetAPIKey(tc.input)
			if gotAuth != tc.want.auth {
				t.Errorf("GetAPIKey() auth = %q, want %q", gotAuth, tc.want.auth)
			}

			if gotErr != nil && tc.want.err != nil {
				t.Errorf("expected error %v, got nil", tc.want.err)
			} else if gotErr != nil && tc.want.err == nil {
				t.Errorf("expected nil error, got %v", gotErr)
			} else if gotErr != nil && gotErr.Error() != tc.want.err.Error() {
				t.Errorf("GetAPIKey() error = %v, want %v", gotErr, tc.want.err)
			}

		})
	}
}
