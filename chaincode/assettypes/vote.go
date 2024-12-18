package assettypes

import "github.com/hyperledger-labs/cc-tools/assets"

var Vote = assets.AssetType{
	Tag:         "vote",
	Label:       "Vote",
	Description: "Vote on a proposal",

	Props: []assets.AssetProp{
		{
			Required:    true,
			IsKey:       true,
			Tag:         "sub",
			Label:       "Sub",
			Description: "Unique system ID of vote",
			DataType:    "string",
		},
		{
			Required:    true,
			IsKey:       true,
			Tag:         "proposal",
			Label:       "Proposal",
			Description: "Proposal to vote on",
			DataType:    "->proposal",
		},
		{
			Required:    true,
			IsKey:       true,
			Tag:         "user",
			Label:       "User",
			Description: "User that votes",
			DataType:    "->user",
		},
		{
			Required:    true,
			Tag:         "vote",
			Label:       "Vote",
			Description: "Vote on the proposal",
			DataType:    "voteType",
		}
	}
}