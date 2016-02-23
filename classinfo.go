package steamapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// ClassInfo is the details of the specific classid
type classInfoJSON struct {
	Result map[string]json.RawMessage `json:"result"`
}

// Info is the details about the class info
type Info struct {
	ClassID        string `json:"classid"`
	IconURL        string `json:"icon_url"`
	MarketHashName string `json:"market_hash_name"`
	Tradable       string
	Marketable     string
}

// GetAssetClassInfo returns asset details
func GetAssetClassInfo(appID, classID uint64, language, apiKey string) (*Info, error) {

	var getAssetClassInfo = NewSteamMethod("ISteamEconomy", "GetAssetClassInfo", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("appid", strconv.FormatUint(appID, 10))
	vals.Add("language", language)
	vals.Add("class_count", "1")
	vals.Add("classid0", strconv.FormatUint(classID, 10))

	var resp classInfoJSON
	err := getAssetClassInfo.Request(vals, &resp)
	if err != nil {
		return nil, err
	}

	var info Info
	for _, object := range resp.Result {
		err := json.Unmarshal(object, &info)
		if err != nil {
			continue
		}
	}

	return &info, nil
}
