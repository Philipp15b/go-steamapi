package steamapi

// GetMockOKGetPlayerItems steam API mock response gives you item_id, def_index...
// https://api.steampowered.com/IEconItems_XXXX/GetPlayerItems/v1/?key=XXXXXX&format=json&steamid=XXXXX
func GetMockOKGetPlayerItems() string {
	return `
		{
			"result": {
				"status": 1,
				"num_backpack_slots": 840,
				"items": [
					{
						"id": 1234567894,
						"original_id": 123456,
						"defindex": 5470,
						"level": 1,
						"quality": 4,
						"inventory": 19,
						"quantity": 1,
						"equipped": [
							{
								"class": 23,
								"slot": 6
							}
						]				
					},
					{
						"id": 1234567897,
						"original_id": 1234567897,
						"defindex": 5508,
						"level": 1,
						"quality": 4,
						"inventory": 25,
						"quantity": 1,
						"attributes": [
				            {
				              "defindex": 8,
				              "value": 1049511890,
				              "float_value": 0.27789169549942017
				            }
			            ]
					}
				]
				
			}
		}
	`
}
