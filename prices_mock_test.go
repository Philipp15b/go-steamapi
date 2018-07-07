package steamapi

// GetMockOKGetAssetPrices gives def_index, class_id...
// https://api.steampowered.com/ISteamEconomy/GetAssetPrices/v1/?key=XXXXX&format=json&appid=XXX&currency=usd
func GetMockOKGetAssetPrices() string {
	return `
		{
			"result": {
				"success": true,
				"assets": [
					{
						"prices": {
							"USD": 0
						},
						"name": "4004",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4004"
							}
						]
						,
						"classid": "57939754"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4008",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4008"
							}
						]
						,
						"classid": "57939594"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4009",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4009"
							}
						]
						,
						"classid": "57939591"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4010",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4010"
							}
						]
						,
						"classid": "57939593"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4049",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4049"
							}
						]
						,
						"classid": "93966736"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4097",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4097"
							}
						]
						,
						"classid": "147888890"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "4110",
						"date": "1/13/2014",
						"class": [
							{
								"name": "def_index",
								"value": "4110"
							}
						]
						,
						"classid": "57939654"
					},
					{
						"prices": {
							"USD": 0
						},
						"name": "20886",
						"date": "9/2/2015",
						"class": [
							{
								"name": "def_index",
								"value": "20886"
							}
						]
						,
						"classid": "1218050854"
					}
				]		
			}
		}
	`
}
