// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: access.proto

package access_v1

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

	user_v1 "github.com/8thgencore/microservice-auth/pkg/user/v1"
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

	_ = user_v1.Role(0)
)

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CheckRequestMultiError, or
// nil if none found.
func (m *CheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEndpoint()); l < 1 || l > 255 {
		err := CheckRequestValidationError{
			field:  "Endpoint",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CheckRequest_Endpoint_Pattern.MatchString(m.GetEndpoint()) {
		err := CheckRequestValidationError{
			field:  "Endpoint",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_/-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CheckRequestMultiError(errors)
	}

	return nil
}

// CheckRequestMultiError is an error wrapping multiple validation errors
// returned by CheckRequest.ValidateAll() if the designated constraints aren't met.
type CheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckRequestMultiError) AllErrors() []error { return m }

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}

var _CheckRequest_Endpoint_Pattern = regexp.MustCompile("^[a-zA-Z0-9_/-]+$")

// Validate checks the field values on AddRoleEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddRoleEndpointRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddRoleEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddRoleEndpointRequestMultiError, or nil if none found.
func (m *AddRoleEndpointRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddRoleEndpointRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEndpoint()); l < 1 || l > 255 {
		err := AddRoleEndpointRequestValidationError{
			field:  "Endpoint",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_AddRoleEndpointRequest_Endpoint_Pattern.MatchString(m.GetEndpoint()) {
		err := AddRoleEndpointRequestValidationError{
			field:  "Endpoint",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_/-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetAllowedRoles()) < 1 {
		err := AddRoleEndpointRequestValidationError{
			field:  "AllowedRoles",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddRoleEndpointRequestMultiError(errors)
	}

	return nil
}

// AddRoleEndpointRequestMultiError is an error wrapping multiple validation
// errors returned by AddRoleEndpointRequest.ValidateAll() if the designated
// constraints aren't met.
type AddRoleEndpointRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddRoleEndpointRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddRoleEndpointRequestMultiError) AllErrors() []error { return m }

// AddRoleEndpointRequestValidationError is the validation error returned by
// AddRoleEndpointRequest.Validate if the designated constraints aren't met.
type AddRoleEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddRoleEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddRoleEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddRoleEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddRoleEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddRoleEndpointRequestValidationError) ErrorName() string {
	return "AddRoleEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddRoleEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddRoleEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddRoleEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddRoleEndpointRequestValidationError{}

var _AddRoleEndpointRequest_Endpoint_Pattern = regexp.MustCompile("^[a-zA-Z0-9_/-]+$")

// Validate checks the field values on UpdateRoleEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleEndpointRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleEndpointRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleEndpointRequestMultiError, or nil if none found.
func (m *UpdateRoleEndpointRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleEndpointRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEndpoint()); l < 1 || l > 255 {
		err := UpdateRoleEndpointRequestValidationError{
			field:  "Endpoint",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_UpdateRoleEndpointRequest_Endpoint_Pattern.MatchString(m.GetEndpoint()) {
		err := UpdateRoleEndpointRequestValidationError{
			field:  "Endpoint",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_/-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetAllowedRoles()) < 1 {
		err := UpdateRoleEndpointRequestValidationError{
			field:  "AllowedRoles",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateRoleEndpointRequestMultiError(errors)
	}

	return nil
}

// UpdateRoleEndpointRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateRoleEndpointRequest.ValidateAll() if the
// designated constraints aren't met.
type UpdateRoleEndpointRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleEndpointRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleEndpointRequestMultiError) AllErrors() []error { return m }

// UpdateRoleEndpointRequestValidationError is the validation error returned by
// UpdateRoleEndpointRequest.Validate if the designated constraints aren't met.
type UpdateRoleEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleEndpointRequestValidationError) ErrorName() string {
	return "UpdateRoleEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleEndpointRequestValidationError{}

var _UpdateRoleEndpointRequest_Endpoint_Pattern = regexp.MustCompile("^[a-zA-Z0-9_/-]+$")

// Validate checks the field values on DeleteRoleEndpointRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteRoleEndpointRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRoleEndpointRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteRoleEndpointRequestMultiError, or nil if none found.
func (m *DeleteRoleEndpointRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRoleEndpointRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEndpoint()); l < 1 || l > 255 {
		err := DeleteRoleEndpointRequestValidationError{
			field:  "Endpoint",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_DeleteRoleEndpointRequest_Endpoint_Pattern.MatchString(m.GetEndpoint()) {
		err := DeleteRoleEndpointRequestValidationError{
			field:  "Endpoint",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_/-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteRoleEndpointRequestMultiError(errors)
	}

	return nil
}

// DeleteRoleEndpointRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteRoleEndpointRequest.ValidateAll() if the
// designated constraints aren't met.
type DeleteRoleEndpointRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRoleEndpointRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRoleEndpointRequestMultiError) AllErrors() []error { return m }

// DeleteRoleEndpointRequestValidationError is the validation error returned by
// DeleteRoleEndpointRequest.Validate if the designated constraints aren't met.
type DeleteRoleEndpointRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRoleEndpointRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRoleEndpointRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRoleEndpointRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRoleEndpointRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRoleEndpointRequestValidationError) ErrorName() string {
	return "DeleteRoleEndpointRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteRoleEndpointRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRoleEndpointRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRoleEndpointRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRoleEndpointRequestValidationError{}

var _DeleteRoleEndpointRequest_Endpoint_Pattern = regexp.MustCompile("^[a-zA-Z0-9_/-]+$")

// Validate checks the field values on ListRoleEndpointsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListRoleEndpointsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRoleEndpointsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListRoleEndpointsResponseMultiError, or nil if none found.
func (m *ListRoleEndpointsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRoleEndpointsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEndpointPermissions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListRoleEndpointsResponseValidationError{
						field:  fmt.Sprintf("EndpointPermissions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListRoleEndpointsResponseValidationError{
						field:  fmt.Sprintf("EndpointPermissions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRoleEndpointsResponseValidationError{
					field:  fmt.Sprintf("EndpointPermissions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListRoleEndpointsResponseMultiError(errors)
	}

	return nil
}

// ListRoleEndpointsResponseMultiError is an error wrapping multiple validation
// errors returned by ListRoleEndpointsResponse.ValidateAll() if the
// designated constraints aren't met.
type ListRoleEndpointsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRoleEndpointsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRoleEndpointsResponseMultiError) AllErrors() []error { return m }

// ListRoleEndpointsResponseValidationError is the validation error returned by
// ListRoleEndpointsResponse.Validate if the designated constraints aren't met.
type ListRoleEndpointsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoleEndpointsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoleEndpointsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoleEndpointsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoleEndpointsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoleEndpointsResponseValidationError) ErrorName() string {
	return "ListRoleEndpointsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRoleEndpointsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoleEndpointsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoleEndpointsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoleEndpointsResponseValidationError{}

// Validate checks the field values on EndpointPermissions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EndpointPermissions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EndpointPermissions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EndpointPermissionsMultiError, or nil if none found.
func (m *EndpointPermissions) ValidateAll() error {
	return m.validate(true)
}

func (m *EndpointPermissions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetEndpoint()); l < 1 || l > 255 {
		err := EndpointPermissionsValidationError{
			field:  "Endpoint",
			reason: "value length must be between 1 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_EndpointPermissions_Endpoint_Pattern.MatchString(m.GetEndpoint()) {
		err := EndpointPermissionsValidationError{
			field:  "Endpoint",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_/-]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetAllowedRoles()) < 1 {
		err := EndpointPermissionsValidationError{
			field:  "AllowedRoles",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return EndpointPermissionsMultiError(errors)
	}

	return nil
}

// EndpointPermissionsMultiError is an error wrapping multiple validation
// errors returned by EndpointPermissions.ValidateAll() if the designated
// constraints aren't met.
type EndpointPermissionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointPermissionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointPermissionsMultiError) AllErrors() []error { return m }

// EndpointPermissionsValidationError is the validation error returned by
// EndpointPermissions.Validate if the designated constraints aren't met.
type EndpointPermissionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointPermissionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointPermissionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointPermissionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointPermissionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointPermissionsValidationError) ErrorName() string {
	return "EndpointPermissionsValidationError"
}

// Error satisfies the builtin error interface
func (e EndpointPermissionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointPermissions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointPermissionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointPermissionsValidationError{}

var _EndpointPermissions_Endpoint_Pattern = regexp.MustCompile("^[a-zA-Z0-9_/-]+$")
