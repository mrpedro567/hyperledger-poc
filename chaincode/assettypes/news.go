package assettypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
)

var News = assets.AssetType{
	Tag:         "news",
	Label:       "news",
	Description: "News or article to vote on whether it is fake or not",

	Props: []assets.AssetProp{
		{
			Required:    true,
			IsKey:       true,
			Tag:         "sub",
			Label:       "Sub",
			Description: "Unique system ID of news",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "summary",
			Label:       "summary",
			Description: "summary of news",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "content",
			Label:       "content",
			Description: "content of news",
			DataType:    "string",
		},
		{
			Tag:          "status",
			Label:        "status",
			Description:  "status of news",
			DataType:     "newsStatus",
			DefaultValue: 0,
		},
	},
}
