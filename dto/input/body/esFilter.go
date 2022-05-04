// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsFilter)(nil)

type EsFilter struct {
	Search           EsFilterSearch   `json:"search"`
	Sort             EsFilterSort     `json:"sort"`
	Offset           types.IntNull    `json:"offset"`
	Limit            types.IntNull    `json:"limit"`
	Text             types.StringNull `json:"text"`
	SubscriptionName types.StringNull `json:"subscriptionName"`
	ReceiverId       types.StringNull `json:"receiverId"`
}

func (dto EsFilter) GetSearch() EsFilterSearch {
	return dto.Search
}
func (dto EsFilter) GetSort() EsFilterSort {
	return dto.Sort
}
func (dto EsFilter) GetOffset() types.IntNull {
	return dto.Offset
}
func (dto EsFilter) GetLimit() types.IntNull {
	return dto.Limit
}
func (dto EsFilter) GetText() types.StringNull {
	return dto.Text
}
func (dto EsFilter) GetSubscriptionName() types.StringNull {
	return dto.SubscriptionName
}
func (dto EsFilter) GetReceiverId() types.StringNull {
	return dto.ReceiverId
}

type EsFilterSearch []EsSearchItem

func (dto EsFilterSearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchItem(dto))
}

type EsFilterSort []EsSortItem

func (dto EsFilterSort) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSortItem(dto))
}

func (dto *EsFilter) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search           *EsFilterSearch
		Sort             *EsFilterSort
		Offset           types.IntNull
		Limit            types.IntNull
		Text             types.StringNull
		SubscriptionName types.StringNull
		ReceiverId       types.StringNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsFilter", err)
	}
	var errorList validator.ErrorList
	if aux.Search == nil {
		errorList = errorList.With(validator.NewError("search", "field is required"))
	}
	if aux.Sort == nil {
		errorList = errorList.With(validator.NewError("sort", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Search = *aux.Search
	dto.Sort = *aux.Sort
	dto.Offset = aux.Offset
	dto.Limit = aux.Limit
	dto.Text = aux.Text
	dto.SubscriptionName = aux.SubscriptionName
	dto.ReceiverId = aux.ReceiverId

	return nil
}
