package main

import (
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"playgoa/demo/client"
)

type (
	// AddOperandsCommand is the command line data structure for the add action of operands
	AddOperandsCommand struct {
		// Left operand
		Left int
		// Right operand
		Right int
	}
	// DesOperandsCommand is the command line data structure for the des action of operands
	DesOperandsCommand struct {
		// Left operand
		Left int
		// Right operand
		Right int
	}
)

// Run makes the HTTP request corresponding to the AddOperandsCommand command.
func (cmd *AddOperandsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/add/%v/%v", cmd.Left, cmd.Right)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AddOperands(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AddOperandsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var left int
	cc.Flags().IntVar(&cmd.Left, "left", left, `Left operand`)
	var right int
	cc.Flags().IntVar(&cmd.Right, "right", right, `Right operand`)
}

// Run makes the HTTP request corresponding to the DesOperandsCommand command.
func (cmd *DesOperandsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/des/%v/%v", cmd.Left, cmd.Right)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DesOperands(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DesOperandsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var left int
	cc.Flags().IntVar(&cmd.Left, "left", left, `Left operand`)
	var right int
	cc.Flags().IntVar(&cmd.Right, "right", right, `Right operand`)
}
