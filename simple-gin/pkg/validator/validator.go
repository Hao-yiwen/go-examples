// Package validator 提供通用的数据验证工具
// 可被其他项目导入使用
package validator

import (
	"regexp"
	"strings"
)

// 预编译正则表达式提升性能
var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

// IsValidPhone 验证中国手机号格式
func IsValidPhone(phone string) bool {
	return phoneRegex.MatchString(phone)
}

// IsNotEmpty 检查字符串是否非空（去除空格后）
func IsNotEmpty(s string) bool {
	return strings.TrimSpace(s) != ""
}

// IsInRange 检查数字是否在范围内
func IsInRange(value, min, max int) bool {
	return value >= min && value <= max
}

// IsPositive 检查是否为正数
func IsPositive(value float64) bool {
	return value > 0
}

// IsNonNegative 检查是否为非负数
func IsNonNegative(value float64) bool {
	return value >= 0
}

// MinLength 检查字符串最小长度
func MinLength(s string, min int) bool {
	return len(s) >= min
}

// MaxLength 检查字符串最大长度
func MaxLength(s string, max int) bool {
	return len(s) <= max
}

// LengthBetween 检查字符串长度是否在范围内
func LengthBetween(s string, min, max int) bool {
	length := len(s)
	return length >= min && length <= max
}
