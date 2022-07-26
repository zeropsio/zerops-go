// Generated ZEROPS sdk

package path

// specifier
// template pkg/shared/specifier/generator/publicSdk/golang/dtoGenerator/templates/urlDto.go.tmpl

import (
	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/types/uuid"
)

type S3Bucket struct {
	ServiceStackId uuid.ServiceStackId
	Name           types.String
}
