package steamapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIEconGetTradeOffer(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, GetMockActiveStateIEconGetTradeOffer())
		}),
	)

	defer ts.Close()

	expectedItem := CEconAsset{
		AppID:          123,
		ContextID:      2,
		AssetID:        1234553,
		CurrencyID:     0,
		ClassID:        888881,
		InstanceID:     0,
		Amount:         1,
		Missing:        false,
		MarketHashName: "testmkt",
	}

	expectedCETO := CEconTradeOffer{
		TradeOfferID:   123456,
		OtherAccountID: 1234,
		Message:        "",
		ExpirationTime: 1300000000,
		State:          2,
		ToGive:         []*CEconAsset{},
		ToReceive:      []*CEconAsset{&expectedItem},
		IsOurs:         true,
		TimeCreated:    1300000000,
		TimeUpdated:    1300000000,
	}

	TOgot, err := IEconGetTradeOffer(ts.URL, "123", 1, 1)

	if err != nil {
		t.Errorf("IEconGetTradeOffer unexpected err %v", err)
		return
	}

	if TOgot.TradeOfferID != expectedCETO.TradeOfferID ||
		TOgot.OtherAccountID != expectedCETO.OtherAccountID ||
		TOgot.Message != expectedCETO.Message ||
		TOgot.ExpirationTime != expectedCETO.ExpirationTime ||
		TOgot.State != expectedCETO.State ||
		TOgot.IsOurs != expectedCETO.IsOurs ||
		TOgot.TimeCreated != expectedCETO.TimeCreated ||
		TOgot.TimeUpdated != expectedCETO.TimeUpdated ||
		len(TOgot.ToGive) != len(expectedCETO.ToGive) ||
		len(TOgot.ToReceive) != len(expectedCETO.ToReceive) {
		t.Errorf("IEconGetTradeOffer expected %v, got %v", expectedCETO, TOgot)
	}

	if *TOgot.ToReceive[0] != *expectedCETO.ToReceive[0] {
		t.Errorf("IEconGetTradeOffer expected %v, got %v", expectedCETO, TOgot)
	}
}

func TestIEconGetTradeOffers(t *testing.T) {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, GetMockIEconGetTradeOffers())
		}),
	)

	defer ts.Close()

	expectedItem := CEconAsset{
		AppID:          123,
		ContextID:      2,
		AssetID:        1234553,
		CurrencyID:     0,
		ClassID:        888881,
		InstanceID:     0,
		Amount:         1,
		Missing:        false,
		MarketHashName: "testmkt",
	}

	expectedCETOsSent := []*CEconTradeOffer{
		&CEconTradeOffer{
			TradeOfferID:   123456,
			OtherAccountID: 1234,
			Message:        "",
			ExpirationTime: 1300000000,
			State:          2,
			ToGive:         []*CEconAsset{},
			ToReceive:      []*CEconAsset{&expectedItem},
			IsOurs:         true,
			TimeCreated:    1300000000,
			TimeUpdated:    1300000000,
		},
		&CEconTradeOffer{
			TradeOfferID:   123457,
			OtherAccountID: 1234,
			Message:        "",
			ExpirationTime: 1300000000,
			State:          2,
			ToGive:         []*CEconAsset{},
			ToReceive:      []*CEconAsset{&expectedItem},
			IsOurs:         true,
			TimeCreated:    1300000000,
			TimeUpdated:    1300000000,
		},
	}

	expectedCETOsReceived := []*CEconTradeOffer{
		&CEconTradeOffer{
			TradeOfferID:   123458,
			OtherAccountID: 12345,
			Message:        "",
			ExpirationTime: 1300000000,
			State:          2,
			ToGive:         []*CEconAsset{},
			ToReceive:      []*CEconAsset{&expectedItem},
			IsOurs:         true,
			TimeCreated:    1300000000,
			TimeUpdated:    1300000000,
		},
	}

	TOsGot, err := IEconGetTradeOffers(ts.URL, "123", true, true, true, false, false, 1)

	if err != nil {
		t.Errorf("IEconGetTradeOffers unexpected err %v", err)
		return
	}

	if len(TOsGot.Sent) != len(expectedCETOsSent) ||
		len(TOsGot.Received) != len(expectedCETOsReceived) {
		t.Errorf("IEconGetTradeOffers expected %d offers sent, %d received "+
			"got %d offers sent, %d received ",
			len(expectedCETOsSent), len(expectedCETOsReceived),
			len(TOsGot.Sent), len(TOsGot.Received),
		)
	}

	if TOsGot.Received[0].OtherAccountID != expectedCETOsReceived[0].OtherAccountID {
		t.Errorf("IEconGetTradeOffers expected accountid %d, got %d",
			expectedCETOsReceived[0].OtherAccountID, TOsGot.Received[0].OtherAccountID)
	}
}
