package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var User = assets.AssetType{
	Tag:         "user",
	Label:       "user",
	Description: "user of the system that can vote on proposals and create proposals",

	Props: []assets.AssetProp{
		{
			Required:    true,
			IsKey:       true,
			Tag:         "sub",
			Label:       "Sub",
			Description: "Unique system ID of fact-checker",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "document",
			Label:       "Document Number",
			Description: "Legal person's unique identification number",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "name",
			Label:       "Name",
			Description: "User's name",
			DataType:    "string",
		},
		{
			Tag:         "externalId",
			Label:       "External ID",
			Description: "External ID of the user",
			DataType:    "string",
		}
	}
}