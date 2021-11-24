// Package controller handles requests and returns responses. api_gen users directly edit this file.
// generated version: 2.3.1
// nolint:dupl
package controller

import (
	types "github.com/54m/how-to-gen-tools/back/interfaces/api/users"
	props "github.com/54m/how-to-gen-tools/back/server/props"
	"github.com/labstack/echo/v4"
)

// PostController ...
type PostController interface {
	Post(c echo.Context, req *types.PostRequest) (res *types.PostResponse, err error)
}

type postController struct {
	*props.ControllerProps
}

// NewPostController ...
func NewPostController(cp *props.ControllerProps) PostController {
	return &postController{
		ControllerProps: cp,
	}
}

// Post - POST
func (ctrl *postController) Post(
	c echo.Context, req *types.PostRequest,
) (res *types.PostResponse, err error) {
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
func (ctrl *postController) AutoBind() bool {
	return true
}
