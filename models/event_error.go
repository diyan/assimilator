package models

type EventErrorType string

const (
	EventErrorInvalidData       EventErrorType = "invalid_data"
	EventErrorInvalidAttribute  EventErrorType = "invalid_attribute"
	EventErrorValueTooLong      EventErrorType = "value_too_long"
	EventErrorUnknownError      EventErrorType = "unknown_error"
	EventErrorSecurityViolation EventErrorType = "security_violation"
	EventErrorRestrictedIP      EventErrorType = "restricted_ip"

	EventErrorJSGenericFetchError        EventErrorType = "js_generic_fetch_error"
	EventErrorJSInvalidHTTPCode          EventErrorType = "js_invalid_http_code"
	EventErrorJSInvalidContent           EventErrorType = "js_invalid_content"
	EventErrorJSNoColumn                 EventErrorType = "js_no_column"
	EventErrorJSMissingSource            EventErrorType = "js_no_source"
	EventErrorJSInvalidSourceMap         EventErrorType = "js_invalid_source"
	EventErrorJSTooManyRemoteSources     EventErrorType = "js_too_many_sources"
	EventErrorJSInvalidSourceEncoding    EventErrorType = "js_invalid_source_encoding"
	EventErrorJSInvalidSourceMapLocation EventErrorType = "js_invalid_sourcemap_location"
	EventErrorJSTooLarge                 EventErrorType = "js_too_large"
	EventErrorJSFetchTimeout             EventErrorType = "js_fetch_timeout"

	EventErrorNativeNoCrashedThread EventErrorType = "native_no_crashed_thread"
	EventErrorNativeInternalFailure EventErrorType = "native_internal_failure"
	EventErrorNativeNoSymsynd       EventErrorType = "native_no_symsynd"
)

type EventError struct {
	Type  EventErrorType
	Name  string
	Value string
}
