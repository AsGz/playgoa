package main

import (
	"fmt"
	"playgoa/demo/app"

	"github.com/goadesign/goa"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	r := ctx.Left + ctx.Right
	msg := fmt.Sprintf(`{"result":%d}`, r)
	ctx.OK([]byte(msg))
	return nil
}

// Add runs the add action.
func (c *OperandsController) Des(ctx *app.DesOperandsContext) error {
	r := ctx.Left - ctx.Right
	msg := fmt.Sprintf(`{"result":%d}`, r)
	ctx.OK([]byte(msg))
	return nil
}
