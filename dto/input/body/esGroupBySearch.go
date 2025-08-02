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
var _ json.Unmarshaler = (*EsGroupBySearch)(nil)

type EsGroupBySearch struct {
	Search           EsGroupBySearchSearch           `json:"search"`
	From             types.DateTime                  `json:"from"`
	Till             types.DateTime                  `json:"till"`
	TimeZone         types.String                    `json:"timeZone"`
	SubscriptionName types.StringNull                `json:"subscriptionName"`
	ReceiverId       types.StringNull                `json:"receiverId"`
	GroupBy          enum.EsGroupBySearchGroupByEnum `json:"groupBy"`
	TimeGroupBy      types.String                    `json:"timeGroupBy"`
}

func (dto EsGroupBySearch) GetSearch() EsGroupBySearchSearch {
	return dto.Search
}
func (dto EsGroupBySearch) GetFrom() types.DateTime {
	return dto.From
}
func (dto EsGroupBySearch) GetTill() types.DateTime {
	return dto.Till
}
func (dto EsGroupBySearch) GetTimeZone() types.String {
	return dto.TimeZone
}
func (dto EsGroupBySearch) GetSubscriptionName() types.StringNull {
	return dto.SubscriptionName
}
func (dto EsGroupBySearch) GetReceiverId() types.StringNull {
	return dto.ReceiverId
}
func (dto EsGroupBySearch) GetGroupBy() enum.EsGroupBySearchGroupByEnum {
	return dto.GroupBy
}
func (dto EsGroupBySearch) GetTimeGroupBy() types.String {
	return dto.TimeGroupBy
}

type EsGroupBySearchSearch []EsSearchItem

func (dto EsGroupBySearchSearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchItem(dto))
}

func (dto *EsGroupBySearch) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search           *EsGroupBySearchSearch
		From             *types.DateTime
		Till             *types.DateTime
		TimeZone         *types.String
		SubscriptionName types.StringNull
		ReceiverId       types.StringNull
		GroupBy          *enum.EsGroupBySearchGroupByEnum
		TimeGroupBy      *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsGroupBySearch", err)
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
