package controller

import (
	"fmt"

	"github.com/takekou0130/meta-curl/adapter/view"
	"github.com/takekou0130/meta-curl/application/inputPort"
)

type Controller struct {
	inputPort *inputPort.InputPort
	view      *view.View
}

func NewController(ip *inputPort.InputPort, v *view.View) *Controller {
	return &Controller{
		inputPort: ip,
		view:      v,
	}
}

func (c *Controller) IndexAction(args []string) {
	if len(args) <= 0 {
		fmt.Errorf("not expected arguments")
	}

	info, err := c.inputPort.MetaInfo(args)
	if err != nil {
		fmt.Errorf("%v", info)
	}
	c.view.Render(info)
}
