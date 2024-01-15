package usecase

import "github.com/GuilhermeReisOlinto/messageria_go/internal/entity"

type ListProductOutputDto struct {
	ID    string
	Name  string
	Price float64
}

type ListProductUseCase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductUseCase(productRepository entity.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{ProductRepository: productRepository}
}

func (u *ListProductUseCase) Execute() ([]*ListProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var productOutput []*ListProductOutputDto

	for _, product := range products {
		productOutput = append(productOutput, &ListProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productOutput, nil
}
