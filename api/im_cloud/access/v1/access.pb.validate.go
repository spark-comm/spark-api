// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: im_cloud/access/v1/access.proto

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

// Validate checks the field values on SendReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SendReqMultiError, or nil if none found.
func (m *SendReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Method

	// no validation rules for Data

	if len(errors) > 0 {
		return SendReqMultiError(errors)
	}

	return nil
}

// SendReqMultiError is an error wrapping multiple validation errors returned
// by SendReq.ValidateAll() if the designated constraints aren't met.
type SendReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendReqMultiError) AllErrors() []error { return m }

// SendReqValidationError is the validation error returned by SendReq.Validate
// if the designated constraints aren't met.
type SendReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendReqValidationError) ErrorName() string { return "SendReqValidationError" }

// Error satisfies the builtin error interface
func (e SendReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendReqValidationError{}

// Validate checks the field values on SendReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendReplyMultiError, or nil
// if none found.
func (m *SendReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SendReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	if len(errors) > 0 {
		return SendReplyMultiError(errors)
	}

	return nil
}

// SendReplyMultiError is an error wrapping multiple validation errors returned
// by SendReply.ValidateAll() if the designated constraints aren't met.
type SendReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendReplyMultiError) AllErrors() []error { return m }

// SendReplyValidationError is the validation error returned by
// SendReply.Validate if the designated constraints aren't met.
type SendReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendReplyValidationError) ErrorName() string { return "SendReplyValidationError" }

// Error satisfies the builtin error interface
func (e SendReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendReplyValidationError{}

// Validate checks the field values on ReceiveReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReceiveReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiveReqMultiError, or
// nil if none found.
func (m *ReceiveReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Method

	// no validation rules for Data

	// no validation rules for Sign

	if len(errors) > 0 {
		return ReceiveReqMultiError(errors)
	}

	return nil
}

// ReceiveReqMultiError is an error wrapping multiple validation errors
// returned by ReceiveReq.ValidateAll() if the designated constraints aren't met.
type ReceiveReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveReqMultiError) AllErrors() []error { return m }

// ReceiveReqValidationError is the validation error returned by
// ReceiveReq.Validate if the designated constraints aren't met.
type ReceiveReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveReqValidationError) ErrorName() string { return "ReceiveReqValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveReqValidationError{}

// Validate checks the field values on ReceiveReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReceiveReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReceiveReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReceiveReplyMultiError, or
// nil if none found.
func (m *ReceiveReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ReceiveReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	// no validation rules for Sign

	if len(errors) > 0 {
		return ReceiveReplyMultiError(errors)
	}

	return nil
}

// ReceiveReplyMultiError is an error wrapping multiple validation errors
// returned by ReceiveReply.ValidateAll() if the designated constraints aren't met.
type ReceiveReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReceiveReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReceiveReplyMultiError) AllErrors() []error { return m }

// ReceiveReplyValidationError is the validation error returned by
// ReceiveReply.Validate if the designated constraints aren't met.
type ReceiveReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiveReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiveReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiveReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiveReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiveReplyValidationError) ErrorName() string { return "ReceiveReplyValidationError" }

// Error satisfies the builtin error interface
func (e ReceiveReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceiveReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiveReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiveReplyValidationError{}

// Validate checks the field values on WebsiteConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WebsiteConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WebsiteConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WebsiteConfigMultiError, or
// nil if none found.
func (m *WebsiteConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *WebsiteConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NodeId

	// no validation rules for PackageName

	// no validation rules for MaxUserNum

	// no validation rules for MaxDayAlive

	// no validation rules for MaxGroupNum

	// no validation rules for PackagePrice

	// no validation rules for PurchaseDuration

	if len(errors) > 0 {
		return WebsiteConfigMultiError(errors)
	}

	return nil
}

// WebsiteConfigMultiError is an error wrapping multiple validation errors
// returned by WebsiteConfig.ValidateAll() if the designated constraints
// aren't met.
type WebsiteConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WebsiteConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WebsiteConfigMultiError) AllErrors() []error { return m }

// WebsiteConfigValidationError is the validation error returned by
// WebsiteConfig.Validate if the designated constraints aren't met.
type WebsiteConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WebsiteConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WebsiteConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WebsiteConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WebsiteConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WebsiteConfigValidationError) ErrorName() string { return "WebsiteConfigValidationError" }

// Error satisfies the builtin error interface
func (e WebsiteConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWebsiteConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WebsiteConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WebsiteConfigValidationError{}

// Validate checks the field values on TenantApplyItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TenantApplyItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TenantApplyItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TenantApplyItemMultiError, or nil if none found.
func (m *TenantApplyItem) ValidateAll() error {
	return m.validate(true)
}

func (m *TenantApplyItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NodeId

	// no validation rules for AssociatedUsersID

	// no validation rules for Phone

	// no validation rules for ApplyType

	// no validation rules for PackageType

	// no validation rules for StartTime

	// no validation rules for EndTime

	// no validation rules for MaxNumber

	// no validation rules for MaxGroupNumber

	// no validation rules for HandleResult

	// no validation rules for HandleReason

	// no validation rules for HandleUserId

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return TenantApplyItemMultiError(errors)
	}

	return nil
}

// TenantApplyItemMultiError is an error wrapping multiple validation errors
// returned by TenantApplyItem.ValidateAll() if the designated constraints
// aren't met.
type TenantApplyItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantApplyItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantApplyItemMultiError) AllErrors() []error { return m }

// TenantApplyItemValidationError is the validation error returned by
// TenantApplyItem.Validate if the designated constraints aren't met.
type TenantApplyItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantApplyItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantApplyItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantApplyItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantApplyItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantApplyItemValidationError) ErrorName() string { return "TenantApplyItemValidationError" }

// Error satisfies the builtin error interface
func (e TenantApplyItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenantApplyItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantApplyItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantApplyItemValidationError{}

// Validate checks the field values on TenantApplyListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TenantApplyListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TenantApplyListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TenantApplyListReplyMultiError, or nil if none found.
func (m *TenantApplyListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *TenantApplyListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TenantApplyListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TenantApplyListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TenantApplyListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TenantApplyListReplyMultiError(errors)
	}

	return nil
}

// TenantApplyListReplyMultiError is an error wrapping multiple validation
// errors returned by TenantApplyListReply.ValidateAll() if the designated
// constraints aren't met.
type TenantApplyListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantApplyListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantApplyListReplyMultiError) AllErrors() []error { return m }

// TenantApplyListReplyValidationError is the validation error returned by
// TenantApplyListReply.Validate if the designated constraints aren't met.
type TenantApplyListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantApplyListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantApplyListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantApplyListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantApplyListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantApplyListReplyValidationError) ErrorName() string {
	return "TenantApplyListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e TenantApplyListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenantApplyListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantApplyListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantApplyListReplyValidationError{}
