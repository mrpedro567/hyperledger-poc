package datatypes

import (
	"fmt"
	"strconv"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

type VoteType float64

const (
	reject VoteType = iota
	accept 
)


func (b VoteType) CheckType() errors.ICCError {
	switch b {
	case reject:
		return nil
	case accept:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

func (b VoteType) Value() float64 {
	return (float64)(b)
}

var voteType = assets.DataType{
	AcceptedFormats: []string{"number"},
	DropDownValues: map[string]interface{}{
		"reject": reject,
		"accept": accept,
	},
	Description: ``,

	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataVal float64
		switch v := data.(type) {
		case float64:
			dataVal = v
		case int:
			dataVal = (float64)(v)
		case VoteType:
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

		retVal := (VoteType)(dataVal)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}
