package txdefs

import (
    "encoding/json"

    "github.com/goledgerdev/cc-tools/assets"
    "github.com/goledgerdev/cc-tools/errors"
    sw "github.com/goledgerdev/cc-tools/stubwrapper"
    tx "github.com/goledgerdev/cc-tools/transactions"
)

var CreateNewTask = tx.Transaction{
    Tag:         "createNewTask",
    Label:       "Create New Task",
    Description: "Create a new task",
    Method:      "POST",
    Callers:     []string{"$org2MSP", "$orgMSP"}, // Only org2 and org can call this transaction

    Args: []tx.Argument{
        {
            Tag:         "id",
            Label:       "ID",
            Description: "ID of the task",
            DataType:    "string",
            Required:    true,
        },
        {
            Tag:         "finished",
            Label:       "Finished",
            Description: "Whether the task is finished or not",
            DataType:    "boolean",
            Required:    true,
        },
        {
            Tag:         "description",
            Label:       "Description",
            Description: "Description of the task",
            DataType:    "text",
        },
        {
            Tag:         "serviceId",
            Label:       "Service ID",
            Description: "ID of the service this task belongs to",
            DataType:    "->services",
            Required:    true,
        },
    },

    Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
        id, _ := req["id"].(string)
        finished, _ := req["finished"].(bool)
        description, _ := req["description"].(string)
        serviceID, _ := req["serviceId"].(string)

        taskMap := make(map[string]interface{})
        taskMap["@assetType"] = "tasks"
        taskMap["id"] = id
        taskMap["finished"] = finished
        taskMap["description"] = description
        taskMap["serviceId"] = serviceID

        taskAsset, err := assets.NewAsset(taskMap)
        if err != nil {
            return nil, errors.WrapError(err, "Failed to create a new asset")
        }

        // Save the new task on the channel
        _, err = taskAsset.PutNew(stub)
        if err != nil {
            return nil, errors.WrapError(err, "Error saving asset on blockchain")
        }

        // Marshal the task asset back to JSON format
        taskJSON, nerr := json.Marshal(taskAsset)
        if nerr != nil {
            return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
        }

        return taskJSON, nil
    },
}
