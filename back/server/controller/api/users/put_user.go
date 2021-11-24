// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: 2.3.1
// nolint:dupl
package controller

import (
	types "github.com/54m/how-to-gen-tools/back/interfaces/api/users"
	props "github.com/54m/how-to-gen-tools/back/server/props"
	"github.com/labstack/echo/v4"
)

// PutUserController ...
type PutUserController interface {
	PutUser(c echo.Context, req *types.PutUserRequest) (res *types.PutUserResponse, err error)
}

type putUserController struct {
	*props.ControllerProps
}

// NewPutUserController ...
func NewPutUserController(cp *props.ControllerProps) PutUserController {
	return &putUserController{
		ControllerProps: cp,
	}
}

// PutUser - PUT user
func (ctrl *putUserController) PutUser(
	c echo.Context, req *types.PutUserRequest,
) (res *types.PutUserResponse, err error) {
	// return nil, apierror.NewAPIError(http.StatusBadRequest)
	//
	// return nil, apierror.NewAPIError(http.StatusBadRequest).SetError(err)
	//
	// body := map[string]interface{}{
	// 	"code": http.StatusBadRequest,
	// 	"message": "invalid request parameter.",
	// }
	// return nil, apierror.NewAPIError(http.StatusBadRequest, body).SetError(err)
	panic("require implements.") // FIXME require implements.
}

// AutoBind - use echo.Bind
func (ctrl *putUserController) AutoBind() bool {
	return true
}
