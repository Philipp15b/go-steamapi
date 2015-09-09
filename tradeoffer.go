package tradeoffer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ieconGetTradeOffersResponse struct {
	Response CEconTradeOffers `json:"response"`
}

// IEconGetTradeOffers retrieves a list of tradeoffers
func IEconGetTradeOffers(baseSteamAPIURL string, apiKey string) (*CEconTradeOffers, error) {

	tosResp := &ieconGetTradeOffersResponse{}

	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("get_sent_offers", "1")
	querystring.Add("get_received_offers", "1")
	querystring.Add("get_descriptions", "1")
	querystring.Add("language", "en")
	querystring.Add("active_only", "0")
	querystring.Add("historical_only", "0")
	querystring.Add("time_historical_cutoff", "1")

	resp, err := http.Get(baseSteamAPIURL + "/IEconService/GetTradeOffers/v0001?" + querystring.Encode())

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffers http.Get: error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffers http.Get: http status %v", resp.Status)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(tosResp)

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffers Decode: error %v", err)
	}

	return &tosResp.Response, nil
}

type ieconGetTradeOfferResponse struct {
	Response CEconTradeOffer `json:"response"`
}

// IEconGetTradeOffer retrieves details about a specific tradeoffer
func IEconGetTradeOffer(baseSteamAPIURL string, apiKey string, steamID uint64, tradeOfferID SteamTradeOfferID) (
	*CEconTradeOffer, error,
) {

	toResp := &ieconGetTradeOfferResponse{}

	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("steamid", steamID.String())
	querystring.Add("format", "json")
	querystring.Add("tradeofferid", tradeOfferID.String())
	querystring.Add("language", "en")

	resp, err := http.Get(baseSteamAPIURL + "/IEconService/GetTradeOffer/v0001?" + querystring.Encode())

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffer http.Get: error %v", err)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(toResp)

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffer Decode: error %v", err)
	}

	return &toResp.Response, nil
}

// IEconActionTradeOffer declines a TO created by someone else
func IEconActionTradeOffer(baseSteamAPIURL string, action string, apiKey string, tradeOfferID SteamTradeOfferID) error {

	if action != "Decline" && action != "Cancel" {
		return fmt.Errorf("tradeoffer IEconActionTradeOffer doesn't support %v action", action)
	}
	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("tradeofferid", tradeOfferID.String())

	resp, err := http.Get(
		baseSteamAPIURL + "/IEconService/" + action + "TradeOffer/v0001?" + querystring.Encode())

	if resp.StatusCode != 200 || err != nil {
		return fmt.Errorf("tradeoffer IEcon%sTradeOffer http.Get: %v error %v", action, resp.StatusCode, err)
	}

	err = resp.Body.Close()

	if err != nil {
		return fmt.Errorf("tradeoffer IEcon%sTradeOffer resp.Body.Close(): error %v", action, err)
	}

	return nil

}
