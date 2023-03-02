package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Project on channel
// POST Method
var CreateNewProject = tx.Transaction{
	Tag:         "createNewProject",
	Label:       "Create New Project",
	Description: "Create a New Project",
	Method:      "POST",
	Callers:     []string{"$org1MSP", "$orgMSP"}, // Only org1 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the project",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "description",
			Label:       "Description",
			Description: "Description of the project",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "supervisorCPF",
			Label:       "Supervisor CPF",
			Description: "CPF of the project supervisor",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "taskId",
			Label:       "Task ID",
			Description: "ID of the project task",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		description, _ := req["description"].(string)
		supervisorCPF, _ := req["supervisorCPF"].(string)
		taskId, _ := req["taskId"].(string)

		projectMap := make(map[string]interface{})
		projectMap["@assetType"] = "project"
		projectMap["name"] = name
		projectMap["description"] = description
		projectMap["supervisorCPF"] = supervisorCPF
		projectMap["taskId"] = taskId

		projectAsset, err := assets.NewAsset(projectMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Check if caller is org1MSP
		callerMSP, err := stub.GetMSPID()
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get caller MSP ID")
		}
		if callerMSP == "org1MSP" {
			projectAsset.SetPrivateData("collectionOrg3", "excludedOrgs", []string{"org1MSP"})
		}

		// Save the new project on channel
		_, err = projectAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		projectJSON, nerr := json.Marshal(projectAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return projectJSON, nil
	},
}
