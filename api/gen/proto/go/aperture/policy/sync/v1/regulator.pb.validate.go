// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aperture/policy/sync/v1/regulator.proto

package syncv1

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

// Validate checks the field values on RegulatorWrapper with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegulatorWrapper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegulatorWrapper with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegulatorWrapperMultiError, or nil if none found.
func (m *RegulatorWrapper) ValidateAll() error {
	return m.validate(true)
}

func (m *RegulatorWrapper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonAttributes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegulatorWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegulatorWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegulatorWrapperValidationError{
				field:  "CommonAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRegulator()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegulatorWrapperValidationError{
					field:  "Regulator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegulatorWrapperValidationError{
					field:  "Regulator",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegulator()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegulatorWrapperValidationError{
				field:  "Regulator",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegulatorWrapperMultiError(errors)
	}

	return nil
}

// RegulatorWrapperMultiError is an error wrapping multiple validation errors
// returned by RegulatorWrapper.ValidateAll() if the designated constraints
// aren't met.
type RegulatorWrapperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegulatorWrapperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegulatorWrapperMultiError) AllErrors() []error { return m }

// RegulatorWrapperValidationError is the validation error returned by
// RegulatorWrapper.Validate if the designated constraints aren't met.
type RegulatorWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegulatorWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegulatorWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegulatorWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegulatorWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegulatorWrapperValidationError) ErrorName() string { return "RegulatorWrapperValidationError" }

// Error satisfies the builtin error interface
func (e RegulatorWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegulatorWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegulatorWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegulatorWrapperValidationError{}

// Validate checks the field values on RegulatorDynamicConfigWrapper with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegulatorDynamicConfigWrapper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegulatorDynamicConfigWrapper with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// RegulatorDynamicConfigWrapperMultiError, or nil if none found.
func (m *RegulatorDynamicConfigWrapper) ValidateAll() error {
	return m.validate(true)
}

func (m *RegulatorDynamicConfigWrapper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonAttributes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegulatorDynamicConfigWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegulatorDynamicConfigWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegulatorDynamicConfigWrapperValidationError{
				field:  "CommonAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRegulatorDynamicConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegulatorDynamicConfigWrapperValidationError{
					field:  "RegulatorDynamicConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegulatorDynamicConfigWrapperValidationError{
					field:  "RegulatorDynamicConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegulatorDynamicConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegulatorDynamicConfigWrapperValidationError{
				field:  "RegulatorDynamicConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegulatorDynamicConfigWrapperMultiError(errors)
	}

	return nil
}

// RegulatorDynamicConfigWrapperMultiError is an error wrapping multiple
// validation errors returned by RegulatorDynamicConfigWrapper.ValidateAll()
// if the designated constraints aren't met.
type RegulatorDynamicConfigWrapperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegulatorDynamicConfigWrapperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegulatorDynamicConfigWrapperMultiError) AllErrors() []error { return m }

// RegulatorDynamicConfigWrapperValidationError is the validation error
// returned by RegulatorDynamicConfigWrapper.Validate if the designated
// constraints aren't met.
type RegulatorDynamicConfigWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegulatorDynamicConfigWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegulatorDynamicConfigWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegulatorDynamicConfigWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegulatorDynamicConfigWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegulatorDynamicConfigWrapperValidationError) ErrorName() string {
	return "RegulatorDynamicConfigWrapperValidationError"
}

// Error satisfies the builtin error interface
func (e RegulatorDynamicConfigWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegulatorDynamicConfigWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegulatorDynamicConfigWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegulatorDynamicConfigWrapperValidationError{}

// Validate checks the field values on RegulatorDecisionWrapper with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegulatorDecisionWrapper) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegulatorDecisionWrapper with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegulatorDecisionWrapperMultiError, or nil if none found.
func (m *RegulatorDecisionWrapper) ValidateAll() error {
	return m.validate(true)
}

func (m *RegulatorDecisionWrapper) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonAttributes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegulatorDecisionWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegulatorDecisionWrapperValidationError{
					field:  "CommonAttributes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegulatorDecisionWrapperValidationError{
				field:  "CommonAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRegulatorDecision()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegulatorDecisionWrapperValidationError{
					field:  "RegulatorDecision",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegulatorDecisionWrapperValidationError{
					field:  "RegulatorDecision",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRegulatorDecision()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegulatorDecisionWrapperValidationError{
				field:  "RegulatorDecision",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegulatorDecisionWrapperMultiError(errors)
	}

	return nil
}

// RegulatorDecisionWrapperMultiError is an error wrapping multiple validation
// errors returned by RegulatorDecisionWrapper.ValidateAll() if the designated
// constraints aren't met.
type RegulatorDecisionWrapperMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegulatorDecisionWrapperMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegulatorDecisionWrapperMultiError) AllErrors() []error { return m }

// RegulatorDecisionWrapperValidationError is the validation error returned by
// RegulatorDecisionWrapper.Validate if the designated constraints aren't met.
type RegulatorDecisionWrapperValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegulatorDecisionWrapperValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegulatorDecisionWrapperValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegulatorDecisionWrapperValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegulatorDecisionWrapperValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegulatorDecisionWrapperValidationError) ErrorName() string {
	return "RegulatorDecisionWrapperValidationError"
}

// Error satisfies the builtin error interface
func (e RegulatorDecisionWrapperValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegulatorDecisionWrapper.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegulatorDecisionWrapperValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegulatorDecisionWrapperValidationError{}

// Validate checks the field values on RegulatorDecision with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegulatorDecision) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegulatorDecision with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegulatorDecisionMultiError, or nil if none found.
func (m *RegulatorDecision) ValidateAll() error {
	return m.validate(true)
}

func (m *RegulatorDecision) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AcceptPercentage

	if len(errors) > 0 {
		return RegulatorDecisionMultiError(errors)
	}

	return nil
}

// RegulatorDecisionMultiError is an error wrapping multiple validation errors
// returned by RegulatorDecision.ValidateAll() if the designated constraints
// aren't met.
type RegulatorDecisionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegulatorDecisionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegulatorDecisionMultiError) AllErrors() []error { return m }

// RegulatorDecisionValidationError is the validation error returned by
// RegulatorDecision.Validate if the designated constraints aren't met.
type RegulatorDecisionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegulatorDecisionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegulatorDecisionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegulatorDecisionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegulatorDecisionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegulatorDecisionValidationError) ErrorName() string {
	return "RegulatorDecisionValidationError"
}

// Error satisfies the builtin error interface
func (e RegulatorDecisionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegulatorDecision.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegulatorDecisionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegulatorDecisionValidationError{}