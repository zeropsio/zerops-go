// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type Settings struct {
	CurrencyList               SettingsCurrencyList      `json:"currencyList"`
	LanguageList               SettingsLanguageList      `json:"languageList"`
	ServiceStackList           SettingsServiceStackList  `json:"serviceStackList"`
	LocationList               SettingsLocationList      `json:"locationList"`
	ObjectStorageList          SettingsObjectStorageList `json:"objectStorageList"`
	ObjectStorageDiskGBytesMax types.Int                 `json:"objectStorageDiskGBytesMax"`
}

type SettingsCurrencyList []Currency

func (dto SettingsCurrencyList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Currency(dto))
}

type SettingsLanguageList []Language

func (dto SettingsLanguageList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Language(dto))
}

type SettingsServiceStackList []ServiceStackType

func (dto SettingsServiceStackList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ServiceStackType(dto))
}

type SettingsLocationList []Location

func (dto SettingsLocationList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]Location(dto))
}

type SettingsObjectStorageList []ObjectStorage

func (dto SettingsObjectStorageList) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ObjectStorage(dto))
}
