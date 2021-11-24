package werror

// nolint: dupl

import (
	"net/http"

	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
)

// NewInternalError - リポジトリエラー
func NewInternalError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusInternalServerError,
		FailedReasons: []FailedReason{{
			Message: errorkey.Internal,
		}},
		err:     err,
		message: "internal server error",
	}
}
