package steamapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMockOkGetAssetClassInfo(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetMockOKGetAssetClassInfo())
	}))
	defer ts.Close()

	expectedInfo := Info{
		ClassID:        "123456789",
		IconURL:        "W_I_5GLm4wPcv9jJQ7z7tz_l_0sEIYUhRfbF4arNQkgGQGKd3kMuVpMgCwRZrhSfeEqb1qNMeO7lDgsvJYj2VkHyNb-A-UWkTe9Xc8Rgd2sbj9_ugkgSUXffBrFHXNQrvM7K0Ay7XgXDLWdun9gFgPqagJWGCPPO6UywK3ID03w",
		MarketHashName: "Ye Olde Pipe",
		Tradable:       "1",
		Marketable:     "1",
	}

	appID := uint64(2)
	classID := uint64(1234)
	language := "en"
	apiKey := "123"

	infos, err := GetAssetClassInfo(ts.URL, appID, classID, language, apiKey)

	if err != nil {
		t.Errorf("GetAssetClassInfo failure: %v", err)
	}

	if *infos != expectedInfo {
		t.Errorf("GetAssetClassInfo(%v, %v, %v, %v, %v) == %#v, expected %#v",
			ts.URL, appID, classID, language, apiKey, infos, expectedInfo)
	}

}
