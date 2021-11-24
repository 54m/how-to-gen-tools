package werror

// nolint: dupl

import (
	"net/http"

	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
)

// NewUnauthorizedError - 認証されてない時に使用
func NewUnauthorizedError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusUnauthorized,
		FailedReasons: []FailedReason{{
			Message: errorkey.Unauthorized,
		}},
		err:     err,
		message: "unauthorized error",
	}
}
