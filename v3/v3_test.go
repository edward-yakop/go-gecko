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
func setupGock(bodyFileName, headerFileName, url string) error {
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

func baseResult(maxAge time.Duration, expires time.Time) types.BaseResult {
	return types.BaseResult{
		CacheMaxAge:  secs(maxAge),
		CacheExpires: expires,
	}
}

var commonBaseResult = baseResult(120, time.Date(2023, time.January, 11, 12, 44, 47, 0, time.UTC))

func basePageResult(maxAge time.Duration, expires time.Time, nextPageIndex, lastPageIndex, pageSize, totalEntriesCount int) types.BasePageResult {
	return types.BasePageResult{
		BaseResult:        baseResult(maxAge, expires),
		NextPageIndex:     nextPageIndex,
		LastPageIndex:     lastPageIndex,
		PageSize:          pageSize,
		TotalEntriesCount: totalEntriesCount,
	}
}

var commonBasePageResult = basePageResult(
	120, time.Date(2023, time.January, 11, 12, 44, 47, 0, time.UTC),
	2, 63, 100, 6247,
)
