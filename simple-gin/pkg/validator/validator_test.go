package validator

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid email", "test@example.com", true},
		{"valid email with subdomain", "test@mail.example.com", true},
		{"invalid - no @", "testexample.com", false},
		{"invalid - no domain", "test@", false},
		{"invalid - empty", "", false},
		{"invalid - spaces", "test @example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.email); got != tt.want {
				t.Errorf("IsValidEmail(%q) = %v, want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  bool
	}{
		{"valid phone", "13800138000", true},
		{"valid phone 2", "15912345678", true},
		{"invalid - too short", "1380013800", false},
		{"invalid - too long", "138001380001", false},
		{"invalid - wrong prefix", "12800138000", false},
		{"invalid - empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidPhone(tt.phone); got != tt.want {
				t.Errorf("IsValidPhone(%q) = %v, want %v", tt.phone, got, tt.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"non-empty", "hello", true},
		{"empty", "", false},
		{"only spaces", "   ", false},
		{"with spaces", "  hello  ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.s); got != tt.want {
				t.Errorf("IsNotEmpty(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestIsInRange(t *testing.T) {
	tests := []struct {
		name       string
		value, min, max int
		want       bool
	}{
		{"in range", 5, 1, 10, true},
		{"at min", 1, 1, 10, true},
		{"at max", 10, 1, 10, true},
		{"below min", 0, 1, 10, false},
		{"above max", 11, 1, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInRange(tt.value, tt.min, tt.max); got != tt.want {
				t.Errorf("IsInRange(%d, %d, %d) = %v, want %v",
					tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

// Benchmark 性能测试示例
func BenchmarkIsValidEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidEmail("test@example.com")
	}
}

func BenchmarkIsValidPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidPhone("13800138000")
	}
}
