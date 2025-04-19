package pagination

import "math"

// Paginator обрабатывает параметры пагинации
type Paginator struct {
	Page     int
	PageSize int
}

// NewPaginator создает новый экземпляр Paginator с валидацией
func NewPaginator(page, pageSize int) Paginator {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10 // значение по умолчанию
	}
	if pageSize > 100 {
		pageSize = 100 // максимальное значение
	}

	return Paginator{
		Page:     page,
		PageSize: pageSize,
	}
}

// Offset вычисляет смещение для SQL-запросов
func (p Paginator) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit возвращает лимит для SQL-запросов
func (p Paginator) Limit() int {
	return p.PageSize
}

// TotalPages вычисляет общее количество страниц
func (p Paginator) TotalPages(totalItems int) int {
	return int(math.Ceil(float64(totalItems) / float64(p.PageSize)))
}

// IsValid проверяет валидность параметров пагинации
func (p Paginator) IsValid(totalItems int) bool {
	if p.Page < 1 || p.PageSize < 1 {
		return false
	}
	return p.Offset() < totalItems
}
