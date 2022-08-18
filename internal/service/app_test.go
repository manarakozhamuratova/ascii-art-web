package service

import "testing"

func TestApp(t *testing.T) {
	type args struct {
		message string
		banner  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := App(tt.args.message, tt.args.banner)
			if (err != nil) != tt.wantErr {
				t.Errorf("App() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("App() = %v, want %v", got, tt.want)
			}
		})
	}
}
