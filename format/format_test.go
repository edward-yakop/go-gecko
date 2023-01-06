package format

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
			got := Bool2String(tt.args.b)

			assert.Equalf(t, tt.want, got, "Bool2String(%v)", tt.args.b)
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
			got := Int2String(tt.args.i)

			assert.Equalf(t, tt.want, got, "Int2String(%v)", tt.args.i)
		})
	}
}
