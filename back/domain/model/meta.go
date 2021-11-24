package model

import (
	"time"
)

// Meta - メタ情報
type Meta struct {
	CreatedAt time.Time  `json:"createdAt"` // 作成日時
	CreatedBy string     `json:"createdBy"` // 作成者
	UpdatedAt time.Time  `json:"updatedAt"` // 更新日時
	UpdatedBy string     `json:"updatedBy"` // 更新者
	DeletedAt *time.Time `json:"deletedAt"` // 削除日時
	DeletedBy string     `json:"deletedBy"` // 削除者
	Version   int        `json:"version"`   // 楽観的排他ロック用バージョン
}

// IsDeleted - 削除済みかどうかを返す
func (m *Meta) IsDeleted() bool {
	return m.DeletedAt != nil
}
