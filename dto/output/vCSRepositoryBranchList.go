// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type VCSRepositoryBranchList struct {
	Branches VCSRepositoryBranchListBranches `json:"branches"`
}

type VCSRepositoryBranchListBranches []VCSRepositoryBranch

func (dto VCSRepositoryBranchListBranches) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]VCSRepositoryBranch(dto))
}
