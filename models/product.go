package models

import "strings"

// Product представляет собой структуру продукта
type Product struct {
	ID              int
	Name            string
	ProcessingTime  float64
	TimeCalculation string
}

// Products представляет собой срез продуктов с методами для работы
type Products []Product

// Search ищет продукты по запросу
func (p Products) Search(query string) Products {
	if query == "" {
		return p
	}

	var result Products
	query = strings.ToLower(query)
	for _, product := range p {
		if strings.Contains(strings.ToLower(product.Name), query) {
			result = append(result, product)
		}
	}
	return result
}

// Delete удаляет продукт по ID
func (p *Products) Delete(id int) {
	for i, product := range *p {
		if product.ID == id {
			*p = append((*p)[:i], (*p)[i+1:]...)
			break
		}
	}
}

// Update обновляет продукт
func (p *Products) Update(product Product) {
	for i, prod := range *p {
		if prod.ID == product.ID {
			(*p)[i] = product
			break
		}
	}
}

// GetNextID возвращает следующий доступный ID
func (p Products) GetNextID() int {
	maxID := 0
	for _, product := range p {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}
