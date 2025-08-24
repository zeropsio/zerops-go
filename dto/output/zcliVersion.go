// Generated ZEROPS sdk

package output

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ZcliVersion struct {
	Url              types.String      `json:"url"`
	Assets_url       types.String      `json:"assets_url"`
	Upload_url       types.String      `json:"upload_url"`
	Html_url         types.String      `json:"html_url"`
	Id               types.Int         `json:"id"`
	Author           ZcliVersionAuthor `json:"author"`
	Node_id          types.String      `json:"node_id"`
	Tag_name         types.String      `json:"tag_name"`
	Target_commitish types.String      `json:"target_commitish"`
	Name             types.String      `json:"name"`
	Draft            types.Bool        `json:"draft"`
	Prerelease       types.Bool        `json:"prerelease"`
	Created_at       types.DateTime    `json:"created_at"`
	Published_at     types.DateTime    `json:"published_at"`
	Assets           ZcliVersionAssets `json:"assets"`
	Tarball_url      types.String      `json:"tarball_url"`
	Zipball_url      types.String      `json:"zipball_url"`
	Body             types.String      `json:"body"`
}

type ZcliVersionAssets []ZcliVersionAsset

func (dto ZcliVersionAssets) MarshalJSON() ([]byte, error) {
	if dto == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]ZcliVersionAsset(dto))
}
