package errorkey

// MessageKeyName - エラーメッセージ
type MessageKeyName = string

// common
const (
	Internal     MessageKeyName = "common.errorMessage.internal"
	Conflict     MessageKeyName = "common.errorMessage.conflict"
	NotFound     MessageKeyName = "common.errorMessage.notFound"
	BadRequest   MessageKeyName = "common.errorMessage.badRequest"
	Unauthorized MessageKeyName = "common.errorMessage.Unauthorized"
	Forbidden    MessageKeyName = "common.errorMessage.Forbidden"
)
