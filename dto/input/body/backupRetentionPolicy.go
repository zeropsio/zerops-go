// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*BackupRetentionPolicy)(nil)

type BackupRetentionPolicy struct {
	MaxTotalFiles types.Int         `json:"maxTotalFiles"`
	MaxTotalGiB   types.Int         `json:"maxTotalGiB"`
	ProtectedTags types.StringArray `json:"protectedTags"`
	MinDaily      types.Int         `json:"minDaily"`
	MaxDaily      types.Int         `json:"maxDaily"`
	MinWeekly     types.Int         `json:"minWeekly"`
	MaxWeekly     types.Int         `json:"maxWeekly"`
	MinMonthly    types.Int         `json:"minMonthly"`
	MaxMonthly    types.Int         `json:"maxMonthly"`
}

func (dto BackupRetentionPolicy) GetMaxTotalFiles() types.Int {
	return dto.MaxTotalFiles
}
func (dto BackupRetentionPolicy) GetMaxTotalGiB() types.Int {
	return dto.MaxTotalGiB
}
func (dto BackupRetentionPolicy) GetProtectedTags() types.StringArray {
	return dto.ProtectedTags
}
func (dto BackupRetentionPolicy) GetMinDaily() types.Int {
	return dto.MinDaily
}
func (dto BackupRetentionPolicy) GetMaxDaily() types.Int {
	return dto.MaxDaily
}
func (dto BackupRetentionPolicy) GetMinWeekly() types.Int {
	return dto.MinWeekly
}
func (dto BackupRetentionPolicy) GetMaxWeekly() types.Int {
	return dto.MaxWeekly
}
func (dto BackupRetentionPolicy) GetMinMonthly() types.Int {
	return dto.MinMonthly
}
func (dto BackupRetentionPolicy) GetMaxMonthly() types.Int {
	return dto.MaxMonthly
}

func (dto *BackupRetentionPolicy) UnmarshalJSON(b []byte) error {
	var aux = struct {
		MaxTotalFiles *types.Int
		MaxTotalGiB   *types.Int
		ProtectedTags *types.StringArray
		MinDaily      *types.Int
		MaxDaily      *types.Int
		MinWeekly     *types.Int
		MaxWeekly     *types.Int
		MinMonthly    *types.Int
		MaxMonthly    *types.Int
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("BackupRetentionPolicy", err)
	}
	var errorList validator.ErrorList
	if aux.MaxTotalFiles == nil {
		errorList = errorList.With(validator.NewError("maxTotalFiles", "field is required"))
	}
	if aux.MaxTotalGiB == nil {
		errorList = errorList.With(validator.NewError("maxTotalGiB", "field is required"))
	}
	if aux.ProtectedTags == nil {
		errorList = errorList.With(validator.NewError("protectedTags", "field is required"))
	}
	if aux.MinDaily == nil {
		errorList = errorList.With(validator.NewError("minDaily", "field is required"))
	}
	if aux.MaxDaily == nil {
		errorList = errorList.With(validator.NewError("maxDaily", "field is required"))
	}
	if aux.MinWeekly == nil {
		errorList = errorList.With(validator.NewError("minWeekly", "field is required"))
	}
	if aux.MaxWeekly == nil {
		errorList = errorList.With(validator.NewError("maxWeekly", "field is required"))
	}
	if aux.MinMonthly == nil {
		errorList = errorList.With(validator.NewError("minMonthly", "field is required"))
	}
	if aux.MaxMonthly == nil {
		errorList = errorList.With(validator.NewError("maxMonthly", "field is required"))
	}
	if errorList != nil {
		return errorList.GetError()
	}
	dto.MaxTotalFiles = *aux.MaxTotalFiles
	dto.MaxTotalGiB = *aux.MaxTotalGiB
	dto.ProtectedTags = *aux.ProtectedTags
	dto.MinDaily = *aux.MinDaily
	dto.MaxDaily = *aux.MaxDaily
	dto.MinWeekly = *aux.MinWeekly
	dto.MaxWeekly = *aux.MaxWeekly
	dto.MinMonthly = *aux.MinMonthly
	dto.MaxMonthly = *aux.MaxMonthly

	return nil
}
