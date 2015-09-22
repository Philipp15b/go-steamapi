package steamapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/vincentserpoul/mangosteam/inventory"
)

// State represents the state of the tradeoffer, see constants
type State uint

const (
	// ETradeOfferStateCreated /!\ non steam status, used to know the TO has been created
	ETradeOfferStateCreated State = iota
	// ETradeOfferStateInvalid Invalid
	ETradeOfferStateInvalid
	// ETradeOfferStateActive This trade offer has been sent, neither party has acted on it yet.
	ETradeOfferStateActive
	// ETradeOfferStateAccepted The trade offer was accepted by the recipient and items were exchanged.
	ETradeOfferStateAccepted
	// ETradeOfferStateCountered The recipient made a counter offer
	ETradeOfferStateCountered
	// ETradeOfferStateExpired The trade offer was not accepted before the expiration date
	ETradeOfferStateExpired
	// ETradeOfferStateCanceled The sender cancelled the offer
	ETradeOfferStateCanceled
	// ETradeOfferStateDeclined The recipient declined the offer
	ETradeOfferStateDeclined
	// ETradeOfferStateInvalidItems Some of the items in the offer are no longer available
	// (indicated by the missing flag in the output)
	ETradeOfferStateInvalidItems
	// ETradeOfferStateEmailPending The offer hasn't been sent yet and is awaiting email confirmation
	ETradeOfferStateEmailPending
	// ETradeOfferStateEmailCanceled The receiver cancelled the offer via email
	ETradeOfferStateEmailCanceled
)

// CEconAsset represents an asset in steam web api
type CEconAsset struct {
	AppID      uint   `json:",string"`
	ContextID  uint64 `json:",string"`
	AssetID    uint64 `json:",string"`
	CurrencyID uint64 `json:",string"`
	ClassID    uint64 `json:",string"`
	InstanceID uint64 `json:",string"`
	Amount     uint64 `json:",string"`
	Missing    bool
}

// CEconTradeOffer represent the to from the steam API
type CEconTradeOffer struct {
	TradeOfferID   uint64 `json:",string"`
	OtherAccountID uint64 `json:"accountid_other"`
	Message        string
	ExpirationTime uint32        `json:"expiration_time"`
	State          State         `json:"trade_offer_state"`
	ToGive         []*CEconAsset `json:"items_to_give"`
	ToReceive      []*CEconAsset `json:"items_to_receive"`
	IsOurs         bool          `json:"is_our_offer"`
	TimeCreated    uint32        `json:"time_created"`
	TimeUpdated    uint32        `json:"time_updated"`
}

// CEconTradeOffers contains a list of tradeoffers, sent and received
type CEconTradeOffers struct {
	Sent         []*CEconTradeOffer     `json:"trade_offers_sent"`
	Received     []*CEconTradeOffer     `json:"trade_offers_received"`
	Descriptions inventory.Descriptions `json:"descriptions"`
}

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
func IEconGetTradeOffer(baseSteamAPIURL string, apiKey string, steamID uint64, tradeOfferID uint64) (
	*CEconTradeOffer, error,
) {

	toResp := &ieconGetTradeOfferResponse{}

	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("steamid", strconv.FormatUint(steamID, 10))
	querystring.Add("format", "json")
	querystring.Add("tradeofferid", strconv.FormatUint(tradeOfferID, 10))
	querystring.Add("language", "en")

	resp, err := http.Get(baseSteamAPIURL + "/IEconService/GetTradeOffer/v1?" + querystring.Encode())

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
func IEconActionTradeOffer(baseSteamAPIURL string, action string, apiKey string, tradeOfferID uint64) error {

	if action != "Decline" && action != "Cancel" {
		return fmt.Errorf("tradeoffer IEconActionTradeOffer doesn't support %v action", action)
	}
	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("tradeofferid", strconv.FormatUint(tradeOfferID, 10))

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
