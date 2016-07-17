//************************************************************************//
// API "adder": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=playgoa/demo/design
// --out=$(GOPATH)/src/playgoa/demo
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// OperandsController is the controller interface for the Operands actions.
type OperandsController interface {
	goa.Muxer
	Add(*AddOperandsContext) error
}

// MountOperandsController "mounts" a Operands resource controller on the given service.
func MountOperandsController(service *goa.Service, ctrl OperandsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewAddOperandsContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("GET", "/add/:left/:right", ctrl.MuxHandler("Add", h, nil))
	service.LogInfo("mount", "ctrl", "Operands", "action", "Add", "route", "GET /add/:left/:right")
}
