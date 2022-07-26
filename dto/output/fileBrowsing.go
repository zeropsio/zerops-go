// Generated ZEROPS sdk

package output

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/jsonOutputDto.go.tmpl

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

var _ strconv.NumError

type FileBrowsing struct {
	AccessToken        uuid.FileBrowsingAccessToken `json:"accessToken"`
	Expiration         types.DateTime               `json:"expiration"`
	ListUrl            types.String                 `json:"listUrl"`
	ReadFileUrl        types.String                 `json:"readFileUrl"`
	ReadFileContentUrl types.String                 `json:"readFileContentUrl"`
}
