// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type GetProjectsByName struct {
	Projects GetProjectsByNameProjects `json:"projects"`
}

type GetProjectsByNameProjects []Project

func (dto GetProjectsByNameProjects) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Project(dto))
}
