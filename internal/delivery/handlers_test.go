package delivery

import (
	"net/http"
	"testing"
)

func Test_handleError(t *testing.T) {
	type args struct {
		w           http.ResponseWriter
		statusError int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleError(tt.args.w, tt.args.statusError)
		})
	}
}

func TestHome(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Home(tt.args.w, tt.args.r)
		})
	}
}

func TestResult(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Result(tt.args.w, tt.args.r)
		})
	}
}
