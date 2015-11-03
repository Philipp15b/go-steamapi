package steamapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	AppID          uint   `json:",string"`
	ContextID      uint64 `json:",string"`
	AssetID        uint64 `json:",string"`
	CurrencyID     uint64 `json:",string"`
	ClassID        uint64 `json:",string"`
	InstanceID     uint64 `json:",string"`
	Amount         uint64 `json:",string"`
	Missing        bool
	MarketHashName string
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
	TradeID        uint64        `json:"tradeid,string"`
}

// CEconTradeOffers represent the list of different tradeoffers types
type CEconTradeOffers struct {
	Sent     []*CEconTradeOffer `json:"trade_offers_sent"`
	Received []*CEconTradeOffer `json:"trade_offers_received"`
}

type ieconGetTradeOffersResponse struct {
	Response struct {
		CEconTradeOffers
	}
}

// IEconGetTradeOffers retrieves a list of tradeoffers
func IEconGetTradeOffers(
	baseSteamAPIURL string,
	apiKey string,
	getSentOffers bool,
	getReceivedOffers bool,
	getDescriptions bool,
	activeOnly bool,
	historicalOnly bool,
	timeHistoricalCutoff int64,
) (*CEconTradeOffers, error) {

	querystring := url.Values{}
	querystring.Add("key", apiKey)
	querystring.Add("get_sent_offers", boolToStr(getSentOffers))
	querystring.Add("get_received_offers", boolToStr(getReceivedOffers))
	querystring.Add("get_descriptions", boolToStr(getDescriptions))
	querystring.Add("language", "en")
	querystring.Add("active_only", boolToStr(activeOnly))
	querystring.Add("historical_only", boolToStr(historicalOnly))
	querystring.Add("time_historical_cutoff", strconv.FormatInt(timeHistoricalCutoff, 10))

	resp, err := http.Get(baseSteamAPIURL + "/IEconService/GetTradeOffers/v0001?" + querystring.Encode())

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffers http.Get: error %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffers http.Get: http status %v", resp.Status)
	}

	defer resp.Body.Close()

	tosResp := &ieconGetTradeOffersResponse{}
	err = json.NewDecoder(resp.Body).Decode(tosResp)

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffers Decode: error %v", err)
	}

	return &tosResp.Response.CEconTradeOffers, nil
}

type ieconGetTradeOfferResponse struct {
	Response struct {
		Offer        CEconTradeOffer
		Descriptions []ItemDescription
	}
}

// ItemDescription represents the details about the items unique w classid instanceid
type ItemDescription struct {
	AppID          uint   `json:"appid"`
	ClassID        uint64 `json:"classid,string"`
	InstanceID     uint64 `json:"instanceid,string"`
	MarketHashName string `json:"market_hash_name"`
	IconURL        string `json:"icon_url"`
	NameColor      string `json:"name_color"`
	Name           string `json:"name"`
}

func findMarketHashName(itemD []ItemDescription, appID uint, classID, instanceID uint64) string {
	for _, description := range itemD {
		if description.AppID == appID &&
			description.ClassID == classID &&
			description.InstanceID == instanceID {
			return description.MarketHashName
		}
	}

	return ""
}

// IEconGetTradeOffer retrieves details about a specific tradeoffer
func IEconGetTradeOffer(baseSteamAPIURL string, apiKey string, steamID uint64, tradeOfferID uint64) (
	*CEconTradeOffer, error,
) {

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
	if resp.StatusCode != http.StatusOK {
		body, errBody := ioutil.ReadAll(resp.Body)
		return nil,
			fmt.Errorf("tradeoffer IEconGetTradeOffer: steam responded with a status %d with the message: %s (%v)",
				resp.StatusCode,
				body,
				errBody,
			)
	}

	defer resp.Body.Close()

	toResp := ieconGetTradeOfferResponse{}
	err = json.NewDecoder(resp.Body).Decode(&toResp)

	if err != nil {
		return nil, fmt.Errorf("tradeoffer IEconGetTradeOffer Decode: error %v", err)
	}

	// If the state is 0, it means there is a mistake
	if toResp.Response.Offer.State == 0 {
		body, errBody := ioutil.ReadAll(resp.Body)
		return nil,
			fmt.Errorf("tradeoffer IEconGetTradeOffer: steam responded with a status %d with the message: %s (%v)",
				resp.StatusCode,
				body,
				errBody,
			)
	}

	for giveIndex, asset := range toResp.Response.Offer.ToGive {
		toResp.Response.Offer.ToGive[giveIndex].MarketHashName =
			findMarketHashName(toResp.Response.Descriptions, asset.AppID, asset.ClassID, asset.InstanceID)
	}

	for receiveIndex, asset := range toResp.Response.Offer.ToReceive {
		toResp.Response.Offer.ToReceive[receiveIndex].MarketHashName =
			findMarketHashName(toResp.Response.Descriptions, asset.AppID, asset.ClassID, asset.InstanceID)
	}

	return &toResp.Response.Offer, nil
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

	if err != nil {
		return fmt.Errorf("tradeoffer IEconGetTradeOffer http.Get: error %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, errBody := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("tradeoffer IEcon%sTradeOffer: steam responded with a status %d with the message: %s (%v)",
			action,
			resp.StatusCode,
			body,
			errBody,
		)
	}

	err = resp.Body.Close()

	if err != nil {
		return fmt.Errorf("tradeoffer IEcon%sTradeOffer resp.Body.Close(): error %v", action, err)
	}

	return nil

}

func boolToStr(b bool) string {
	if b {
		return "1"
	}

	return "0"

}
