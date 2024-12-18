package txdefs

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateProposal = tx.Transaction{
	Tag:         "createProposal",
	Label:       "Create Proposal",
	Description: "Register a new proposal",
	Method:      "POST",

	Args: []tx.Argument{
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

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		news := req["news"].(assets.Key)
		minVotes, _ := req["minVotes"].(float64)
		expiration, _ := req["expiration"].(time.Time)
		sub := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v", news))))

		_, err := news.Get(stub)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "createProposal: error getting news", 500)
		}

		ProposalAssetMap := map[string]interface{}{
			"@assetType": "proposal",
			"sub":        sub,
			"news":       news,
			"minVotes":   minVotes,
			"expiration": expiration,
		}

		ProposalAsset, err := assets.NewAsset(ProposalAssetMap)

		if err != nil {
			return nil, errors.WrapError(err, "createProposal: error creating proposal")
		}

		ProposalData, err := ProposalAsset.Put(stub)
		if err != nil {
			return nil, errors.WrapError(err, "createProposal: error putting proposal")
		}

		return []byte(fmt.Sprintf("Proposal created successfully: %s", ProposalData["sub"].(string))), nil
	},
}
