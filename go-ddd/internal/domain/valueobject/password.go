package valueobject

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters")
	ErrPasswordTooWeak    = errors.New("password must contain at least one uppercase, one lowercase, and one digit")
	ErrPasswordHashFailed = errors.New("failed to hash password")
	ErrPasswordMismatch   = errors.New("password does not match")
)

// Password 密码值对象
// 值对象封装了密码的创建和验证逻辑
// 存储的是加密后的哈希值，而不是明文
type Password struct {
	hash string
}

// NewPassword 从明文创建密码值对象
func NewPassword(plaintext string) (Password, error) {
	// 验证密码强度
	if len(plaintext) < 8 {
		return Password{}, ErrPasswordTooShort
	}

	if !isStrongPassword(plaintext) {
		return Password{}, ErrPasswordTooWeak
	}

	// 使用bcrypt加密
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, ErrPasswordHashFailed
	}

	return Password{hash: string(hashedBytes)}, nil
}

// NewPasswordFromHash 从哈希值创建密码值对象（用于从数据库读取）
func NewPasswordFromHash(hash string) Password {
	return Password{hash: hash}
}

// Hash 返回密码哈希值
func (p Password) Hash() string {
	return p.hash
}

// Verify 验证明文密码是否匹配
func (p Password) Verify(plaintext string) error {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plaintext))
	if err != nil {
		return ErrPasswordMismatch
	}
	return nil
}

// IsEmpty 检查密码是否为空
func (p Password) IsEmpty() bool {
	return p.hash == ""
}

// isStrongPassword 检查密码强度
func isStrongPassword(password string) bool {
	var hasUpper, hasLower, hasDigit bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	return hasUpper && hasLower && hasDigit
}
