package users

import (
	"github.com/54m/how-to-gen-tools/back/domain/model"
	"github.com/54m/how-to-gen-tools/back/domain/werror"
)

// @Summary ユーザー情報取得API
// @Description 指定したユーザーを取得する

// GetUserRequest - ユーザー情報取得APIリクエスト
type GetUserRequest struct {
	UserID string `json:"userID" param:"user_id" validate:"required,len=20"` // ユーザーID
}

// GetUserResponse - ユーザー情報取得APIレスポンス
type GetUserResponse struct {
	Status   int                   `json:"status"`             // ステータスコード
	Messages []werror.FailedReason `json:"messages,omitempty"` // エラーなどのメッセージ
	Payload  *model.User           `json:"payload,omitempty"`  // payload
}

// @Summary ユーザー情報更新API
// @Description 指定したユーザー情報を更新する

// PutUserRequest - ユーザー情報更新APIリクエスト
type PutUserRequest struct {
	UserID  string `json:"userID"  param:"user_id" validate:"required,len=20"` // ユーザーID
	Name    string `json:"name"`                                               // ニックネーム
	Age     int    `json:"age"`                                                // 年齢
	Version int    `json:"version"`                                            // optimistic exclusive lock
}

// PutUserResponse - ユーザー情報更新APIレスポンス
type PutUserResponse struct {
	Status   int                   `json:"status"`             // ステータスコード
	Messages []werror.FailedReason `json:"messages,omitempty"` // エラーなどのメッセージ
	Payload  *model.User           `json:"payload,omitempty"`  // payload
}

// @Summary ユーザー情報削除API
// @Description 指定したユーザーを削除する

// DeleteUserRequest - ユーザー情報削除APIリクエスト
type DeleteUserRequest struct {
	UserID string `json:"userID" param:"user_id" validate:"required,len=20"` // ユーザーID
}

// DeleteUserResponse - ユーザー情報削除APIレスポンス
type DeleteUserResponse struct {
	Status   int                   `json:"status"`             // ステータスコード
	Messages []werror.FailedReason `json:"messages,omitempty"` // エラーなどのメッセージ
}
