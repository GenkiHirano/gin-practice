package handler

type requestUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// コメントの上限を10文字にバリデーションしてください。
	Comment string `json:"comment"`
}

type responseUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

// CreateUserを実装してください。
// 引数、戻り値、関数の中身はお好きに変更可能です。
func CreateUser() {
}

// GetUserを実装してください。
// 引数、戻り値、関数の中身はお好きに変更可能です。
func GetUser() {
}
