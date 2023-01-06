package coingecko

import (
	"github.com/h2non/gock"
	"io"
	"net/http"
	"os"
)

func init() {
	defer gock.Off()
}

var c = NewClient(nil)
var mockURL = "https://api.coingecko.com/api/v3"

// Util: Setup Gock
func setupGock(filename string, url string) error {
	testJSON, err := os.Open(filename)
	defer func(testJSON *os.File) {
		_ = testJSON.Close()
	}(testJSON)

	if err != nil {
		return err
	}
	testByte, err := io.ReadAll(testJSON)
	if err != nil {
		return err
	}
	gock.New(mockURL).
		Get(url).
		Reply(http.StatusOK).
		JSON(testByte)

	return nil
}
