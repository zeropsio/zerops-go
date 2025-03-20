// Generated ZEROPS sdk

package uuid

import (
	"encoding/json"

	"github.com/zeropsio/zerops-go/types"
)

type BuildCacheSnapshotId types.UuidShort

func NewBuildCacheSnapshotIdFromString(value string) (out BuildCacheSnapshotId, err error) {
	return BuildCacheSnapshotId(value), nil
}

func (parameter BuildCacheSnapshotId) Native() string {
	return string(parameter)
}

func (parameter BuildCacheSnapshotId) TypedString() types.String {
	return types.String(parameter)
}

type BuildCacheSnapshotIdNull struct {
	value  BuildCacheSnapshotId
	filled bool
}

func (parameter BuildCacheSnapshotId) BuildCacheSnapshotIdNull() BuildCacheSnapshotIdNull {
	return NewBuildCacheSnapshotIdNull(parameter)
}

func NewBuildCacheSnapshotIdNull(value BuildCacheSnapshotId) BuildCacheSnapshotIdNull {
	return BuildCacheSnapshotIdNull{
		filled: true,
		value:  value,
	}
}

func NewBuildCacheSnapshotIdNullFromString(value string) BuildCacheSnapshotIdNull {
	return BuildCacheSnapshotIdNull{
		filled: true,
		value:  BuildCacheSnapshotId(value),
	}
}

func (parameter BuildCacheSnapshotIdNull) Get() (BuildCacheSnapshotId, bool) {
	return parameter.value, parameter.filled
}

func (parameter BuildCacheSnapshotIdNull) Filled() bool {
	return parameter.filled
}

func (parameter BuildCacheSnapshotIdNull) MarshalJSON() ([]byte, error) {
	if parameter.filled {
		return json.Marshal(parameter.value)
	}

	return []byte("null"), nil
}

func (parameter *BuildCacheSnapshotIdNull) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		parameter.filled = false
		return nil
	}

	err := json.Unmarshal(data, &parameter.value)
	parameter.filled = err == nil

	return err
}
