package controller

import (
	"github.com/pkg/errors"
	"github.com/takekou0130/meta-curl/adapter/view"
	"github.com/takekou0130/meta-curl/application/inputPort"
)

type Controller struct {
	inputPort inputPort.InputPort
	view      view.View
}

func NewController(ip *inputPort.InputPort, v *view.View) *Controller {
	return &Controller{
		inputPort: *ip,
		view:      *v,
	}
}

var ErrArguments = errors.New("not expected arguments")

func (c *Controller) IndexAction(args []string) error {
	if len(args) <= 0 {
		return errors.Wrap(ErrArguments, "this controller should have args")
	}

	info, err := c.inputPort.MetaInfo(args)
	if err != nil {
		return errors.Wrap(err, "failed to inputPort.MetaInfo")
	}

	if err = c.view.Render(info); err != nil {
		return errors.Wrap(err, "failed to view.Render")
	}

	return nil
}
