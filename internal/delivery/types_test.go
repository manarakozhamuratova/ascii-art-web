package delivery

import "testing"

func Test_input_validate(t *testing.T) {
	type fields struct {
		text   string
		banner string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &input{
				text:   tt.fields.text,
				banner: tt.fields.banner,
			}
			if got := i.validate(); got != tt.want {
				t.Errorf("input.validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
