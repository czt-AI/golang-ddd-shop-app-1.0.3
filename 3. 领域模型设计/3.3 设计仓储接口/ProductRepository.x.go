package domain

// ProductRepository defines the interface for product data access.
type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	FindById(ctx context.Context, productId int) (*Product, error)
	FindAll(ctx context.Context) ([]Product, error)
	FindByCategory(ctx context.Context, categoryId int) ([]Product, error)
	Delete(ctx context.Context, productId int) error
}

// InMemoryProductRepository is an in-memory implementation of ProductRepository.
type InMemoryProductRepository struct {
	products map[int]*Product
}

// Save saves a product to the in-memory store.
func (repo *InMemoryProductRepository) Save(ctx context.Context, product *Product) error {
	repo.products[product.ProductId] = product
	return nil
}

// FindById finds a product by ID.
func (repo *InMemoryProductRepository) FindById(ctx context.Context, productId int) (*Product, error) {
	product, exists := repo.products[productId]
	if !exists {
		return nil, fmt.Errorf("product not found with ID: %d", productId)
　　}
	return product, nil
}

// FindAll finds all products.
func (repo *InMemoryProductRepository) FindAll(ctx context.Context) ([]Product, error) {
	products := make([]Product, 0, len(repo.products))
	for _, product := range repo.products {
		products = append(products, *product)
	}
	return products, nil
}

// FindByCategory finds products by category ID.
func (repo *InMemoryProductRepository) FindByCategory(ctx context.Context, categoryId int) ([]Product, error) {
	products := make([]Product, 0)
	for _, product := range repo.products {
		if product.CategoryId == categoryId {
			products = append(products, *product)
		}
	}
	return products, nil
}

// Delete deletes a product from the in-memory store.
func (repo *InMemoryProductRepository) Delete(ctx context.Context, productId int) error {
	if _, exists := repo.products[productId]; !exists {
		return fmt.Errorf("product not found with ID: %d", productId)
	}
	delete(repo.products, productId)
	return nil
}
