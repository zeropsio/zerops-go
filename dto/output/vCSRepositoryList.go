// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type VCSRepositoryList struct {
	Repositories VCSRepositoryListRepositories `json:"repositories"`
}

type VCSRepositoryListRepositories []VCSRepository

func (dto VCSRepositoryListRepositories) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]VCSRepository(dto))
}
