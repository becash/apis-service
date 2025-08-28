package grpc

import (
	"context"

	"github.com/becash/apis/gen_go/swallow"
	"github.com/pkg/errors"
)

func (s *Server) GetProduct(
	ctx context.Context,
	input *swallow.FieldFilter,
) (*swallow.Product, error) {
	//res, err := s.useCases.GetProduct(ctx, input.GetId())
	//return s.converter.FromBannerPublic(res), errors.Wrap(err, "Server GetBannerPublic")
	//return nil, errors.Wrap(err, "Server GetBannerPublic")
	return nil, errors.Wrap(nil, "Server GetBannerPublic")
}
