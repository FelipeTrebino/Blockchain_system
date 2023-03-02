package txdefs

import (
    "encoding/json"

    "github.com/goledgerdev/cc-tools/assets"
    "github.com/goledgerdev/cc-tools/errors"
    sw "github.com/goledgerdev/cc-tools/stubwrapper"
    tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return all reports from a specific supervisor CPF, ordered by creation date.
// GET method
var GetReportsBySupervisor = tx.Transaction{
    Tag:         "getReportsBySupervisor",
    Label:       "Get Reports by Supervisor CPF",
    Description: "Return all reports that have a specific supervisor CPF, ordered by creation date.",
    Method:      "GET",
    Callers:     []string{"$org2MSP"}, // Only org2 and org3 can call this transaction

    Args: []tx.Argument{
        {
            Tag:         "supervisorCPF",
            Label:       "Supervisor CPF",
            Description: "The CPF of the supervisor",
            DataType:    "cpf",
            Required:    true,
        },
        {
            Tag:         "limit",
            Label:       "Limit",
            Description: "The maximum number of reports to return",
            DataType:    "number",
        },
    },

    Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
        supervisorCPF, _ := req["supervisorCPF"].(string)
        limit, hasLimit := req["limit"].(float64)

        if hasLimit && limit <= 0 {
            return nil, errors.NewCCError("limit must be greater than 0", 400)
        }

        // Prepare couchdb query
        query := map[string]interface{}{
            "selector": map[string]interface{}{
                "@assetType":      "reports",
                "supervisorCPF":   supervisorCPF,
            },
            "sort": []map[string]string{
                {"_created": "asc"},
            },
        }

        if hasLimit {
            query["limit"] = limit
        }

        var err error
        response, err := assets.Search(stub, query, "", true)
        if err != nil {
            return nil, errors.WrapErrorWithStatus(err, "error searching for reports", 500)
        }

        responseJSON, err := json.Marshal(response)
        if err != nil {
            return nil, errors.WrapErrorWithStatus(err, "error marshaling response", 500)
        }

        return responseJSON, nil
    },
}
