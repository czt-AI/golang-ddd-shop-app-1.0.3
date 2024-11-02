package domain

// Product defines the structure for a product entity.
type Product struct {
	ProductId    int
	ProductName  string
	Description  string
	Price        decimal.Decimal
	Stock        int
	CategoryId   int
	Images       []string
	Reviews      []Review
}

// NewProduct creates a new Product instance.
func NewProduct(productId int, productName, description string, price decimal.Decimal, stock int, categoryId int, images []string, reviews []Review) *Product {
	return &Product{
		ProductId:    productId,
		ProductName:  productName,
		Description:  description,
		Price:        price,
		Stock:        stock,
		CategoryId:   categoryId,
		Images:       images,
		Reviews:      reviews,
	}
}

// ProductDTO defines the data transfer object for a product.
type ProductDTO struct {
	ProductId    int
	ProductName  string
	Description  string
	Price        float64
	Stock        int
	CategoryId   int
	Images       []string
	Reviews      []ReviewDTO
}

// ToDTO converts a Product to a ProductDTO.
func (p *Product) ToDTO() *ProductDTO {
	return &ProductDTO{
		ProductId:    p.ProductId,
		ProductName:  p.ProductName,
		Description:  p.Description,
		Price:        p.Price.Float64(),
		Stock:        p.Stock,
		CategoryId:   p.CategoryId,
		Images:       p.Images,
		Reviews:      make([]ReviewDTO, len(p.Reviews)),
	}
}

// Review defines the structure for a review entity.
type Review struct {
	ReviewId    int
	ProductId   int
	UserId      int
	Content     string
	CreateTime  time.Time
}

// NewReview creates a new Review instance.
func NewReview(reviewId, productId, userId int, content string, createTime time.Time) *Review {
	return &Review{
		ReviewId:    reviewId,
		ProductId:   productId,
		UserId:      userId,
		Content:     content,
		CreateTime:  createTime,
	}
}

// ReviewDTO defines the data transfer object for a review.
type ReviewDTO struct {
	ReviewId    int
	ProductId   int
	UserId      int
	Content     string
	CreateTime  time.Time
}
