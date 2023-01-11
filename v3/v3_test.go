package coingecko

import (
	"encoding/json"
	"fmt"
	"github.com/edward-yakop/go-gecko/v3/types"
	"github.com/h2non/gock"
	"net/http"
	"os"
	"time"
)

func init() {
	defer gock.Off()
}

var c = NewClient(nil)
var mockURL = "https://api.coingecko.com/api/v3"

// Util: Setup Gock
func setupGock(filename string, url string) error {
	return setupGockWithHeader(filename, "", url)
}

func setupGockWithHeader(bodyFileName, headerFileName, url string) error {
	bodyBA, err := os.ReadFile(bodyFileName)
	if err != nil {
		return fmt.Errorf("fail to read %s: %v", bodyFileName, err)
	}

	g := gock.New(mockURL).
		Get(url).
		Reply(http.StatusOK).
		JSON(bodyBA)

	if headerFileName != "" {
		if headerBA, hErr := os.ReadFile(headerFileName); hErr == nil {
			if jErr := json.Unmarshal(headerBA, &g.Header); jErr != nil {
				return fmt.Errorf("fail to unmarshal json [%s]: %v", headerFileName, jErr)
			}
		} else {
			return fmt.Errorf("fail to read header [%s] file: %v", headerFileName, hErr)
		}
	}

	return nil
}

func secs(i time.Duration) time.Duration {
	return time.Second * i
}

func baseResult(age, maxAge time.Duration, expires time.Time) types.BaseResult {
	return types.BaseResult{
		Age:     secs(age),
		MaxAge:  secs(maxAge),
		Expires: expires,
	}
}
