package werror

// nolint: dupl

import (
	"net/http"

	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
)

// NewNotFoundError - データが見つからなかった時に使用
func NewNotFoundError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusNotFound,
		FailedReasons: []FailedReason{{
			Message: errorkey.NotFound,
		}},
		err:     err,
		message: "not found error",
	}
}
