package types

import (
	"fmt"
)

type (
	ErrorCode uint

	Error struct {
		Code    ErrorCode
		Message string
	}
)

const (
	ErrTopicNotFoundFmt = "Topic %s not found"
	ErrInternalFmt      = "Internal error: %s"
	ErrCriticalFmt      = "Critical error: %s"
	ErrDefaultFmt       = "%s"
)

const (
	ErrTopicNotFound = ErrorCode(1404)
	ErrInternal      = ErrorCode(1500)
	ErrCritical      = ErrorCode(2000)

	ErrShutdown = ErrorCode(0)
	ErrKill     = ErrorCode(1)
)

var (
	codeToErrorMapping = map[ErrorCode]string{
		ErrTopicNotFound: ErrTopicNotFoundFmt,
		ErrInternal:      ErrInternalFmt,
		ErrCritical:      ErrCriticalFmt,
		ErrShutdown:      ErrDefaultFmt,
		ErrKill:          ErrDefaultFmt,
	}
)

func (e *Error) Error() string {
	return e.Message
}

func (c ErrorCode) Format(args ...any) *Error {
	if f, ok := codeToErrorMapping[c]; ok {
		return &Error{
			Code:    c,
			Message: fmt.Sprintf(f, args...),
		}
	}

	return &Error{
		Code:    c,
		Message: fmt.Sprintf(ErrDefaultFmt, args...),
	}
}
