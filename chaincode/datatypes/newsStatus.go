package datatypes

import (
	"fmt"
	"strconv"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type NewsStatus float64

const (
	notVoted NewsStatus = iota
	approved
	rejected
)

func (b NewsStatus) CheckType() errors.ICCError {
	switch b {
	case notVoted:
		return nil
	case approved:
		return nil
	case rejected:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

var newsStatus = assets.DataType{
	AcceptedFormats: []string{"number"},
	DropDownValues: map[string]interface{}{
		"notVoted": notVoted,
		"approved": approved,
		"rejected": rejected,
	},
	Description: ``,

	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = (float64)(v)
		case NewsStatus:
			dataVal = (float64)(v)
		case string:
			var err error
			dataVal, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer, is %t", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer, is %t", 400)
		}

		retVal := (NewsStatus)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
