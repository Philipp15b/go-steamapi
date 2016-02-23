package steamapi

import (
	"net/url"
	"strconv"
)

type schemaJSON struct {
	Result Schema
}

// Schema is a game schema
type Schema struct {
	Status                               int
	FullSchemaURL                        string `json:"items_game_url"`
	Qualities                            map[string]int
	QualityNames                         map[string]string
	OriginNames                          []Origin
	Items                                []SchemaItem
	Attributes                           []SchemaAttribute
	ItemSets                             []ItemSet
	AttributeControlledAttachedParticles []ParticleEffect     `json:"attribute_controlled_attached_particles"`
	ItemLevels                           []ItemRankSet        `json:"item_levels"`
	KillEaterScoreTypes                  []KillEaterScoreType `json:"kill_eater_score_types"`
}

// Item finds an item by its defindex in a Schema.
func (s *Schema) Item(defindex int) *SchemaItem {
	for _, item := range s.Items {
		if item.Defindex == defindex {
			return &item
		}
	}
	return nil
}

// SchemaItem is a schema description of an Item.
type SchemaItem struct {
	Name               string
	Defindex           int
	ItemClass          string                `json:"item_class"`
	ItemTypeName       string                `json:"item_type_name"`
	ItemName           string                `json:"item_name"`
	Description        string                `json:"item_description,omitempty"`
	ProperName         bool                  `json:"proper_name"`
	Slot               string                `json:"item_slot"`
	DefaultQuality     int                   `json:"item_quality"`
	InventoryImage     *string               `json:"image_inventory"` // this is null for the "Free Trial Premium Upgrade"
	ImageURL           string                `json:"image_url"`
	ImageURLLarge      string                `json:"image_url_large"`
	DropType           string                `json:"drop_type,omitempty"`
	ItemSet            string                `json:"item_set,omitempty"`
	HolidayRestriction string                `json:"holiday_restriction"`
	MinLevel           int                   `json:"min_ilevel"`
	MaxLevel           int                   `json:"max_ilevel"`
	CraftClass         string                `json:"craft_class,omitempty"`
	Capabilities       map[string]bool       `json:",omitempty"`
	UsedByClasses      []string              `json:"used_by_classes,omitempty"`
	ClassLoadoutSlots  map[string]string     `json:"per_class_loadout_slots,omitempty"`
	Styles             []Style               `json:",omitempty"`
	Attributes         []SchemaItemAttribute `json:",omitempty"`
}

// Origin is an item origin as defined in the Schema.
type Origin struct {
	ID   int `json:"origin"`
	Name string
}

// Style is an item style
type Style struct {
	Name string
}

// SchemaItemAttribute is the schema items attributes
type SchemaItemAttribute struct {
	Name  string
	Class string
	Value float64
}

// SchemaAttribute is a schema attribute
type SchemaAttribute struct {
	Name              string
	Defindex          int
	Class             string `json:"attribute_class"`
	MinValue          float64
	MaxValue          float64
	Description       string `json:"description_string,omitempty"`
	DescriptionFormat string `json:"description_format,omitempty"`
	EffectType        string `json:"effect_type"`
	Hidden            bool
	StoredAsInteger   bool `json:"stored_as_integer"`
}

// ItemSet is an item set
type ItemSet struct {
	InternalName string `json:"item_set"`
	Name         string
	StoreBundle  string `json:"store_bundle,omitempty"`
	Items        []string
	Attributes   []SchemaItemAttribute `json:",omitempty"`
}

// ParticleEffect is the particle effect
type ParticleEffect struct {
	System           string
	ID               int
	AttachToRootbone bool   `json:"attach_to_rootbone"`
	Attachment       string `json:",omitempty"`
	Name             string
}

// ItemRankSet is the set of the possible ranks
type ItemRankSet struct {
	Name   string
	Levels []ItemRank
}

// ItemRank is the rank of the item
type ItemRank struct {
	Level         int
	RequiredScore int `json:"required_score"`
	Name          string
}

// KillEaterScoreType is the type of kill eater score
type KillEaterScoreType struct {
	ID   int    `json:"type"`
	Name string `json:"type_name"`
}

// GetSchema Fetches the Schema for the given game.
func GetSchema(appID int, language, APIKey string) (*Schema, error) {
	getSchema := NewSteamMethod("IEconItems_"+strconv.Itoa(appID), "GetSchema", 1)

	vals := url.Values{}
	vals.Add("key", APIKey)
	vals.Add("language", language)

	var resp schemaJSON
	err := getSchema.Request(vals, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
