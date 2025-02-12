// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/model/v2/msg.proto

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

// Validate checks the field values on MsgReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MsgReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MsgReqMultiError, or nil if none found.
func (m *MsgReq) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identifier

	// no validation rules for Token

	if utf8.RuneCountInString(m.GetSendID()) < 1 {
		err := MsgReqValidationError{
			field:  "SendID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTrackID()) < 1 {
		err := MsgReqValidationError{
			field:  "TrackID",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetMsgIncr()) < 1 {
		err := MsgReqValidationError{
			field:  "MsgIncr",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Data

	if len(errors) > 0 {
		return MsgReqMultiError(errors)
	}

	return nil
}

// MsgReqMultiError is an error wrapping multiple validation errors returned by
// MsgReq.ValidateAll() if the designated constraints aren't met.
type MsgReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgReqMultiError) AllErrors() []error { return m }

// MsgReqValidationError is the validation error returned by MsgReq.Validate if
// the designated constraints aren't met.
type MsgReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgReqValidationError) ErrorName() string { return "MsgReqValidationError" }

// Error satisfies the builtin error interface
func (e MsgReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgReqValidationError{}

// Validate checks the field values on MsgReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MsgReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MsgReplyMultiError, or nil
// if none found.
func (m *MsgReply) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identifier

	// no validation rules for TrackID

	// no validation rules for MsgIncr

	// no validation rules for Code

	// no validation rules for Msg

	// no validation rules for Reason

	// no validation rules for Data

	// no validation rules for ServerMsgId

	// no validation rules for Seq

	if len(errors) > 0 {
		return MsgReplyMultiError(errors)
	}

	return nil
}

// MsgReplyMultiError is an error wrapping multiple validation errors returned
// by MsgReply.ValidateAll() if the designated constraints aren't met.
type MsgReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgReplyMultiError) AllErrors() []error { return m }

// MsgReplyValidationError is the validation error returned by
// MsgReply.Validate if the designated constraints aren't met.
type MsgReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgReplyValidationError) ErrorName() string { return "MsgReplyValidationError" }

// Error satisfies the builtin error interface
func (e MsgReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgReplyValidationError{}

// Validate checks the field values on GetMaxSeqMsg with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetMaxSeqMsg) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMaxSeqMsg with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetMaxSeqMsgMultiError, or
// nil if none found.
func (m *GetMaxSeqMsg) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMaxSeqMsg) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SessionType

	// no validation rules for Uid

	if len(errors) > 0 {
		return GetMaxSeqMsgMultiError(errors)
	}

	return nil
}

// GetMaxSeqMsgMultiError is an error wrapping multiple validation errors
// returned by GetMaxSeqMsg.ValidateAll() if the designated constraints aren't met.
type GetMaxSeqMsgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMaxSeqMsgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMaxSeqMsgMultiError) AllErrors() []error { return m }

// GetMaxSeqMsgValidationError is the validation error returned by
// GetMaxSeqMsg.Validate if the designated constraints aren't met.
type GetMaxSeqMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMaxSeqMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMaxSeqMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMaxSeqMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMaxSeqMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMaxSeqMsgValidationError) ErrorName() string { return "GetMaxSeqMsgValidationError" }

// Error satisfies the builtin error interface
func (e GetMaxSeqMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMaxSeqMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMaxSeqMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMaxSeqMsgValidationError{}

// Validate checks the field values on ResMaxSeqMsg with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResMaxSeqMsg) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResMaxSeqMsg with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResMaxSeqMsgMultiError, or
// nil if none found.
func (m *ResMaxSeqMsg) ValidateAll() error {
	return m.validate(true)
}

func (m *ResMaxSeqMsg) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SessionType

	// no validation rules for MaxSeq

	// no validation rules for Uid

	if len(errors) > 0 {
		return ResMaxSeqMsgMultiError(errors)
	}

	return nil
}

// ResMaxSeqMsgMultiError is an error wrapping multiple validation errors
// returned by ResMaxSeqMsg.ValidateAll() if the designated constraints aren't met.
type ResMaxSeqMsgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResMaxSeqMsgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResMaxSeqMsgMultiError) AllErrors() []error { return m }

