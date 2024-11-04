// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package v1

const (
	OAuthScopes = "oAuth.Scopes"
)

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Description of the error.
	Error string `json:"error"`
}