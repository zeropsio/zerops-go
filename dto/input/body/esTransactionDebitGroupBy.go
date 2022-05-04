// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/enum"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*EsTransactionDebitGroupBy)(nil)

type EsTransactionDebitGroupBy struct {
	Search           EsTransactionDebitGroupBySearch    `json:"search"`
	From             types.DateTime                     `json:"from"`
	Till             types.DateTime                     `json:"till"`
	TimeZone         types.String                       `json:"timeZone"`
	SubscriptionName types.StringNull                   `json:"subscriptionName"`
	ReceiverId       types.StringNull                   `json:"receiverId"`
	GroupBy          enum.EsTransactionDebitGroupByEnum `json:"groupBy"`
	TimeGroupBy      types.String                       `json:"timeGroupBy"`
}

func (dto EsTransactionDebitGroupBy) GetSearch() EsTransactionDebitGroupBySearch {
	return dto.Search
}
func (dto EsTransactionDebitGroupBy) GetFrom() types.DateTime {
	return dto.From
}
func (dto EsTransactionDebitGroupBy) GetTill() types.DateTime {
	return dto.Till
}
func (dto EsTransactionDebitGroupBy) GetTimeZone() types.String {
	return dto.TimeZone
}
func (dto EsTransactionDebitGroupBy) GetSubscriptionName() types.StringNull {
	return dto.SubscriptionName
}
func (dto EsTransactionDebitGroupBy) GetReceiverId() types.StringNull {
	return dto.ReceiverId
}
func (dto EsTransactionDebitGroupBy) GetGroupBy() enum.EsTransactionDebitGroupByEnum {
	return dto.GroupBy
}
func (dto EsTransactionDebitGroupBy) GetTimeGroupBy() types.String {
	return dto.TimeGroupBy
}

type EsTransactionDebitGroupBySearch []EsSearchItem

func (dto EsTransactionDebitGroupBySearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchItem(dto))
}

func (dto *EsTransactionDebitGroupBy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search           *EsTransactionDebitGroupBySearch
		From             *types.DateTime
		Till             *types.DateTime
		TimeZone         *types.String
		SubscriptionName types.StringNull
		ReceiverId       types.StringNull
		GroupBy          *enum.EsTransactionDebitGroupByEnum
		TimeGroupBy      *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsTransactionDebitGroupBy", err)
	}
	var errorList validator.ErrorList
	if aux.Search == nil {
		errorList = errorList.With(validator.NewError("search", "field is required"))
	}
	if aux.From == nil {
		errorList = errorList.With(validator.NewError("from", "field is required"))
	}
	if aux.Till == nil {
		errorList = errorList.With(validator.NewError("till", "field is required"))
	}
	if aux.TimeZone == nil {
		errorList = errorList.With(validator.NewError("timeZone", "field is required"))
	}
	if aux.GroupBy == nil {
		errorList = errorList.With(validator.NewError("groupBy", "field is required"))
	}
	if aux.TimeGroupBy == nil {
		errorList = errorList.With(validator.NewError("timeGroupBy", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.Search = *aux.Search
	dto.From = *aux.From
	dto.Till = *aux.Till
	dto.TimeZone = *aux.TimeZone
	dto.SubscriptionName = aux.SubscriptionName
	dto.ReceiverId = aux.ReceiverId
	dto.GroupBy = *aux.GroupBy
	dto.TimeGroupBy = *aux.TimeGroupBy

	return nil
}
