package steamapi

import (
	"errors"
	"net/url"
	"strconv"
)

type storeJSON struct {
	Result struct {
		Success bool
		Assets  []Asset
	}
}

// Asset is an item in the store.
type Asset struct {
	Prices   map[string]int
	Defindex int `json:"name,string"`
	Date     string
	Tags     []string
	TagIDs   []int64
}

// HasTag return bool if the asset has a tag
func (i *Asset) HasTag(tag string) bool {
	for _, t := range i.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// GetAssetPrices returns a list of assets with their prices
func GetAssetPrices(appid uint64, language, currency, apiKey string) ([]Asset, error) {

	var getAssetPrices = NewSteamMethod("ISteamEconomy", "GetAssetPrices", 1)

	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("appid", strconv.FormatUint(appid, 10))
	vals.Add("language", language)
	vals.Add("currency", currency)

	var resp storeJSON
	err := getAssetPrices.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Result.Success {
		return nil, errors.New("API call 'GetAssetPrices' did not succeed")
	}
	return resp.Result.Assets, nil
}
