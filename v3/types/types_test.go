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
		{"invalid: nil", args{nil}, time.Time{}},
		{"valid: happy path", args{arrStrings("Wed, 11 Jan 2023 12:08:21 GMT")}, time.Date(2023, time.January, 11, 12, 8, 21, 0, time.UTC)},
		{"invalid: not expected format", args{arrStrings("whatever")}, time.Time{}},
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
		{"invalid: nil", args{nil}, time.Nanosecond},
		{"valid: happy path", args{arrStrings("public, max-age=120")}, secs(120)},
		{"invalid: not int", args{arrStrings("invalid string")}, time.Nanosecond},
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
		"Cache-Control": arrStrings("public, max-age=120"),
		"Expires":       arrStrings("Wed, 11 Jan 2023 12:08:23 GMT"),
	}

	result := NewBaseResult(h)
	assert.Equal(t, secs(120), result.CacheMaxAge)
	assert.Equal(t, time.Date(2023, time.January, 11, 12, 8, 23, 0, time.UTC), result.CacheExpires)
}

func TestNewBasePageResult(t *testing.T) {
	h := http.Header{
		"Cache-Control": arrStrings("public, max-age=120"),
		"Expires":       arrStrings("Wed, 11 Jan 2023 12:08:23 GMT"),
		"Per-Page":      arrStrings("100"),
		"Total":         arrStrings("6247"),
	}
	type args struct {
		currentPageIndex int
	}
	tests := []struct {
		name              string
		args              args
		wantNextPageIndex int
	}{
		{"invalid: before first page", args{-10}, 2},
		{"valid: non-last page", args{1}, 2},
		{"valid: last page", args{63}, -1},
		{"invalid: beyond last page", args{64}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBasePageResult(h, tt.args.currentPageIndex)

			assert.Equal(t, secs(120), got.CacheMaxAge)
			assert.Equal(t, time.Date(2023, time.January, 11, 12, 8, 23, 0, time.UTC), got.CacheExpires)
			assert.Equal(t, tt.wantNextPageIndex, got.NextPageIndex)
			assert.Equal(t, 63, got.LastPageIndex)
			assert.Equal(t, 6247, got.TotalEntriesCount)
		})
	}
}

func Test_toInt(t *testing.T) {
	type args struct {
		value        []string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"invalid: nil", args{nil, -1}, -1},
		{"valid: 1", args{arrStrings("1"), -1}, 1},
		{"valid: 100", args{arrStrings("100"), -1}, 100},
		{"invalid: non int", args{arrStrings("ABC"), -3}, -3},
		{"invalid: float", args{arrStrings("123.3"), -2}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, toInt(tt.args.value, tt.args.defaultValue), "toInt(%v, %v)", tt.args.value, tt.args.defaultValue)
		})
	}
}
