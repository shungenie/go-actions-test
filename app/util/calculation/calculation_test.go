package calculation

import (
	"math"
	"testing"
)

// Addのテスト
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"正の数を足す", 5, 3, 8},
		{"負の数を足す", -5, -3, -8},
		{"正と負を足す", 5, -3, 2},
		{"ゼロを足す", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Subtractのテスト
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"正の数同士", 8, 3, 5},
		{"負の数同士", -5, -3, -2},
		{"正と負", 5, -3, 8},
		{"ゼロとの計算", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Multiplyのテスト
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"正の数同士", 4, 3, 12},
		{"負の数同士", -4, -3, 12},
		{"正と負", 4, -3, -12},
		{"ゼロとの計算", 4, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Divideのテスト
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        int
		expected    int
		expectError bool
	}{
		{"正の数同士", 12, 3, 4, false},
		{"負の数同士", -12, -3, 4, false},
		{"正と負", 12, -3, -4, false},
		{"ゼロによる除算", 12, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			// エラーチェック
			if (err != nil) != tt.expectError {
				t.Errorf("Divide(%d, %d) error = %v; expectError %v", tt.a, tt.b, err, tt.expectError)
				return
			}

			// エラーが期待されない場合は結果も確認
			if !tt.expectError && result != tt.expected {
				t.Errorf("Divide(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// CalculateAverageのテスト
func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expected    float64
		expectError bool
	}{
		{"通常の計算", []int{1, 2, 3, 4, 5}, 3.0, false},
		{"負の数を含む", []int{-1, -2, 3, 4}, 1.0, false},
		{"空のスライス", []int{}, 0.0, true},
		{"同じ値のみ", []int{5, 5, 5}, 5.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculateAverage(tt.input)

			// エラーチェック
			if (err != nil) != tt.expectError {
				t.Errorf("CalculateAverage(%v) error = %v; expectError %v", tt.input, err, tt.expectError)
				return
			}

			// エラーが期待されない場合は結果も確認
			if !tt.expectError && math.Abs(result-tt.expected) > 0.0001 {
				t.Errorf("CalculateAverage(%v) = %f; expected %f", tt.input, result, tt.expected)
			}
		})
	}
}

// IsEvenのテスト
func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"偶数のケース", 4, true},
		{"奇数のケース", 5, false},
		{"ゼロのケース", 0, true},
		{"負の偶数", -6, true},
		{"負の奇数", -7, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("IsEven(%d) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

// CountWordsのテスト
func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"標準的な文", "これは テスト です", 3},
		{"空白の多い文", "  これは  テスト   です  ", 3},
		{"空文字列", "", 0},
		{"1つの単語", "テスト", 1},
		{"複数行", "これは\nテスト\nです", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWords(tt.input)
			if result != tt.expected {
				t.Errorf("CountWords(%q) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}
