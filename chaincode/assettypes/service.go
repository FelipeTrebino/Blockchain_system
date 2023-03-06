package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Service = assets.AssetType{
	Tag:         "service",
	Label:       "Service",
	Description: "",

	Props: []assets.AssetProp{
		{
			//Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "id",
			DataType: "string",
			Writers:  []string{`org1MSP`, "orgMSP"},
		},
		{
			Tag:      "description",
			Label:    "description",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "projectId",
			Label:    "The ID of the project",
			DataType: "->project",
		},
		{
			Required: true,
			Tag:      "startOfWork",
			Label:    "Start of work",
			DataType: "datetime",
		},
	},
}
