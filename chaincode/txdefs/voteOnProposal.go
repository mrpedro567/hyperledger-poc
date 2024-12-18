package txdefs

import (
	"crypto/md5"
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var VoteOnProposal = tx.Transaction{
	Tag:         "voteOnProposal",
	Label:       "Vote on Proposal",
	Description: "Vote on a proposal",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required:    true,
			Tag:         "proposal",
			Label:       "Proposal",
			Description: "Proposal to vote on",
			DataType:    "->proposal",
		},
		{
			Required:    true,
			Tag:         "voter",
			Label:       "Voter",
			Description: "User voting on the proposal",
			DataType:    "->user",
		},
		{
			Required:    true,
			Tag:         "type",
			Label:       "Vote",
			Description: "Vote on the proposal",
			DataType:    "voteType",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		proposal := req["proposal"].(assets.Key)
		voter := req["voter"].(assets.Key)
		vote, _ := req["type"].(int)

		voteAssetMap := map[string]interface{}{
			"@assetType": "vote",
			"sub":        fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v%v", proposal.Key(), voter.Key())))),
			"proposal":   proposal,
			"user":       voter,
			"vote":       vote,
		}

		VoteAsset, err := assets.NewAsset(voteAssetMap)

		if err != nil {
			return nil, errors.WrapError(err, "VoteOnProposal: Error creating vote asset")
		}

		VoteData, err := VoteAsset.Put(stub)
		if err != nil {
			return nil, errors.WrapError(err, "VoteOnProposal: Error putting vote asset")
		}

		return []byte(fmt.Sprintf("%v", VoteData["sub"])), nil

	},
}
