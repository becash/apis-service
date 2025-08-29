package api

import (
	"context"
	"fmt"

	"github.com/becash/apis/gen_go/swallow_channel_to_service"
)

func (s *Server) GetAvailabilityOfProduct(
	ctx context.Context,
	input *swallow_channel_to_service.ProductAvailabilitiesRequest,
) (*swallow_channel_to_service.Availabilities, error) {

	//return s.useCases.GetProductAvailabilities(ctx, input.GetProductId())
	return nil, fmt.Errorf("not implemented")
}
