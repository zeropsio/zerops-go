// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"encoding/json"
	"strconv"
)

var _ strconv.NumError

type SettingsCountryList struct {
	List SettingsCountryListList `json:"list"`
}

type SettingsCountryListList []Country

func (dto SettingsCountryListList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Country(dto))
}
