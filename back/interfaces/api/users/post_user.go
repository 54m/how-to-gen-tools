package users

import (
	"github.com/54m/how-to-gen-tools/back/domain/model"
	"github.com/54m/how-to-gen-tools/back/domain/werror"
)

// @Summary ユーザー情報作成API
// @Description ユーザー情報を登録する

// PostRequest - ユーザー情報作成APIリクエスト
type PostRequest struct {
	Name string `json:"name" validate:"required"` // ニックネーム
	Age  int    `json:"age"`                      // 年齢
}

// PostResponse - ユーザー情報作成APIレスポンス
type PostResponse struct {
	Status   int                   `json:"status"`             // ステータスコード
	Messages []werror.FailedReason `json:"messages,omitempty"` // エラーなどのメッセージ
	Payload  *model.User           `json:"payload,omitempty"`  // payload
}
