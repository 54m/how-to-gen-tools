package werror_test

import (
	"fmt"
	"testing"

	"github.com/54m/how-to-gen-tools/back/domain/werror"
	"github.com/54m/how-to-gen-tools/back/pkg/errorkey"
	"github.com/go-playground/validator/v10"
	"github.com/google/go-cmp/cmp"
)

func TestNewValidationError(t *testing.T) {
	type args struct {
		er error
	}

	validate := validator.New()
	tests := []struct {
		name string
		args args
		want []werror.FailedReason
	}{
		{
			name: "single field(error)",
			args: args{
				er: validate.Struct(&struct {
					Description string `validate:"required"`
				}{}),
			},
			want: []werror.FailedReason{
				{
					Message: "Description: required",
				},
			},
		},
		{
			name: "single field(not error)",
			args: args{
				er: validate.Struct(&struct {
					Count int `validate:"gte=1,lt=10"`
				}{
					Count: 5,
				}),
			},
			want: nil,
		},
		{
			name: "multiple field(error)",
			args: args{
				er: validate.Struct(&struct {
					Name string `validate:"required"`
					Age  int    `validate:"gte=1,lt=10"`
				}{}),
			},
			want: []werror.FailedReason{
				{
					Message: "Age: gte",
				},
				{
					Message: "Name: required",
				},
			},
		},
		{
			name: "multiple field(not error)",
			args: args{
				er: validate.Struct(&struct {
					Name string `validate:"required"`
					Age  int    `validate:"gte=1,lt=10"`
				}{
					Name: "name",
					Age:  5,
				}),
			},
			want: nil,
		},
		{
			name: "guard test(different error)",
			args: args{
				er: fmt.Errorf("error"),
			},
			want: []werror.FailedReason{
				{
					Message: errorkey.Internal,
				},
			},
		},
		{
			name: "guard test(nil)",
			args: args{
				er: nil,
			},
			want: []werror.FailedReason{
				{
					Message: errorkey.Internal,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt // ESCAPE: using the variable on range scope `tt` in function literal
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.er == nil {
				return
			}
			got := werror.NewValidationError(tt.args.er)
			if diff := cmp.Diff(tt.want, got.FailedReasons); diff != "" {
				t.Errorf("service returned wrong response mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
