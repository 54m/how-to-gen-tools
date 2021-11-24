package definitions

import "time"

// Pagination - ページング設定
type Pagination struct {
	PagingKey time.Time `json:"pagingKey,omitempty" query:"pagingKey"` // ページングキー
	SortOrder string    `json:"sortOrder,omitempty" query:"sortOrder"` // ソートキー / default desc
	Limit     int       `json:"limit,omitempty"     query:"limit"`     // 取得数
}
