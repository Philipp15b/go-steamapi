package steamapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestIEconGetTradeOffer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, getMockActiveStateIEconGetTradeOffer())
	}))
	defer ts.Close()

	expectedItem := CEconAsset{
		AppID:      123,
		ContextID:  2,
		AssetID:    1234553,
		CurrencyID: 0,
		ClassID:    888881,
		InstanceID: 0,
		Amount:     1,
		Missing:    false,
	}

	expectedCETO := CEconTradeOffer{
		TradeOfferID:   123456,
		OtherAccountID: 1234,
		Message:        "",
		ExpirationTime: 1300000000,
		State:          2,
		ToGive:         []*CEconAsset{&expectedItem},
		ToReceive:      []*CEconAsset{},
		IsOurs:         true,
		TimeCreated:    1300000000,
		TimeUpdated:    1300000000,
	}

	TOgot, err := IEconGetTradeOffer(ts.URL, "123", 1, 1)

	if err != nil {
		t.Errorf("IEconGetTradeOffer unexpected err %v", err)
		return
	}

	if !reflect.DeepEqual(TOgot, expectedCETO) {
		t.Errorf("IEconGetTradeOffer expected %v, got %v", expectedCETO, TOgot)
	}

}
