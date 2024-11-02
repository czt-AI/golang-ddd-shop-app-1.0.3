package domain

// Category defines the structure for a category entity.
type Category struct {
	CategoryId int
	CategoryName string
}

// NewCategory creates a new Category instance.
func NewCategory(categoryId int, categoryName string) *Category {
	return &Category{
		CategoryId: categoryId,
		CategoryName: categoryName,
	}
}

// CategoryDTO defines the data transfer object for a category.
type CategoryDTO struct {
	CategoryId int
	CategoryName string
}

// ToDTO converts a Category to a CategoryDTO.
func (c *Category) ToDTO() *CategoryDTO {
	return &CategoryDTO{
		CategoryId: c.CategoryId,
		CategoryName: c.CategoryName,
	}
}
