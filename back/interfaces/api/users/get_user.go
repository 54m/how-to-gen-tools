package users

import (
	"time"

	"github.com/54m/how-to-gen-tools/back/domain/model"
	"github.com/54m/how-to-gen-tools/back/domain/werror"
	"github.com/54m/how-to-gen-tools/back/pkg/definitions"
)

// @Summary ユーザー情報検索API
// @Description 条件に一致するユーザー情報リストを取得する

// GetRequest - ユーザー情報検索APIリクエスト
type GetRequest struct {
	Name string `json:"name,omitempty" query:"name"` // ニックネーム
	Age  int    `json:"age,omitempty"  query:"age"`  // 年齢
	definitions.Pagination
}

// GetResponse - ユーザー情報検索APIレスポンス
type GetResponse struct {
	Status        int                   `json:"status"`                  // ステータスコード
	Messages      []werror.FailedReason `json:"messages,omitempty"`      // エラーなどのメッセージ
	Payload       []*model.User         `json:"payload"`                 // payload
	NextPagingKey *time.Time            `json:"nextPagingKey,omitempty"` // next paging key
}
