// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/mysqlIssueDemo.proto

package mysql_issue_demo

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

// Validate checks the field values on WriteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WriteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WriteRequestMultiError, or
// nil if none found.
func (m *WriteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return WriteRequestMultiError(errors)
	}

	return nil
}

// WriteRequestMultiError is an error wrapping multiple validation errors
// returned by WriteRequest.ValidateAll() if the designated constraints aren't met.
type WriteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteRequestMultiError) AllErrors() []error { return m }

// WriteRequestValidationError is the validation error returned by
// WriteRequest.Validate if the designated constraints aren't met.
type WriteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteRequestValidationError) ErrorName() string { return "WriteRequestValidationError" }

// Error satisfies the builtin error interface
func (e WriteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteRequestValidationError{}

// Validate checks the field values on WriteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WriteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WriteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WriteResponseMultiError, or
// nil if none found.
func (m *WriteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *WriteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return WriteResponseMultiError(errors)
	}

	return nil
}

// WriteResponseMultiError is an error wrapping multiple validation errors
// returned by WriteResponse.ValidateAll() if the designated constraints
// aren't met.
type WriteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WriteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WriteResponseMultiError) AllErrors() []error { return m }

// WriteResponseValidationError is the validation error returned by
// WriteResponse.Validate if the designated constraints aren't met.
type WriteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WriteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WriteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WriteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WriteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WriteResponseValidationError) ErrorName() string { return "WriteResponseValidationError" }

// Error satisfies the builtin error interface
func (e WriteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWriteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WriteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WriteResponseValidationError{}

// Validate checks the field values on GetAllStoresRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAllStoresRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAllStoresRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAllStoresRequestMultiError, or nil if none found.
func (m *GetAllStoresRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAllStoresRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetAllStoresRequestMultiError(errors)
	}

	return nil
}

// GetAllStoresRequestMultiError is an error wrapping multiple validation
// errors returned by GetAllStoresRequest.ValidateAll() if the designated
// constraints aren't met.
type GetAllStoresRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAllStoresRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAllStoresRequestMultiError) AllErrors() []error { return m }

// GetAllStoresRequestValidationError is the validation error returned by
// GetAllStoresRequest.Validate if the designated constraints aren't met.
type GetAllStoresRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllStoresRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllStoresRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllStoresRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllStoresRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllStoresRequestValidationError) ErrorName() string {
	return "GetAllStoresRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllStoresRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllStoresRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllStoresRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllStoresRequestValidationError{}

// Validate checks the field values on GetAllStoresResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetAllStoresResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetAllStoresResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetAllStoresResponseMultiError, or nil if none found.
func (m *GetAllStoresResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetAllStoresResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetAllStoresResponseMultiError(errors)
	}

	return nil
}

// GetAllStoresResponseMultiError is an error wrapping multiple validation
// errors returned by GetAllStoresResponse.ValidateAll() if the designated
// constraints aren't met.
type GetAllStoresResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetAllStoresResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetAllStoresResponseMultiError) AllErrors() []error { return m }

// GetAllStoresResponseValidationError is the validation error returned by
// GetAllStoresResponse.Validate if the designated constraints aren't met.
type GetAllStoresResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetAllStoresResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetAllStoresResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetAllStoresResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetAllStoresResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetAllStoresResponseValidationError) ErrorName() string {
	return "GetAllStoresResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetAllStoresResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetAllStoresResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetAllStoresResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetAllStoresResponseValidationError{}
