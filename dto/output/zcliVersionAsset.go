// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ZcliVersionAsset struct {
	Url                  types.String        `json:"url"`
	Id                   types.Int           `json:"id"`
	Node_id              types.String        `json:"node_id"`
	Name                 types.String        `json:"name"`
	Label                types.String        `json:"label"`
	Uploader             ZcliVersionUploader `json:"uploader"`
	Content_type         types.String        `json:"content_type"`
	State                types.String        `json:"state"`
	Size                 types.Int           `json:"size"`
	Download_count       types.Int           `json:"download_count"`
	Created_at           types.DateTime      `json:"created_at"`
	Updated_at           types.DateTime      `json:"updated_at"`
	Browser_download_url types.String        `json:"browser_download_url"`
}
