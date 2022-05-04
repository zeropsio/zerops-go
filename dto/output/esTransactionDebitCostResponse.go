// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type EsTransactionDebitCostResponse struct {
	Client  EsTransactionDebitCostItem            `json:"client"`
	Project EsTransactionDebitCostResponseProject `json:"project"`
	Stack   EsTransactionDebitCostResponseStack   `json:"stack"`
}

type EsTransactionDebitCostResponseProject []EsTransactionDebitCostItem

func (dto EsTransactionDebitCostResponseProject) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsTransactionDebitCostItem(dto))
}

type EsTransactionDebitCostResponseStack []EsTransactionDebitCostItem

func (dto EsTransactionDebitCostResponseStack) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsTransactionDebitCostItem(dto))
}
