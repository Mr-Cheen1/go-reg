package tests

import (
	"testing"
	"github.com/Mr-Cheen1/go-reg/utils"
)

func TestCalculateTime(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"Простое сложение", "1.5 + 2.5", 4.0},
		{"Несколько чисел", "1 + 2 + 3", 6.0},
		{"С пробелами", "1 + 2.5 + 3.5", 7.0},
		{"Одно число", "5", 5.0},
		{"Пустая строка", "", 0.0},
		{"Некорректные данные", "abc + def", 0.0},
		{"Смешанные данные", "1.5 + abc + 2.5", 4.0},
		{"Отрицательные числа", "-1 + 3", 2.0},
		{"Дробные числа", "0.3 + 0.7", 1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := utils.CalculateTime(tt.input)
			if result != tt.expected {
				t.Errorf("CalculateTime(%q) = %v, ожидалось %v", 
					tt.input, result, tt.expected)
			}
		})
	}
} 