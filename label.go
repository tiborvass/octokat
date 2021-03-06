package octokat

import (
	"fmt"
)

type Label struct {
	URL   string `json:"url,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

func (c *Client) Labels(repo Repo) (labels []*Label, err error) {
	path := fmt.Sprintf("repos/%s/labels", repo)
	err = c.jsonGet(path, &Options{}, &labels)
	return
}

func (c *Client) AppyLabel(repo Repo, issue *Issue, labels []string) error {
	path := fmt.Sprintf("repos/%s/issues/%d/labels", repo, issue.Number)
	out := []*Label{}
	return c.jsonPost(path, &Options{Params: labels}, &out)
}
