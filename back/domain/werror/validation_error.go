package werror

// nolint: dupl

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/go-playground/validator/v10"
)

// NewValidationError - バリデーションエラー constructor
func NewValidationError(er error) *ErrorResponse {
	// guard
	if er == nil {
		return NewInternalError(er)
	}

	errs, ok := er.(validator.ValidationErrors)
	// guard
	if !ok {
		return NewInternalError(er)
	}

	reasons := make([]FailedReason, len(errs))
	for i, err := range errs {
		reasons[i].Message = fmt.Sprintf("%s: %s", err.Field(), err.ActualTag())
	}

	sort.Slice(reasons, func(i, j int) bool {
		return reasons[i].Message < reasons[j].Message
	})

	return &ErrorResponse{
		Status:        http.StatusBadRequest,
		FailedReasons: reasons,
		message:       "validation error",
		err:           er,
	}
}
