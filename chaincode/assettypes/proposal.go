package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Proposal = assets.AssetType{
	Tag:         "proposal",
	Label:       "proposal",
	Description: "Proposal to vote on whether a news article is fake or not",

	Props: []assets.AssetProp{
		{
			Required:    true,
			IsKey:       true,
			Tag:         "sub",
			Label:       "Sub",
			Description: "Unique system ID of proposal",
			DataType:    "string",
		},
		{
			Required:    true,
			Tag:         "news",
			Label:       "news",
			Description: "news to vote on",
			DataType:    "->news",
		},
		{
			Required:    true,
			Tag:         "minVotes",
			Label:       "minVotes",
			Description: "minimum number of votes required to close the proposal",
			DataType:    "number",
		},
		{
			Required:    true,
			Tag:         "expiration",
			Label:       "expiration",
			Description: "expiration date of the proposal",
			DataType:    "datetime",
		},
	},
}
