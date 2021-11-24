package werror

// nolint: dupl

import (
	"net/http"

	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
)

// NewConflictError - バージョンが一致しなかった時に使用 constructor of ConflictError
func NewConflictError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusConflict,
		FailedReasons: []FailedReason{{
			Message: errorkey.Conflict,
		}},
		message: "conflict error",
		err:     err,
	}
}
