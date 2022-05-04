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
var _ json.Unmarshaler = (*EsStatsHistoryFilter)(nil)

type EsStatsHistoryFilter struct {
	Search           EsStatsHistoryFilterSearch     `json:"search"`
	Limit            types.IntNull                  `json:"limit"`
	From             types.DateTimeNull             `json:"from"`
	Till             types.DateTimeNull             `json:"till"`
	TimeZone         types.String                   `json:"timeZone"`
	SubscriptionName types.StringNull               `json:"subscriptionName"`
	ReceiverId       types.StringNull               `json:"receiverId"`
	GroupBy          enum.EsStatsHistoryGroupByEnum `json:"groupBy"`
	TimeGroupBy      types.String                   `json:"timeGroupBy"`
}

func (dto EsStatsHistoryFilter) GetSearch() EsStatsHistoryFilterSearch {
	return dto.Search
}
func (dto EsStatsHistoryFilter) GetLimit() types.IntNull {
	return dto.Limit
}
func (dto EsStatsHistoryFilter) GetFrom() types.DateTimeNull {
	return dto.From
}
func (dto EsStatsHistoryFilter) GetTill() types.DateTimeNull {
	return dto.Till
}
func (dto EsStatsHistoryFilter) GetTimeZone() types.String {
	return dto.TimeZone
}
func (dto EsStatsHistoryFilter) GetSubscriptionName() types.StringNull {
	return dto.SubscriptionName
}
func (dto EsStatsHistoryFilter) GetReceiverId() types.StringNull {
	return dto.ReceiverId
}
func (dto EsStatsHistoryFilter) GetGroupBy() enum.EsStatsHistoryGroupByEnum {
	return dto.GroupBy
}
func (dto EsStatsHistoryFilter) GetTimeGroupBy() types.String {
	return dto.TimeGroupBy
}

type EsStatsHistoryFilterSearch []EsSearchNativeItem

func (dto EsStatsHistoryFilterSearch) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]EsSearchNativeItem(dto))
}

func (dto *EsStatsHistoryFilter) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Search           *EsStatsHistoryFilterSearch
		Limit            types.IntNull
		From             types.DateTimeNull
		Till             types.DateTimeNull
		TimeZone         *types.String
		SubscriptionName types.StringNull
		ReceiverId       types.StringNull
		GroupBy          *enum.EsStatsHistoryGroupByEnum
		TimeGroupBy      *types.String
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("EsStatsHistoryFilter", err)
	}
	var errorList validator.ErrorList
	if aux.Search == nil {
		errorList = errorList.With(validator.NewError("search", "field is required"))
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
	dto.Limit = aux.Limit
	dto.From = aux.From
	dto.Till = aux.Till
	dto.TimeZone = *aux.TimeZone
	dto.SubscriptionName = aux.SubscriptionName
	dto.ReceiverId = aux.ReceiverId
	dto.GroupBy = *aux.GroupBy
	dto.TimeGroupBy = *aux.TimeGroupBy

	return nil
}
