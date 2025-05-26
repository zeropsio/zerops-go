// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*ServiceStackBackupConfig)(nil)

type ServiceStackBackupConfig struct {
	BackupPeriod    types.StringNull       `json:"backupPeriod"`
	RetentionPolicy *BackupRetentionPolicy `json:"retentionPolicy"`
}

func (dto ServiceStackBackupConfig) GetBackupPeriod() types.StringNull {
	return dto.BackupPeriod
}
func (dto ServiceStackBackupConfig) GetRetentionPolicy() *BackupRetentionPolicy {
	return dto.RetentionPolicy
}

func (dto *ServiceStackBackupConfig) UnmarshalJSON(b []byte) error {
	var aux = struct {
		BackupPeriod    types.StringNull
		RetentionPolicy *BackupRetentionPolicy
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("ServiceStackBackupConfig", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.BackupPeriod = aux.BackupPeriod
	dto.RetentionPolicy = aux.RetentionPolicy

	return nil
}
