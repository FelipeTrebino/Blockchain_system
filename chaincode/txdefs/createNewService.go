package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new service on channel
// POST Method
var CreateNewService = tx.Transaction{
	Tag:         "createNewService",
	Label:       "Create New Service",
	Description: "Create a new service on the channel",
	Method:      "POST",
	Callers:     []string{"$org1MSP", "$org2MSP"}, // Only org1 and org2  can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the service",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "provider",
			Label:       "Provider",
			Description: "Provider of the service",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "price",
			Label:       "Price",
			Description: "Price of the service",
			DataType:    "number",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)
		provider, _ := req["provider"].(string)
		price, _ := req["price"].(float64)

		serviceMap := make(map[string]interface{})
		serviceMap["@assetType"] = "service"
		serviceMap["name"] = name
		serviceMap["provider"] = provider
		serviceMap["price"] = price

		serviceAsset, err := assets.NewAsset(serviceMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new service on channel
		_, err = serviceAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		serviceJSON, nerr := json.Marshal(serviceAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Failed to encode asset to JSON format")
		}

		return serviceJSON, nil
	},
}
