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
		"Cache-Control": arrStrings("public, max-age=120"),
		"Expires":       arrStrings("Wed, 11 Jan 2023 12:08:23 GMT"),
	}

	result := NewBaseResult(h)
	assert.Equal(t, secs(120), result.CacheMaxAge)
	assert.Equal(t, time.Date(2023, time.January, 11, 12, 8, 23, 0, time.UTC), result.CacheExpires)
}

func Test_toNextPageIndex(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil}, -1},
		{"happy path", args{arrStrings("\u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=63\u003e; rel=\"last\", \u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=2\u003e; rel=\"next\"")}, 2},
		{"not defined", args{arrStrings("\u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=63\u003e; rel=\"last\"")}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, toNextPageIndex(tt.args.value), "toNextPageIndex(%v)", tt.args.value)
		})
	}
}

func Test_toLastPageIndex(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil}, -1},
		{"happy path", args{arrStrings("\u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=63\u003e; rel=\"last\", \u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=2\u003e; rel=\"next\"")}, 63},
		{"not defined", args{arrStrings("<https://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true&include_exchange_logo=true&order=volume_desc&page=2>; rel=\"next\"")}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, toLastPageIndex(tt.args.value), "toLastPageIndex(%v)", tt.args.value)
		})
	}
}

func TestNewBasePageResult(t *testing.T) {
	h := http.Header{
		"Cache-Control": arrStrings("public, max-age=120"),
		"Expires":       arrStrings("Wed, 11 Jan 2023 12:08:23 GMT"),
		"Link":          arrStrings("\u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=63\u003e; rel=\"last\", \u003chttps://api.coingecko.com/api/v3/coins/bitcoin/tickers?depth=true\u0026include_exchange_logo=true\u0026order=volume_desc\u0026page=2\u003e; rel=\"next\""),
		"Per-Page":      arrStrings("100"),
		"Total":         arrStrings("6247"),
	}

	result := NewBasePageResult(h)
	assert.Equal(t, secs(120), result.CacheMaxAge)
	assert.Equal(t, time.Date(2023, time.January, 11, 12, 8, 23, 0, time.UTC), result.CacheExpires)
	assert.Equal(t, 2, result.NextPageIndex)
	assert.Equal(t, 63, result.LastPageIndex)
	assert.Equal(t, 6247, result.TotalEntriesCount)
}
