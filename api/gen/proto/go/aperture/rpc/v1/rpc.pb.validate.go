// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aperture/rpc/v1/rpc.proto

package rpcv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPayload()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "Payload",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RequestValidationError{
					field:  "Payload",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Response) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Response with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResponseMultiError, or nil
// if none found.
func (m *Response) ValidateAll() error {
	return m.validate(true)
}

func (m *Response) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	switch v := m.Result.(type) {
	case *Response_Payload:
		if v == nil {
			err := ResponseValidationError{
				field:  "Result",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Payload
	case *Response_Error:
		if v == nil {
			err := ResponseValidationError{
				field:  "Result",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetError()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResponseValidationError{
						field:  "Error",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResponseValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ResponseMultiError(errors)
	}

	return nil
}

// ResponseMultiError is an error wrapping multiple validation errors returned
// by Response.ValidateAll() if the designated constraints aren't met.
type ResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResponseMultiError) AllErrors() []error { return m }

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on ServerToClient with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ServerToClient) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ServerToClient with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ServerToClientMultiError,
// or nil if none found.
func (m *ServerToClient) ValidateAll() error {
	return m.validate(true)
}

func (m *ServerToClient) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Command.(type) {
	case *ServerToClient_Request:
		if v == nil {
			err := ServerToClientValidationError{
				field:  "Command",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ServerToClientValidationError{
						field:  "Request",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ServerToClientValidationError{
						field:  "Request",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ServerToClientValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ServerToClient_CancelId:
		if v == nil {
			err := ServerToClientValidationError{
				field:  "Command",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for CancelId
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ServerToClientMultiError(errors)
	}

	return nil
}

// ServerToClientMultiError is an error wrapping multiple validation errors
// returned by ServerToClient.ValidateAll() if the designated constraints
// aren't met.
type ServerToClientMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ServerToClientMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ServerToClientMultiError) AllErrors() []error { return m }

// ServerToClientValidationError is the validation error returned by
// ServerToClient.Validate if the designated constraints aren't met.
type ServerToClientValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerToClientValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerToClientValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerToClientValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerToClientValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerToClientValidationError) ErrorName() string { return "ServerToClientValidationError" }

// Error satisfies the builtin error interface
func (e ServerToClientValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServerToClient.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerToClientValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerToClientValidationError{}

// Validate checks the field values on ClientToServer with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientToServer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientToServer with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientToServerMultiError,
// or nil if none found.
func (m *ClientToServer) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientToServer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Msg.(type) {
	case *ClientToServer_Hello_:
		if v == nil {
			err := ClientToServerValidationError{
				field:  "Msg",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHello()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClientToServerValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClientToServerValidationError{
						field:  "Hello",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHello()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientToServerValidationError{
					field:  "Hello",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ClientToServer_Response:
		if v == nil {
			err := ClientToServerValidationError{
				field:  "Msg",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetResponse()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClientToServerValidationError{
						field:  "Response",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClientToServerValidationError{
						field:  "Response",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientToServerValidationError{
					field:  "Response",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ClientToServerMultiError(errors)
	}

	return nil
}

// ClientToServerMultiError is an error wrapping multiple validation errors
// returned by ClientToServer.ValidateAll() if the designated constraints
// aren't met.
type ClientToServerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientToServerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientToServerMultiError) AllErrors() []error { return m }

// ClientToServerValidationError is the validation error returned by
// ClientToServer.Validate if the designated constraints aren't met.
type ClientToServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientToServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientToServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientToServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientToServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientToServerValidationError) ErrorName() string { return "ClientToServerValidationError" }

// Error satisfies the builtin error interface
func (e ClientToServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientToServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientToServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientToServerValidationError{}

// Validate checks the field values on ClientToServer_Hello with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClientToServer_Hello) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientToServer_Hello with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClientToServer_HelloMultiError, or nil if none found.
func (m *ClientToServer_Hello) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientToServer_Hello) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for NextId

	if len(errors) > 0 {
		return ClientToServer_HelloMultiError(errors)
	}

	return nil
}

// ClientToServer_HelloMultiError is an error wrapping multiple validation
// errors returned by ClientToServer_Hello.ValidateAll() if the designated
// constraints aren't met.
type ClientToServer_HelloMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientToServer_HelloMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientToServer_HelloMultiError) AllErrors() []error { return m }

// ClientToServer_HelloValidationError is the validation error returned by
// ClientToServer_Hello.Validate if the designated constraints aren't met.
type ClientToServer_HelloValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientToServer_HelloValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientToServer_HelloValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientToServer_HelloValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientToServer_HelloValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientToServer_HelloValidationError) ErrorName() string {
	return "ClientToServer_HelloValidationError"
}

// Error satisfies the builtin error interface
func (e ClientToServer_HelloValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientToServer_Hello.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientToServer_HelloValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientToServer_HelloValidationError{}