// ResMaxSeqMsgValidationError is the validation error returned by
// ResMaxSeqMsg.Validate if the designated constraints aren't met.
type ResMaxSeqMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResMaxSeqMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResMaxSeqMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResMaxSeqMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResMaxSeqMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResMaxSeqMsgValidationError) ErrorName() string { return "ResMaxSeqMsgValidationError" }

// Error satisfies the builtin error interface
func (e ResMaxSeqMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResMaxSeqMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResMaxSeqMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResMaxSeqMsgValidationError{}

// Validate checks the field values on PullingMsg with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PullingMsg) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PullingMsg with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PullingMsgMultiError, or
// nil if none found.
func (m *PullingMsg) ValidateAll() error {
	return m.validate(true)
}

func (m *PullingMsg) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SessionType

	// no validation rules for MsgType

	// no validation rules for Uid

	if len(m.GetSeqs()) < 1 {
		err := PullingMsgValidationError{
			field:  "Seqs",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PullingMsgMultiError(errors)
	}

	return nil
}

// PullingMsgMultiError is an error wrapping multiple validation errors
// returned by PullingMsg.ValidateAll() if the designated constraints aren't met.
type PullingMsgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PullingMsgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PullingMsgMultiError) AllErrors() []error { return m }

// PullingMsgValidationError is the validation error returned by
// PullingMsg.Validate if the designated constraints aren't met.
type PullingMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PullingMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PullingMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PullingMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PullingMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PullingMsgValidationError) ErrorName() string { return "PullingMsgValidationError" }

// Error satisfies the builtin error interface
func (e PullingMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPullingMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PullingMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PullingMsgValidationError{}

// Validate checks the field values on ResPullingMsg with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResPullingMsg) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResPullingMsg with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResPullingMsgMultiError, or
// nil if none found.
func (m *ResPullingMsg) ValidateAll() error {
	return m.validate(true)
}

func (m *ResPullingMsg) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SessionType

	// no validation rules for Uid

	// no validation rules for MinSeq

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResPullingMsgValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResPullingMsgValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResPullingMsgValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Count

	if len(errors) > 0 {
		return ResPullingMsgMultiError(errors)
	}

	return nil
}

// ResPullingMsgMultiError is an error wrapping multiple validation errors
// returned by ResPullingMsg.ValidateAll() if the designated constraints
// aren't met.
type ResPullingMsgMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResPullingMsgMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResPullingMsgMultiError) AllErrors() []error { return m }

// ResPullingMsgValidationError is the validation error returned by
// ResPullingMsg.Validate if the designated constraints aren't met.
type ResPullingMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResPullingMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResPullingMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResPullingMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResPullingMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResPullingMsgValidationError) ErrorName() string { return "ResPullingMsgValidationError" }

// Error satisfies the builtin error interface
func (e ResPullingMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResPullingMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResPullingMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResPullingMsgValidationError{}

// Validate checks the field values on OfflinePushInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OfflinePushInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OfflinePushInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OfflinePushInfoMultiError, or nil if none found.
func (m *OfflinePushInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *OfflinePushInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Desc

	// no validation rules for Ex

	// no validation rules for IOSPushSound

	// no validation rules for IOSBadgeCount

	if len(errors) > 0 {
		return OfflinePushInfoMultiError(errors)
	}

	return nil
}

// OfflinePushInfoMultiError is an error wrapping multiple validation errors
// returned by OfflinePushInfo.ValidateAll() if the designated constraints
// aren't met.
type OfflinePushInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OfflinePushInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OfflinePushInfoMultiError) AllErrors() []error { return m }

// OfflinePushInfoValidationError is the validation error returned by
// OfflinePushInfo.Validate if the designated constraints aren't met.
type OfflinePushInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OfflinePushInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OfflinePushInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OfflinePushInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OfflinePushInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OfflinePushInfoValidationError) ErrorName() string { return "OfflinePushInfoValidationError" }

