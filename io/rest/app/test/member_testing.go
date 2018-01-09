// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "ocg": member TestHelpers
//
// Command:
// $ goagen
// --design=github.com/aweisser/ocg-poc/io/rest/design
// --out=$(GOPATH)\src\github.com\aweisser\ocg-poc\io\rest\
// --version=v1.1.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aweisser/ocg-poc/io/rest/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// ShowMemberBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowMemberBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.MemberController, memberID string, ifModifiedSince *string, ifNoneMatch *string) (http.ResponseWriter, *app.JSONAPIErrors) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/members/%v", memberID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	if ifModifiedSince != nil {
		sliceVal := []string{*ifModifiedSince}
		req.Header["If-Modified-Since"] = sliceVal
	}
	if ifNoneMatch != nil {
		sliceVal := []string{*ifNoneMatch}
		req.Header["If-None-Match"] = sliceVal
	}
	prms := url.Values{}
	prms["memberID"] = []string{fmt.Sprintf("%v", memberID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "MemberTest"), rw, req, prms)
	showCtx, _err := app.NewShowMemberContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt *app.JSONAPIErrors
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.JSONAPIErrors)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.JSONAPIErrors", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowMemberInternalServerError runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowMemberInternalServerError(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.MemberController, memberID string, ifModifiedSince *string, ifNoneMatch *string) (http.ResponseWriter, *app.JSONAPIErrors) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/members/%v", memberID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	if ifModifiedSince != nil {
		sliceVal := []string{*ifModifiedSince}
		req.Header["If-Modified-Since"] = sliceVal
	}
	if ifNoneMatch != nil {
		sliceVal := []string{*ifNoneMatch}
		req.Header["If-None-Match"] = sliceVal
	}
	prms := url.Values{}
	prms["memberID"] = []string{fmt.Sprintf("%v", memberID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "MemberTest"), rw, req, prms)
	showCtx, _err := app.NewShowMemberContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 500 {
		t.Errorf("invalid response status code: got %+v, expected 500", rw.Code)
	}
	var mt *app.JSONAPIErrors
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.JSONAPIErrors)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.JSONAPIErrors", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowMemberNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowMemberNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.MemberController, memberID string, ifModifiedSince *string, ifNoneMatch *string) (http.ResponseWriter, *app.JSONAPIErrors) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/members/%v", memberID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	if ifModifiedSince != nil {
		sliceVal := []string{*ifModifiedSince}
		req.Header["If-Modified-Since"] = sliceVal
	}
	if ifNoneMatch != nil {
		sliceVal := []string{*ifNoneMatch}
		req.Header["If-None-Match"] = sliceVal
	}
	prms := url.Values{}
	prms["memberID"] = []string{fmt.Sprintf("%v", memberID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "MemberTest"), rw, req, prms)
	showCtx, _err := app.NewShowMemberContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}
	var mt *app.JSONAPIErrors
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.JSONAPIErrors)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.JSONAPIErrors", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowMemberNotModified runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowMemberNotModified(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.MemberController, memberID string, ifModifiedSince *string, ifNoneMatch *string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/members/%v", memberID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	if ifModifiedSince != nil {
		sliceVal := []string{*ifModifiedSince}
		req.Header["If-Modified-Since"] = sliceVal
	}
	if ifNoneMatch != nil {
		sliceVal := []string{*ifNoneMatch}
		req.Header["If-None-Match"] = sliceVal
	}
	prms := url.Values{}
	prms["memberID"] = []string{fmt.Sprintf("%v", memberID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "MemberTest"), rw, req, prms)
	showCtx, _err := app.NewShowMemberContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 304 {
		t.Errorf("invalid response status code: got %+v, expected 304", rw.Code)
	}

	// Return results
	return rw
}

// ShowMemberOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowMemberOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.MemberController, memberID string, ifModifiedSince *string, ifNoneMatch *string) (http.ResponseWriter, *app.MemberSingle) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/members/%v", memberID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	if ifModifiedSince != nil {
		sliceVal := []string{*ifModifiedSince}
		req.Header["If-Modified-Since"] = sliceVal
	}
	if ifNoneMatch != nil {
		sliceVal := []string{*ifNoneMatch}
		req.Header["If-None-Match"] = sliceVal
	}
	prms := url.Values{}
	prms["memberID"] = []string{fmt.Sprintf("%v", memberID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "MemberTest"), rw, req, prms)
	showCtx, _err := app.NewShowMemberContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.MemberSingle
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.MemberSingle)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.MemberSingle", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}