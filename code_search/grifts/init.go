package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/obayanju/felastab/code_search/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
