package steamapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMockOkGetPlayerItems(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetMockOKGetPlayerItems())
	}))
	defer ts.Close()

	appID := uint64(2)
	steamID := uint64(1234)
	apiKey := "123"

	expectedPlayerItems := Inventory{
		Status:        uint(1),
		BackpackSlots: 840,
		Items: []Item{
			Item{
				ID:                uint64(1234567894),
				OriginalID:        uint64(1234567894),
				Defindex:          5470,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       false,
				Uncraftable:       false,
				InventoryToken:    uint64(0x0),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes:        []Attribute(nil),
				Equipped: []EquipInfo{
					EquipInfo{
						Class: 23,
						Slot:  6,
					},
				},
			},
			Item{
				ID:                uint64(1234567895),
				OriginalID:        uint64(1234567895),
				Defindex:          15001,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       true,
				Uncraftable:       false,
				InventoryToken:    uint64(0x0),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes: []Attribute{
					Attribute{
						Defindex:    153,
						Value:       1065353216,
						FloatValue:  float64(0),
						AccountInfo: (*AccountInfo)(nil),
					},
					Attribute{
						Defindex:    16,
						Value:       1,
						FloatValue:  float64(0),
						AccountInfo: (*AccountInfo)(nil),
					},
				},
				Equipped: []EquipInfo(nil),
			},
			Item{
				ID:                uint64(1234567897),
				OriginalID:        uint64(1234567897),
				Defindex:          10068,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       true,
				Uncraftable:       false,
				InventoryToken:    uint64(0x0),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes: []Attribute{
					Attribute{
						Defindex:    153,
						Value:       1065353216,
						FloatValue:  float64(0),
						AccountInfo: (*AccountInfo)(nil),
					},
				},
				Equipped: []EquipInfo(nil),
			},
			Item{
				ID:                uint64(1234567898),
				OriginalID:        uint64(1234567898),
				Defindex:          5508,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       false,
				Uncraftable:       false,
				InventoryToken:    uint64(0x0),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes:        []Attribute(nil),
				Equipped: []EquipInfo{
					EquipInfo{
						Class: 4,
						Slot:  4,
					},
				},
			},
			Item{
				ID:                0x499602da,
				OriginalID:        0x499602da,
				Defindex:          7480,
				Level:             1,
				Quantity:          1,
				Origin:            0,
				Untradeable:       true,
				Uncraftable:       false,
				InventoryToken:    uint64(0x0),
				Quality:           4,
				CustomName:        "",
				CustomDescription: "",
				Attributes: []Attribute{
					Attribute{
						Defindex:    153,
						Value:       1,
						FloatValue:  float64(0),
						AccountInfo: (*AccountInfo)(nil),
					},
					Attribute{
						Defindex:    213,
						Value:       1,
						FloatValue:  float64(0),
						AccountInfo: (*AccountInfo)(nil),
					},
				},
				Equipped: []EquipInfo{
					EquipInfo{
						Class: 15,
						Slot:  0,
					},
				},
			},
		},
	}

	playerItems, err := GetPlayerItems(ts.URL, steamID, appID, apiKey)

	if err != nil {
		t.Errorf("GetPlayerItems failure: %v", err)
	}

	// Need a better test
	if playerItems.Status != expectedPlayerItems.Status ||
		playerItems.BackpackSlots != expectedPlayerItems.BackpackSlots ||
		len(playerItems.Items) != len(expectedPlayerItems.Items) {

		t.Errorf("GetPlayerItems(%v, %v, %v, %v) == %#v, expected %#v",
			ts.URL, steamID, appID, apiKey, playerItems, expectedPlayerItems)
	}

}
