package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"

	"fmt"
)

var Project = assets.assettypes{
	Tag: "project",
	Label: "project",
	Description: "A project represents a major action by a 
	civil construction company, such as the construction of a building",


	Props: []assets.AssetProp{
		{
			//Primary key
			Required: true,
			isKey: true,
			Tag: "id",
			Label: "id",
			DataType: "string",
			Writers: []string{`org1MSP`, "orgMSP"}
		}
		{
			Tag: "description",
			Label: "description",
			DataType: "string",
		}
		{
			Required: true,
			Tag: "projectId",
			Label: "The ID of the project",
			DataType: "->project",
			Validate: func(projectId interface{}) error {
				projectIdStr := projectId.(string)
				if projectId == "" {
					return fmt.Errorf("projectId must be non-empty")
				}
				return nil
			},
		}
		{
			Required: true,
			Tag: "startOfWork",
			Label: "Start of work",
			DataType: "datetime",
			// Validate funcion
			Validate: func(startOfWork interface{}) error {
				startOfWorkStr := startOfWork.(string)
				if startOfWorkStr == "" {
					return fmt.Errorf("startOfWork must be non-empty")
				}
				return nil
			},
		}
	}

}