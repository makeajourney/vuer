package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/makeajourney/vuer/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
