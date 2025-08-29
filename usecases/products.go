package usecases

import (
	"apis_service/domain"
	"context"
)

func (u *UseCases) GetProduct(
	ctx context.Context,
	input int32,
) (*domain.Product, error) {
	return u.repoProducts.Get(ctx, input, nil)
}
