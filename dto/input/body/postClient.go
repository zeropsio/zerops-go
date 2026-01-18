// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PostClient)(nil)

type PostClient struct {
	AccountName types.String `json:"accountName"`
}

func (dto PostClient) GetAccountName() types.String {
	return dto.AccountName
}

func (dto *PostClient) UnmarshalJSON(b []byte) error {
	var aux = struct {
		AccountName *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PostClient", err)
	}
	var errorList validator.ErrorList
	if aux.AccountName == nil {
		errorList = errorList.With(validator.NewError("accountName", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.AccountName = *aux.AccountName

	return nil
}
