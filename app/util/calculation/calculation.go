package calculation

import (
	"errors"
	"math"
	"strings"
)

// Add は2つの数値を足し算する関数
func Add(a, b int) int {
	return a + b
}

// Subtract は2つの数値を引き算する関数
func Subtract(a, b int) int {
	return a - b
}

// Multiply は2つの数値を掛け算する関数
func Multiply(a, b int) int {
	// 乗算結果を返す
	return a * b
}

// Divide は2つの数値を割り算する関数
// 0による割り算の場合はエラーを返す
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ゼロによる除算はできません")
	}
	return a / b, nil
}

// CalculateAverage はスライスの平均値を計算する
func CalculateAverage(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("空のスライスの平均値は計算できません")
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return float64(sum) / float64(len(numbers)), nil
}

// IsEven は数値が偶数かどうかを判定する
func IsEven(number int) bool {
	// 2で割り切れるかどうかで判定
	return number%2 == 0
}

// CountWords はテキスト内の単語数を数える
func CountWords(text string) int {
	if text == "" {
		return 0
	}

	// 空白で分割して単語をカウント
	words := strings.Fields(text)
	return len(words)
}

// CalculateCircleArea は円の面積を計算する
func CalculateCircleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("半径は負の値にできません")
	}

	return math.Pi * radius * radius, nil
}