// Error satisfies the builtin error interface
func (e OfflinePushInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOfflinePushInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OfflinePushInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OfflinePushInfoValidationError{}

// Validate checks the field values on MsgData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MsgData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MsgData with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MsgDataMultiError, or nil if none found.
func (m *MsgData) ValidateAll() error {
	return m.validate(true)
}

func (m *MsgData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SendID

	// no validation rules for ReceiveID

	// no validation rules for GroupID

	// no validation rules for ClientMsgID

	// no validation rules for ServerMsgID

	// no validation rules for SenderPlatform

	// no validation rules for SenderNickname

	// no validation rules for SenderFaceURL

	// no validation rules for SessionType

	// no validation rules for MsgFrom

	// no validation rules for ContentType

	if len(m.GetContent()) < 1 {
		err := MsgDataValidationError{
			field:  "Content",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Seq

	// no validation rules for SendTime

	// no validation rules for CreateTime

	// no validation rules for Status

	// no validation rules for Options

	if all {
		switch v := interface{}(m.GetOfflinePushInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MsgDataValidationError{
					field:  "OfflinePushInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MsgDataValidationError{
					field:  "OfflinePushInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOfflinePushInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MsgDataValidationError{
				field:  "OfflinePushInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for MsgDataList

	// no validation rules for AttachedInfo

	// no validation rules for Ex

	// no validation rules for IsReact

	// no validation rules for IsExternalExtensions

	// no validation rules for MsgFirstModifyTime

	if len(errors) > 0 {
		return MsgDataMultiError(errors)
	}

	return nil
}

// MsgDataMultiError is an error wrapping multiple validation errors returned
// by MsgData.ValidateAll() if the designated constraints aren't met.
type MsgDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MsgDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MsgDataMultiError) AllErrors() []error { return m }

// MsgDataValidationError is the validation error returned by MsgData.Validate
// if the designated constraints aren't met.
type MsgDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MsgDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MsgDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MsgDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MsgDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MsgDataValidationError) ErrorName() string { return "MsgDataValidationError" }

// Error satisfies the builtin error interface
func (e MsgDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMsgData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MsgDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MsgDataValidationError{}

// Validate checks the field values on MessageTransfer with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MessageTransfer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageTransfer with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageTransferMultiError, or nil if none found.
func (m *MessageTransfer) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageTransfer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for TenantId

	if all {
		switch v := interface{}(m.GetMsgDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MessageTransferValidationError{
					field:  "MsgDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MessageTransferValidationError{
					field:  "MsgDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMsgDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MessageTransferValidationError{
				field:  "MsgDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MessageTransferMultiError(errors)
	}

	return nil
}

// MessageTransferMultiError is an error wrapping multiple validation errors
// returned by MessageTransfer.ValidateAll() if the designated constraints
// aren't met.
type MessageTransferMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageTransferMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageTransferMultiError) AllErrors() []error { return m }

// MessageTransferValidationError is the validation error returned by
// MessageTransfer.Validate if the designated constraints aren't met.
type MessageTransferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageTransferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageTransferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageTransferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageTransferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageTransferValidationError) ErrorName() string { return "MessageTransferValidationError" }

// Error satisfies the builtin error interface
func (e MessageTransferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageTransfer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageTransferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageTransferValidationError{}

// Validate checks the field values on MessageSaveToMq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MessageSaveToMq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageSaveToMq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageSaveToMqMultiError, or nil if none found.
func (m *MessageSaveToMq) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageSaveToMq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MessageSaveToMqValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MessageSaveToMqValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MessageSaveToMqValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MessageSaveToMqMultiError(errors)
	}

	return nil
}

// MessageSaveToMqMultiError is an error wrapping multiple validation errors
// returned by MessageSaveToMq.ValidateAll() if the designated constraints
// aren't met.
type MessageSaveToMqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageSaveToMqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageSaveToMqMultiError) AllErrors() []error { return m }

// MessageSaveToMqValidationError is the validation error returned by
// MessageSaveToMq.Validate if the designated constraints aren't met.
type MessageSaveToMqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageSaveToMqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageSaveToMqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageSaveToMqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageSaveToMqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageSaveToMqValidationError) ErrorName() string { return "MessageSaveToMqValidationError" }

// Error satisfies the builtin error interface
func (e MessageSaveToMqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageSaveToMq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageSaveToMqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageSaveToMqValidationError{}
