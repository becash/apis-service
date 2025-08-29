package api

import (
	"context"

	"github.com/becash/apis/gen_go/swallow_channel_to_service"
	"github.com/pkg/errors"
)

func (s *Server) GetAvailabilityOfProduct(
	ctx context.Context,
	input *swallow_channel_to_service.ProductAvailabilitiesRequest,
) (*swallow_channel_to_service.Availabilities, error) {
	//res, err := s.useCases.GetProduct(ctx, input.GetId())
	//return s.converter.FromBannerPublic(res), errors.Wrap(err, "Server GetBannerPublic")
	//return nil, errors.Wrap(err, "Server GetBannerPublic")
	return nil, errors.Wrap(nil, "Server GetBannerPublic")
}
