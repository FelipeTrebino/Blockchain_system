package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Report on channel
// POST Method
var CreateNewReport = tx.Transaction{
	Tag:         "createNewReport",
	Label:       "Create New Report",
	Description: "Create a New Report",
	Method:      "POST",
	Callers:     []string{"$org2MSP", "$orgMSP"}, 

	Args: []tx.Argument{
		{
			Tag:         "supervisorCPF",
			Label:       "Supervisor CPF",
			Description: "CPF of the supervisor responsible for the task",
			DataType:    "cpf",
			Required:    true,
		},
		{
			Tag:         "taskId",
			Label:       "Task ID",
			Description: "ID of the task being reported on",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "text",
			Label:       "Report Text",
			Description: "Text describing the report",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "assessment",
			Label:       "Report Assessment",
			Description: "Assessment score of the report",
			DataType:    "int",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		// Get the supervisorCPF and taskId from the request arguments
		supervisorCPF, _ := req["supervisorCPF"].(string)
		taskId, _ := req["taskId"].(string)

		// Check if the supervisor and task assets exist
		supervisorAsset, err := assets.GetAssetByKey(stub, "supervisor", supervisorCPF)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to retrieve supervisor asset")
		}
		taskAsset, err := assets.GetAssetByKey(stub, "tasks", taskId)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to retrieve task asset")
		}

		// Create a new report asset
		reportMap := make(map[string]interface{})
		reportMap["@assetType"] = "reports"
		reportMap["supervisorCPF"] = supervisorAsset.Key
		reportMap["supervisor"] = supervisorAsset
		reportMap["taskId"] = taskAsset.Key
		reportMap["task"] = taskAsset
		reportMap["text"] = req["text"]
		reportMap["assessment"] = req["assessment"]

		reportAsset, err := assets.NewAsset(reportMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new report on channel
		_, err = reportAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		reportJSON, nerr := json.Marshal(reportAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}
		
		return libraryJSON, nil
	},
}

		