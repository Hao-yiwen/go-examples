package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	emailRegex      = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// Email 邮箱值对象
// 值对象是DDD中的重要概念，特点是：
// 1. 不可变（immutable）
// 2. 没有唯一标识
// 3. 相等性由所有属性决定
// 4. 可以包含验证逻辑
type Email struct {
	value string
}

// NewEmail 创建邮箱值对象
func NewEmail(email string) (Email, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if !emailRegex.MatchString(email) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: email}, nil
}

// String 返回邮箱字符串
func (e Email) String() string {
	return e.value
}

// Equals 比较两个邮箱是否相等
func (e Email) Equals(other Email) bool {
	return e.value == other.value
}

// Domain 返回邮箱域名
func (e Email) Domain() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

// LocalPart 返回邮箱本地部分（@之前的部分）
func (e Email) LocalPart() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

// IsEmpty 检查邮箱是否为空
func (e Email) IsEmpty() bool {
	return e.value == ""
}
