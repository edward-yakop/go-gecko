package types

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_toExpires(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"nil", args{nil}, time.Time{}},
		{"happy path", args{arrStrings("Wed, 11 Jan 2023 12:08:21 GMT")}, time.Date(2023, time.January, 11, 12, 8, 21, 0, time.UTC)},
		{"invalid string", args{arrStrings("whatever")}, time.Time{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toExpires(tt.args.value)

			assert.Equalf(t, tt.want, got, "toExpires(%v)", tt.args.value)
		})
	}
}

func arrStrings(v ...string) []string {
	return v
}

func Test_toMaxAge(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"nil", args{nil}, time.Nanosecond},
		{"happy path", args{arrStrings("public, max-age=120")}, secs(120)},
		{"invalid string", args{arrStrings("invalid string")}, time.Nanosecond},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, toMaxAge(tt.args.value), "toAge(%v)", tt.args.value)
		})
	}
}

func secs(i time.Duration) time.Duration {
	return time.Second * i
}

func TestNewBaseResult(t *testing.T) {
	h := http.Header{
		"Age":           arrStrings("109"),
		"Cache-Control": arrStrings("public, max-age=120"),
		"Expires":       arrStrings("Wed, 11 Jan 2023 12:08:23 GMT"),
	}

	result := NewBaseResult(h)
	assert.Equal(t, secs(109), result.Age)
	assert.Equal(t, secs(120), result.MaxAge)
	assert.Equal(t, time.Date(2023, time.January, 11, 12, 8, 23, 0, time.UTC), result.Expires)
}
