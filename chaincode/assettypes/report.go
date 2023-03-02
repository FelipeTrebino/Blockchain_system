package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Report = assets.AssetType{
	Tag:         "reports",
	Label:       "Reports",
	Description: "A report represents an assessment made by a supervisor on a task.",
	Props: []assets.AssetProp{
		{
			Required: true,
			IsKey:    true,
			Tag:      "supervisorCPF",
			Label:    "Supervisor CPF",
			DataType: "->supervisor",
			Writers:  []string{"org2MSP"},
		},
		{
			Required: true,
			IsKey:    true,
			Tag:      "taskId",
			Label:    "Task ID",
			DataType: "->task",
			Writers:  []string{"org2MSP"},
		},
		{
			Required: true,
			Tag:      "text",
			Label:    "Text",
			DataType: "string",
		},
		{
			Tag:      "assessment",
			Label:    "Assessment",
			DataType: "number",
		},
	},
}
