package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"playgoa/demo/app"
	"testing"
)

// AddOperandsOK test setup
func AddOperandsOK(t *testing.T, ctrl app.OperandsController, left int, right int) {
	AddOperandsOKCtx(t, context.Background(), ctrl, left, right)
}

// AddOperandsOKCtx test setup
func AddOperandsOKCtx(t *testing.T, ctx context.Context, ctrl app.OperandsController, left int, right int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/add/%v/%v", left, right), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["left"] = []string{fmt.Sprintf("%v", left)}
	prms["right"] = []string{fmt.Sprintf("%v", right)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OperandsTest"), rw, req, prms)
	addCtx, err := app.NewAddOperandsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Add(addCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// DesOperandsOK test setup
func DesOperandsOK(t *testing.T, ctrl app.OperandsController, left int, right int) {
	DesOperandsOKCtx(t, context.Background(), ctrl, left, right)
}

// DesOperandsOKCtx test setup
func DesOperandsOKCtx(t *testing.T, ctx context.Context, ctrl app.OperandsController, left int, right int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/des/%v/%v", left, right), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["left"] = []string{fmt.Sprintf("%v", left)}
	prms["right"] = []string{fmt.Sprintf("%v", right)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OperandsTest"), rw, req, prms)
	desCtx, err := app.NewDesOperandsContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Des(desCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}
