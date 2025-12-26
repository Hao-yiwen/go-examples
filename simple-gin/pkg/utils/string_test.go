package utils

import (
	"reflect"
	"testing"
)

func TestGenerateID(t *testing.T) {
	id1 := GenerateID(16)
	id2 := GenerateID(16)

	if len(id1) != 16 {
		t.Errorf("GenerateID(16) length = %d, want 16", len(id1))
	}

	if id1 == id2 {
		t.Error("GenerateID should generate unique IDs")
	}
}

func TestContains(t *testing.T) {
	// 测试 int 切片
	intSlice := []int{1, 2, 3, 4, 5}
	if !Contains(intSlice, 3) {
		t.Error("Contains should return true for existing element")
	}
	if Contains(intSlice, 6) {
		t.Error("Contains should return false for non-existing element")
	}

	// 测试 string 切片
	strSlice := []string{"a", "b", "c"}
	if !Contains(strSlice, "b") {
		t.Error("Contains should return true for existing string")
	}
	if Contains(strSlice, "d") {
		t.Error("Contains should return false for non-existing string")
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"with duplicates", []int{1, 2, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unique(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("Unique() length = %d, want %d", len(got), len(tt.want))
			}
		})
	}
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	double := func(x int) int { return x * 2 }

	result := Map(input, double)
	expected := []int{2, 4, 6, 8, 10}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Map() = %v, want %v", result, expected)
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	isEven := func(x int) bool { return x%2 == 0 }

	result := Filter(input, isEven)
	expected := []int{2, 4, 6}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Filter() = %v, want %v", result, expected)
	}
}

// 表驱动测试示例
func TestTrimAll(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello world", "helloworld"},
		{"  spaces  ", "spaces"},
		{"no spaces", "nospaces"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := TrimAll(tt.input); got != tt.want {
				t.Errorf("TrimAll(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
