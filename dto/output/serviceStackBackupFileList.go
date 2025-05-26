// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ServiceStackBackupFileList struct {
	Files                  ServiceStackBackupFileListFiles `json:"files"`
	BackupPeriod           types.String                    `json:"backupPeriod"`
	RetentionPolicy        *BackupRetentionPolicy          `json:"retentionPolicy"`
	DefaultRetentionPolicy BackupRetentionPolicy           `json:"defaultRetentionPolicy"`
}

type ServiceStackBackupFileListFiles []ServiceStackBackupFile

func (dto ServiceStackBackupFileListFiles) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackBackupFile(dto))
}
