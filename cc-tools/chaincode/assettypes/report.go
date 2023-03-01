package assettypes

import (
    "github.com/goledgerdev/cc-tools/assets"
)

var Reports = assets.AssetTypes{
    Tag:        "reports",
    Label:      "Reports",
    Description: "A report represents an assessment made by a supervisor on a task.",
    Props: []assets.AssetProp{
        {
            Required: true,
            IsKey:    true,
            Tag:      "supervisorCPF",
            Label:    "Supervisor CPF",
            DataType: "->supervisor",
            Writers:  []string{"org2MSP"},
            Relation: &assets.AssetRelation{
                TargetAsset: "supervisor",
                Unique:      true,
            },
        },
        {
            Required: true,
            IsKey:    true,
            Tag:      "taskId",
            Label:    "Task ID",
            DataType: "->task",
            Writers:  []string{"org2MSP"},
            Relation: &assets.AssetRelation{
                TargetAsset: "tasks",
                Unique:      true,
            },
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
            DataType: "int",
        },
    },
}
