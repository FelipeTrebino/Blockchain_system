package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
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
		},
		{
			Tag:      "contact",
			Label:    "Contact (xx) xxxxx-xxxx",
			DataType: "string",
		},
	},
}
