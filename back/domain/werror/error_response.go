package werror

// nolint: dupl

import (
	"fmt"

	"golang.org/x/xerrors"
)

// ErrorResponse - エラーレスポンス
type ErrorResponse struct {
	Status        int            `json:"status"`
	FailedReasons []FailedReason `json:"messages"`
	message       string
	err           error
}

var (
	_ error           = (*ErrorResponse)(nil)
	_ xerrors.Wrapper = (*ErrorResponse)(nil)
)

// Error - standard error format
func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("error response: status=%d, messages=%+v, internal=%+v", err.Status, err.message, err.err)
}

// Unwrap - unwrap the error
func (err *ErrorResponse) Unwrap() error {
	return err.err
}
