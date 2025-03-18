// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type ServiceStackBackupFileList struct {
	Files ServiceStackBackupFileListFiles `json:"files"`
}

type ServiceStackBackupFileListFiles []ServiceStackBackupFile

func (dto ServiceStackBackupFileListFiles) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackBackupFile(dto))
}
