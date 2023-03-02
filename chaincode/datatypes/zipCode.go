package datatypes

import (
	"strings"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

var zipCode = assets.DataType{
	AcceptedFormats: []string{"string"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		zipCode, ok := data.(string)
		if !ok {
			return "", nil, errors.NewCCError("property must be a string", 400)
		}

		zipCode = strings.ReplaceAll(zipCode, ".", "")
		zipCode = strings.ReplaceAll(zipCode, "-", "")

		if len(zipCode) != 8 {
			return "", nil, errors.NewCCError("zipCode must have 8 digits", 400)
		}

		stateCode := zipCode[:2]

		if stateCode < "01" || stateCode > "46" {
			return "", nil, errors.NewCCError("Invalid zipCode", 400)
		}

		return zipCode, zipCode, nil
	},
}
