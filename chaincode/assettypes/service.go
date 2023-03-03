package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"

	"fmt"
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
			Validate: func(projectId interface{}) error {
				projectIdStr, _ := projectId.(string)
				if projectIdStr == "" {
					return fmt.Errorf("projectId must be non-empty")
				}
				return nil
			},
		},
		{
			Required: true,
			Tag:      "startOfWork",
			Label:    "Start of work",
			DataType: "datetime",
			// Validate funcion
			Validate: func(startOfWork interface{}) error {
				startOfWorkStr, _ := startOfWork.(string)
				if startOfWorkStr == "" {
					return fmt.Errorf("startOfWork must be non-empty")
				}
				return nil
			},
		},
	},
}
