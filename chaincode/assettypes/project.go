package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Project = assets.AssetType{
	Tag:         "project",
	Label:       "Project",
	Description: "A project represents a major action by a civil construction company, such as the construction of a building",

	Props: []assets.AssetProp{
		{
			//Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "id",
			DataType: "string",
			Writers:  []string{`org1MSP`, `orgMSP`},
		},
		{
			Required: true,
			Tag:      "name",
			Label:    "name",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "street",
			Label:    "street",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "neighborhood",
			Label:    "neighborhood",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "city",
			Label:    "city",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "state",
			Label:    "state",
			DataType: "string",
		},
		{
			Required: true,
			Tag:      "number",
			Label:    "number",
			DataType: "number",
		},
		{
			//CEP (zip code)
			Required: true,
			Tag:      "zipCode",
			Label:    "zipCode (CEP)",
			DataType: "zipCode",
		},
	},
}
