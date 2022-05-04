// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type ProjectImportServiceStack struct {
	Id        uuid.ServiceStackId                `json:"id"`
	Name      types.String                       `json:"name"`
	Error     *ErrorObject                       `json:"error"`
	Processes ProjectImportServiceStackProcesses `json:"processes"`
}

type ProjectImportServiceStackProcesses []Process

func (dto ProjectImportServiceStackProcesses) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Process(dto))
}
