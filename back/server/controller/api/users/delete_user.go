// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: 2.3.1
// nolint:dupl
package controller

import (
	types "github.com/54m/how-to-gen-tools/back/interfaces/api/users"
	props "github.com/54m/how-to-gen-tools/back/server/props"
	"github.com/labstack/echo/v4"
)

// DeleteUserController ...
type DeleteUserController interface {
	DeleteUser(c echo.Context, req *types.DeleteUserRequest) (res *types.DeleteUserResponse, err error)
}

type deleteUserController struct {
	*props.ControllerProps
}

// NewDeleteUserController ...
func NewDeleteUserController(cp *props.ControllerProps) DeleteUserController {
	return &deleteUserController{
		ControllerProps: cp,
	}
}

// DeleteUser - DELETE user
func (ctrl *deleteUserController) DeleteUser(
	c echo.Context, req *types.DeleteUserRequest,
) (res *types.DeleteUserResponse, err error) {
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
func (ctrl *deleteUserController) AutoBind() bool {
	return true
}
