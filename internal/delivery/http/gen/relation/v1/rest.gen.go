//go:build go1.22

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Follow user.
	// (POST /relation/followings)
	PostRelationFollowings(w http.ResponseWriter, r *http.Request)
	// Unfollow user.
	// (DELETE /relation/followings/{userID})
	DeleteRelationFollowingsUserID(w http.ResponseWriter, r *http.Request, userID string)
	// Hide the user.
	// (POST /relation/hidden)
	PostRelationHidden(w http.ResponseWriter, r *http.Request)
	// Unhide the user.
	// (DELETE /relation/hidden/{userID})
	DeleteRelationHiddenUserID(w http.ResponseWriter, r *http.Request, userID string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostRelationFollowings operation middleware
func (siw *ServerInterfaceWrapper) PostRelationFollowings(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, OAuthScopes, []string{"general"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRelationFollowings(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteRelationFollowingsUserID operation middleware
func (siw *ServerInterfaceWrapper) DeleteRelationFollowingsUserID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userID" -------------
	var userID string

	err = runtime.BindStyledParameterWithOptions("simple", "userID", r.PathValue("userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, OAuthScopes, []string{"general"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteRelationFollowingsUserID(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostRelationHidden operation middleware
func (siw *ServerInterfaceWrapper) PostRelationHidden(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, OAuthScopes, []string{"general"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRelationHidden(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteRelationHiddenUserID operation middleware
func (siw *ServerInterfaceWrapper) DeleteRelationHiddenUserID(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "userID" -------------
	var userID string

	err = runtime.BindStyledParameterWithOptions("simple", "userID", r.PathValue("userID"), &userID, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userID", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, OAuthScopes, []string{"general"})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteRelationHiddenUserID(w, r, userID)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{})
}

// ServeMux is an abstraction of http.ServeMux.
type ServeMux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type StdHTTPServerOptions struct {
	BaseURL          string
	BaseRouter       ServeMux
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, m ServeMux) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseRouter: m,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, m ServeMux, baseURL string) http.Handler {
	return HandlerWithOptions(si, StdHTTPServerOptions{
		BaseURL:    baseURL,
		BaseRouter: m,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options StdHTTPServerOptions) http.Handler {
	m := options.BaseRouter

	if m == nil {
		m = http.NewServeMux()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	m.HandleFunc("POST "+options.BaseURL+"/relation/followings", wrapper.PostRelationFollowings)
	m.HandleFunc("DELETE "+options.BaseURL+"/relation/followings/{userID}", wrapper.DeleteRelationFollowingsUserID)
	m.HandleFunc("POST "+options.BaseURL+"/relation/hidden", wrapper.PostRelationHidden)
	m.HandleFunc("DELETE "+options.BaseURL+"/relation/hidden/{userID}", wrapper.DeleteRelationHiddenUserID)

	return m
}

type PostRelationFollowingsRequestObject struct {
	Body *PostRelationFollowingsJSONRequestBody
}

type PostRelationFollowingsResponseObject interface {
	VisitPostRelationFollowingsResponse(w http.ResponseWriter) error
}

type PostRelationFollowings201Response struct {
}

func (response PostRelationFollowings201Response) VisitPostRelationFollowingsResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type PostRelationFollowingsdefaultJSONResponse struct {
	Body       ErrorResponse
	StatusCode int
}

func (response PostRelationFollowingsdefaultJSONResponse) VisitPostRelationFollowingsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type DeleteRelationFollowingsUserIDRequestObject struct {
	UserID string `json:"userID"`
}

type DeleteRelationFollowingsUserIDResponseObject interface {
	VisitDeleteRelationFollowingsUserIDResponse(w http.ResponseWriter) error
}

type DeleteRelationFollowingsUserID200Response struct {
}

func (response DeleteRelationFollowingsUserID200Response) VisitDeleteRelationFollowingsUserIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteRelationFollowingsUserIDdefaultJSONResponse struct {
	Body       ErrorResponse
	StatusCode int
}

func (response DeleteRelationFollowingsUserIDdefaultJSONResponse) VisitDeleteRelationFollowingsUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostRelationHiddenRequestObject struct {
	Body *PostRelationHiddenJSONRequestBody
}

type PostRelationHiddenResponseObject interface {
	VisitPostRelationHiddenResponse(w http.ResponseWriter) error
}

type PostRelationHidden201Response struct {
}

func (response PostRelationHidden201Response) VisitPostRelationHiddenResponse(w http.ResponseWriter) error {
	w.WriteHeader(201)
	return nil
}

type PostRelationHiddendefaultJSONResponse struct {
	Body       ErrorResponse
	StatusCode int
}

func (response PostRelationHiddendefaultJSONResponse) VisitPostRelationHiddenResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type DeleteRelationHiddenUserIDRequestObject struct {
	UserID string `json:"userID"`
}

type DeleteRelationHiddenUserIDResponseObject interface {
	VisitDeleteRelationHiddenUserIDResponse(w http.ResponseWriter) error
}

type DeleteRelationHiddenUserID200Response struct {
}

func (response DeleteRelationHiddenUserID200Response) VisitDeleteRelationHiddenUserIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type DeleteRelationHiddenUserIDdefaultJSONResponse struct {
	Body       ErrorResponse
	StatusCode int
}

func (response DeleteRelationHiddenUserIDdefaultJSONResponse) VisitDeleteRelationHiddenUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Follow user.
	// (POST /relation/followings)
	PostRelationFollowings(ctx context.Context, request PostRelationFollowingsRequestObject) (PostRelationFollowingsResponseObject, error)
	// Unfollow user.
	// (DELETE /relation/followings/{userID})
	DeleteRelationFollowingsUserID(ctx context.Context, request DeleteRelationFollowingsUserIDRequestObject) (DeleteRelationFollowingsUserIDResponseObject, error)
	// Hide the user.
	// (POST /relation/hidden)
	PostRelationHidden(ctx context.Context, request PostRelationHiddenRequestObject) (PostRelationHiddenResponseObject, error)
	// Unhide the user.
	// (DELETE /relation/hidden/{userID})
	DeleteRelationHiddenUserID(ctx context.Context, request DeleteRelationHiddenUserIDRequestObject) (DeleteRelationHiddenUserIDResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostRelationFollowings operation middleware
func (sh *strictHandler) PostRelationFollowings(w http.ResponseWriter, r *http.Request) {
	var request PostRelationFollowingsRequestObject

	var body PostRelationFollowingsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostRelationFollowings(ctx, request.(PostRelationFollowingsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostRelationFollowings")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostRelationFollowingsResponseObject); ok {
		if err := validResponse.VisitPostRelationFollowingsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteRelationFollowingsUserID operation middleware
func (sh *strictHandler) DeleteRelationFollowingsUserID(w http.ResponseWriter, r *http.Request, userID string) {
	var request DeleteRelationFollowingsUserIDRequestObject

	request.UserID = userID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteRelationFollowingsUserID(ctx, request.(DeleteRelationFollowingsUserIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteRelationFollowingsUserID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteRelationFollowingsUserIDResponseObject); ok {
		if err := validResponse.VisitDeleteRelationFollowingsUserIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostRelationHidden operation middleware
func (sh *strictHandler) PostRelationHidden(w http.ResponseWriter, r *http.Request) {
	var request PostRelationHiddenRequestObject

	var body PostRelationHiddenJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostRelationHidden(ctx, request.(PostRelationHiddenRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostRelationHidden")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostRelationHiddenResponseObject); ok {
		if err := validResponse.VisitPostRelationHiddenResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteRelationHiddenUserID operation middleware
func (sh *strictHandler) DeleteRelationHiddenUserID(w http.ResponseWriter, r *http.Request, userID string) {
	var request DeleteRelationHiddenUserIDRequestObject

	request.UserID = userID

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteRelationHiddenUserID(ctx, request.(DeleteRelationHiddenUserIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteRelationHiddenUserID")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteRelationHiddenUserIDResponseObject); ok {
		if err := validResponse.VisitDeleteRelationHiddenUserIDResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
