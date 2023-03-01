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
			Writers: []string{`org1MSP`}
		}
		{
			Required: true,
			Tag: "name",
			Label: "name",
			DataType: "string",
			// Validate funcion
			Validate: func(name interface{}) error {
				nameStr := name.(string)
				if nameStr == "" {
					return fmt.Errorf("name must be non-empty")
				}
				return nil
			},
		}
		{
			Required: true,
			Tag: "street",
			Label: "street",
			DataType: "string",
			// Validate funcion
			Validate: func(street interface{}) error {
				streetStr := street.(string)
				if streetStr == "" {
					return fmt.Errorf("street must be non-empty")
				}
				return nil
			},
		}
		{
			Required: true,
			Tag: "neighborhood",
			Label: "neighborhood",
			DataType: "string",
			// Validate funcion
			Validate: func(neighborhood interface{}) error {
				neighborhoodStr := neighborhood.(string)
				if neighborhoodStr == "" {
					return fmt.Errorf("neighborhood must be non-empty")
				}
				return nil
			},
		}
		{
			Required: true,
			Tag: "city",
			Label: "city",
			DataType: "string",
			// Validate funcion
			Validate: func(city interface{}) error {
				cityStr := city.(string)
				if cityStr == "" {
					return fmt.Errorf("city must be non-empty")
				}
				return nil
			},
		}
		{
			Required: true,
			Tag: "state",
			Label: "state",
			DataType: "string",
			// Validate funcion
			Validate: func(state interface{}) error {
				stateStr := state.(string)
				if stateStr == "" {
					return fmt.Errorf("state must be non-empty")
				}
				return nil
			},
		}
		{
			Required: true,
			Tag: "number",
			Label: "number",
			DataType: "number",
			// Validate funcion
			Validate: func(number interface{}) error {
				numberStr := number.(string)
				if numberStr == "" {
					return fmt.Errorf("number must be non-empty")
				}
				return nil
			},
		}
		{
			//CEP (zip code)
			Required: true,
			Tag: "zipCode",
			Label: "zipCode (CEP)",
			DataType: "zipCode",
			// Validate funcion
			Validate: func(zipCode interface{}) error {
				zipCodeStr := zipCode.(string)
				if zipCodeStr == "" {
					return fmt.Errorf("zipCode must be non-empty")
				}
				return nil
			},
		}
	}

}