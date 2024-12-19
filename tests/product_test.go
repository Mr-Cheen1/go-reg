package tests

import (
	"testing"

	"github.com/Mr-Cheen1/go-reg/models"
)

func TestProductSearch(t *testing.T) {
	products := models.Products{
		{ID: 1, Name: "Тест продукт 1"},
		{ID: 2, Name: "Другой продукт"},
		{ID: 3, Name: "Тестовый образец"},
	}

	tests := []struct {
		name     string
		query    string
		expected int
	}{
		{"Пустой запрос", "", 3},
		{"Поиск по 'тест'", "тест", 2},
		{"Поиск по 'другой'", "другой", 1},
		{"Поиск несуществующего", "xyz", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := products.Search(tt.query)
			if len(result) != tt.expected {
				t.Errorf("Search() = %v результатов, ожидалось %v", len(result), tt.expected)
			}
		})
	}
}

func TestProductDelete(t *testing.T) {
	tests := []struct {
		name        string
		deleteID    int
		initialLen  int
		expectedLen int
	}{
		{
			name:        "Удаление существующего",
			deleteID:    2,
			initialLen:  3,
			expectedLen: 2,
		},
		{
			name:        "Удаление несуществующего",
			deleteID:    99,
			initialLen:  2,
			expectedLen: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products := models.Products{
				{ID: 1, Name: "Продукт 1"},
				{ID: 2, Name: "Продукт 2"},
				{ID: 3, Name: "Продукт 3"},
			}[:tt.initialLen]

			products.Delete(tt.deleteID)
			if len(products) != tt.expectedLen {
				t.Errorf("Delete() = %v элементов, ожидалось %v", len(products), tt.expectedLen)
			}

			for _, p := range products {
				if p.ID == tt.deleteID {
					t.Errorf("Delete() не удалил элемент с ID %v", tt.deleteID)
				}
			}
		})
	}
}

func TestGetNextID(t *testing.T) {
	tests := []struct {
		name     string
		products models.Products
		expected int
	}{
		{
			name:     "Пустой список",
			products: models.Products{},
			expected: 1,
		},
		{
			name: "Последовательные ID",
			products: models.Products{
				{ID: 1}, {ID: 2}, {ID: 3},
			},
			expected: 4,
		},
		{
			name: "Непоследовательные ID",
			products: models.Products{
				{ID: 1}, {ID: 5}, {ID: 3},
			},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.products.GetNextID()
			if result != tt.expected {
				t.Errorf("GetNextID() = %v, ожидалось %v", result, tt.expected)
			}
		})
	}
}
