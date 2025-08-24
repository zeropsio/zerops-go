// Generated ZEROPS sdk

package output

import (
	"strconv"

	"github.com/zeropsio/zerops-go/types"
)

var _ strconv.NumError

type ZcliVersionAuthor struct {
	Login               types.String `json:"login"`
	Id                  types.Int    `json:"id"`
	Node_id             types.String `json:"node_id"`
	Avatar_url          types.String `json:"avatar_url"`
	Gravatar_id         types.String `json:"gravatar_id"`
	Url                 types.String `json:"url"`
	Html_url            types.String `json:"html_url"`
	Followers_url       types.String `json:"followers_url"`
	Following_url       types.String `json:"following_url"`
	Gists_url           types.String `json:"gists_url"`
	Starred_url         types.String `json:"starred_url"`
	Subscriptions_url   types.String `json:"subscriptions_url"`
	Organizations_url   types.String `json:"organizations_url"`
	Repos_url           types.String `json:"repos_url"`
	Events_url          types.String `json:"events_url"`
	Received_events_url types.String `json:"received_events_url"`
	Type                types.String `json:"type"`
	User_view_type      types.String `json:"user_view_type"`
	Site_admin          types.Bool   `json:"site_admin"`
}
