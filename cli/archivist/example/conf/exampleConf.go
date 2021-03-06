// Code generated by archivist. DO NOT EDIT.

package conf

import (
	"time"

	"github.com/kingsgroupos/archivist/lib/go/archivist"
	"github.com/kingsgroupos/misc/wtime"
	"github.com/pkg/errors"
)

var (
	_ = time.After
	_ = errors.New
	_ = archivist.NewArchivist
	_ = wtime.ParseDuration
)

// easyjson:json
type ExampleConf struct {
	NumberKey         int64  `json:"1number_key"`
	AvatarUrl         string `json:"avatar_url"`
	Blog              string `json:"blog"`
	Company           string `json:"company"`
	CreatedAt         string `json:"created_at"`
	Email             string `json:"email"`
	EventsUrl         string `json:"events_url"`
	Followers         int64  `json:"followers"`
	FollowersUrl      string `json:"followers_url"`
	Following         int64  `json:"following"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	GravatarId        string `json:"gravatar_id"`
	Hireable          bool   `json:"hireable"`
	HtmlUrl           string `json:"html_url"`
	Id                int64  `json:"id"`
	Location          string `json:"location"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	OrganizationsUrl  string `json:"organizations_url"`
	PublicGists       int64  `json:"public_gists"`
	PublicRepos       int64  `json:"public_repos"`
	ReceivedEventsUrl string `json:"received_events_url"`
	ReposUrl          string `json:"repos_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	Type              string `json:"type"`
	UpdatedAt         string `json:"updated_at"`
	Url               string `json:"url"`
}

func (this *ExampleConf) bindRefs(c *Collection) error {
	if this == nil {
		return nil
	}

	var ok bool
	_ = ok

	return nil
}
