package format

import "testing"

func TestBool2String(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"true", args{true}, "true"},
		{"false", args{false}, "false"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool2String(tt.args.b); got != tt.want {
				t.Errorf("Bool2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt2String(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{0}, "0"},
		{"10", args{10}, "10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int2String(tt.args.i); got != tt.want {
				t.Errorf("Int2String() = %v, want %v", got, tt.want)
			}
		})
	}
}
