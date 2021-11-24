package werror

// nolint: dupl

import (
	"net/http"

	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
)

// NewBadRequestError - 不正なリクエストエラー constructor
func NewBadRequestError(err error) *ErrorResponse {
	return &ErrorResponse{
		Status: http.StatusBadRequest,
		FailedReasons: []FailedReason{{
			Message: errorkey.BadRequest,
		}},
		err:     err,
		message: "bad request error",
	}
}
