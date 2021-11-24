package werror

// nolint: dupl

import (
	"net/http"

	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
)

// NewForbiddenError - 権限がなかった時に使用
func NewForbiddenError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusForbidden,
		FailedReasons: []FailedReason{{
			Message: errorkey.Forbidden,
		}},
		err:     err,
		message: "forbidden error",
	}
}
