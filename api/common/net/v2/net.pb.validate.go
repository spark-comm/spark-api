// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/net/v2/net.proto

package v2

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

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ResultMultiError, or nil if none found.
func (m *Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Msg

	// no validation rules for Reason

	// no validation rules for ErrMsg

	// no validation rules for Data

	if len(errors) > 0 {
		return ResultMultiError(errors)
	}

	return nil
}

// ResultMultiError is an error wrapping multiple validation errors returned by
// Result.ValidateAll() if the designated constraints aren't met.
type ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultMultiError) AllErrors() []error { return m }

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}

// Validate checks the field values on Pagination with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Pagination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pagination with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PaginationMultiError, or
// nil if none found.
func (m *Pagination) ValidateAll() error {
	return m.validate(true)
}

func (m *Pagination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for Size

	if len(errors) > 0 {
		return PaginationMultiError(errors)
	}

	return nil
}

// PaginationMultiError is an error wrapping multiple validation errors
// returned by Pagination.ValidateAll() if the designated constraints aren't met.
type PaginationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PaginationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PaginationMultiError) AllErrors() []error { return m }

// PaginationValidationError is the validation error returned by
// Pagination.Validate if the designated constraints aren't met.
type PaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaginationValidationError) ErrorName() string { return "PaginationValidationError" }

// Error satisfies the builtin error interface
func (e PaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaginationValidationError{}

// Validate checks the field values on RequestPagination with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RequestPagination) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestPagination with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RequestPaginationMultiError, or nil if none found.
func (m *RequestPagination) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestPagination) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageNumber

	// no validation rules for ShowNumber

	if len(errors) > 0 {
		return RequestPaginationMultiError(errors)
	}

	return nil
}

// RequestPaginationMultiError is an error wrapping multiple validation errors
// returned by RequestPagination.ValidateAll() if the designated constraints
// aren't met.
type RequestPaginationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestPaginationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestPaginationMultiError) AllErrors() []error { return m }

// RequestPaginationValidationError is the validation error returned by
// RequestPagination.Validate if the designated constraints aren't met.
type RequestPaginationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestPaginationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestPaginationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestPaginationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestPaginationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestPaginationValidationError) ErrorName() string {
	return "RequestPaginationValidationError"
}

// Error satisfies the builtin error interface
func (e RequestPaginationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestPagination.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestPaginationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestPaginationValidationError{}

// Validate checks the field values on GetByUserListSdk with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetByUserListSdk) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetByUserListSdk with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetByUserListSdkMultiError, or nil if none found.
func (m *GetByUserListSdk) ValidateAll() error {
	return m.validate(true)
}

func (m *GetByUserListSdk) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserID

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetByUserListSdkValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetByUserListSdkValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetByUserListSdkValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetByUserListSdkMultiError(errors)
	}

	return nil
}

// GetByUserListSdkMultiError is an error wrapping multiple validation errors
// returned by GetByUserListSdk.ValidateAll() if the designated constraints
// aren't met.
type GetByUserListSdkMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetByUserListSdkMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetByUserListSdkMultiError) AllErrors() []error { return m }

// GetByUserListSdkValidationError is the validation error returned by
// GetByUserListSdk.Validate if the designated constraints aren't met.
type GetByUserListSdkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByUserListSdkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByUserListSdkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByUserListSdkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByUserListSdkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByUserListSdkValidationError) ErrorName() string { return "GetByUserListSdkValidationError" }

// Error satisfies the builtin error interface
func (e GetByUserListSdkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByUserListSdk.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByUserListSdkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByUserListSdkValidationError{}

// Validate checks the field values on GetByFormUserListSdk with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetByFormUserListSdk) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetByFormUserListSdk with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetByFormUserListSdkMultiError, or nil if none found.
func (m *GetByFormUserListSdk) ValidateAll() error {
	return m.validate(true)
}

func (m *GetByFormUserListSdk) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FromUserID

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetByFormUserListSdkValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetByFormUserListSdkValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetByFormUserListSdkValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetByFormUserListSdkMultiError(errors)
	}

	return nil
}

// GetByFormUserListSdkMultiError is an error wrapping multiple validation
// errors returned by GetByFormUserListSdk.ValidateAll() if the designated
// constraints aren't met.
type GetByFormUserListSdkMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetByFormUserListSdkMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetByFormUserListSdkMultiError) AllErrors() []error { return m }

// GetByFormUserListSdkValidationError is the validation error returned by
// GetByFormUserListSdk.Validate if the designated constraints aren't met.
type GetByFormUserListSdkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetByFormUserListSdkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetByFormUserListSdkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetByFormUserListSdkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetByFormUserListSdkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetByFormUserListSdkValidationError) ErrorName() string {
	return "GetByFormUserListSdkValidationError"
}

// Error satisfies the builtin error interface
func (e GetByFormUserListSdkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetByFormUserListSdk.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetByFormUserListSdkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetByFormUserListSdkValidationError{}
