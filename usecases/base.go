package usecases

import (
	"context"

	"apis_service/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type UseCases struct {
	cfg               *domain.Config
	log               *zap.SugaredLogger
	repoProducts      RepoProducts
	repoAutoIncrement RepoAutoIncrement
}

type RepoAutoIncrement interface {
	GetIncrement(ctx context.Context, field string) (int32, error)
	GetCurrent(ctx context.Context, field string) (int32, error)
	Upsert(ctx context.Context, field string, value int32) error
}

type RepoProducts interface {
	Get(ctx context.Context, input int32, fields *bson.M) (*domain.Product, error)
	//Upsert(ctx context.Context, input *domain.BannerUpsertRequest) (*domain.Banner, error)
	//List(ctx context.Context, input *domain.BannersListRequest, fields *bson.M) ([]*domain.Banner, error)
	//Delete(ctx context.Context, input int32) error
	//Count(ctx context.Context, input *domain.BannersListRequest) (int64, error)
}

func NewUseCases(
	cfg *domain.Config,
	log *zap.SugaredLogger,
	repoAutoIncrement RepoAutoIncrement,
	repoProducts RepoProducts,
) *UseCases {
	return &UseCases{
		cfg:               cfg,
		log:               log,
		repoAutoIncrement: repoAutoIncrement,
		repoProducts:      repoProducts,
	}
}
