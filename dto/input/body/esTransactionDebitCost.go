// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsTransactionDebitCost)(nil)

type EsTransactionDebitCost struct {
	Search EsTransactionDebitCostSearch `json:"search"`
}

func (dto EsTransactionDebitCost) GetSearch() EsTransactionDebitCostSearch {
	return dto.Search
}

type EsTransactionDebitCostSearch []EsSearchItem

func (dto EsTransactionDebitCostSearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchItem(dto))
}

func (dto *EsTransactionDebitCost) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search *EsTransactionDebitCostSearch
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsTransactionDebitCost", err)
	}
	var errorList validator.ErrorList
	if aux.Search == nil {
		errorList = errorList.With(validator.NewError("search", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Search = *aux.Search

	return nil
}
