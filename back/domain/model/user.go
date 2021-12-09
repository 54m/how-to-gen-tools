package model

import (
	"context"

	"github.com/54m/how-to-gen-tools/back/pkg/validator"
	"golang.org/x/xerrors"
)

// User - ユーザー
type User struct {
	ID   string `json:"id"   firestore:"-"` // ID
	Name string `json:"name"`               // ニックネーム
	Age  int    `json:"age"`                // 年齢
	Meta
}

// Prepare - 新規/更新とValidation
func (obj User) Prepare(context.Context) (*User, error) {
	if err := validator.Validate(obj); err != nil {
		return nil, xerrors.Errorf("validation error: %w", err)
	}

	return &obj, nil
}
