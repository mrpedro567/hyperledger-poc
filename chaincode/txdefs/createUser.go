package txdefs

import (
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var CreateUser = tx.Transaction{
	Tag:         "createUser",
	Label:       "Create User",
	Description: "Register a new user",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Required:    true,
			Tag:         "sub",
			Label:       "Sub",
			Description: "Unique identifier of the user",
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
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		sub, _ := req["sub"].(string)
		document, _ := req["document"].(string)
		name, _ := req["name"].(string)
		externalId, _ := req["externalId"].(string)

		userAssetMap := map[string]interface{}{
			"@assetType": "user",
			"sub":        sub,
			"document":   document,
			"name":       name,
		}

		if externalId != "" {
			userAssetMap["externalId"] = externalId
		}

		UserAsset, err := assets.NewAsset(userAssetMap)

		if err != nil {
			return nil, errors.WrapError(err, "failed to create user asset")
		}

		UserData, err := UserAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to put user asset")
		}

		response := fmt.Sprintf(
			`{"message": "User created", "user": %s}`,
			UserData["sub"].(string),
		)

		return []byte(response), nil
	},
}
