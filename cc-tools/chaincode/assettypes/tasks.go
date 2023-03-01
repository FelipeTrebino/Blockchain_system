package assettypes

import (
    "fmt"

    "github.com/goledgerdev/cc-tools/assets"
)

var Tasks = assets.AssetTypes{
    Tag:        "tasks",
    Label:      "Tasks",
    Description: "A task represents a specific activity that needs to be completed.",

    Props: []assets.AssetProp{
        {
            Required: true,
            IsKey:    true,
            Tag:      "id",
            Label:    "ID",
            DataType: "string",
            Writers:  []string{"org2MSP", "orgMSP"},
        },
        {
            Required: true,
            Tag:      "finished",
            Label:    "Finished",
            DataType: "boolean",
            Validate: func(finished interface{}) error {
                if finished == nil {
                    return fmt.Errorf("finished must be non-nil")
                }
                return nil
            },
        },
        {
            Tag:      "description",
            Label:    "Description",
            DataType: "text",
        },
        {
            Required: true,
            Tag:      "serviceId",
            Label:    "Service ID",
            DataType: "->services",
        },
    },
}
