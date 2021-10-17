package cmd

import (
	"github.com/takekou0130/meta-curl/view"
)

func NewView(mode string) View {
	return View(view.NewTableRenderer())
}
