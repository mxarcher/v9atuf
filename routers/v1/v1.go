package v1

// 使用泛型简化代码
type info[T any] struct {
	Operator string `json:"operator"`
	Model    T      `json:"model"`
}
