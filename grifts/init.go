package grifts

import (
	"github.com/BryanMoslo/go_finanzly/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
