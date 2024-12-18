package txdefs

import (
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateNews = tx.Transaction{
	Tag:         "createNews",
	Label:       "Create News",
	Description: "Register a new news",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required:    true,
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
			Tag:         "status",
			Label:       "status",
			Description: "status of news",
			DataType:    "newsStatus",
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		sub, _ := req["sub"].(string)
		summary, _ := req["summary"].(string)
		content, _ := req["content"].(string)
		status, _ := req["status"].(int)

		newsAssetMap := map[string]interface{}{
			"@assetType": "news",
			"sub":        sub,
			"summary":    summary,
			"content":    content,
			"status":     status,
		}

		NewsAsset, err := assets.NewAsset(newsAssetMap)

		if err != nil {
			return nil, errors.WrapError(err, "failed to create news asset")
		}

		NewsData, err := NewsAsset.PutNew(stub)

		if err != nil {
			return nil, errors.WrapError(err, "failed to put news asset")
		}

		response := fmt.Sprintf("News created successfully: %s", NewsData["sub"].(string))

		return []byte(response), nil
	},
}
