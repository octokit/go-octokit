package octokat

import (
	"github.com/octokit/octokat/hyper"
)

func (c *Client) Root(headers Headers) (root *hyper.Root, err error) {
	resp := c.Get("", headers)
	if resp.HasError() {
		err = resp.Error
		return
	}

	err = resp.Data(&root)
	return
}
