package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"

	"fmt"
)

var Supervisor = assets.AssetType{
	Tag:         "supervisor",
	Label:       "Supervisor",
	Description: "A supervisor represents a person who supervises a construction site.",
	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "cpf",
			Label:    "CPF",
			DataType: "cpf",
			Writers:  []string{"org2MSP"},
		},
		{
			Required: true,
			Tag:      "name",
			Label:    "Name",
			DataType: "string",
			Validate: func(name interface{}) error {
				nameStr, _ := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		},
		{
			Tag:      "contact",
			Label:    "Contact (xx) xxxxx-xxxx",
			DataType: "string",
			Validate: func(contact interface{}) error {
				contactStr, _ := contact.(string)
				if contactStr == "" {
					return fmt.Errorf("Contact must be non-empty")
				}
				return nil
			},
		},
	},
}
