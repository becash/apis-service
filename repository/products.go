package repository

import (
	"context"

	"apis_service/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Products struct {
	col *mongo.Collection
	log *zap.SugaredLogger
}

func NewRepoProducts(
	dbConnection *mongo.Database,
	log *zap.SugaredLogger,
) *Products {
	col := dbConnection.Collection("products")
	repo := &Products{
		col: col,
		log: log,
	}

	return repo
}

func (b *Products) Get(
	ctx context.Context,
	input int32,
	fields *bson.M,
) (*domain.Product, error) {
	var res *domain.Product
	err := b.col.FindOne(ctx, bson.M{"_id": input}).Decode(&res)
	return res, errors.Wrap(err, "BannersRepo: GetOne Decode")
}

//func (b *Products) Count(
//	ctx context.Context,
//	input *domain.BannersListRequest,
//) (int64, error) {
//	filter := mapper.ConvertStructToBSONMap(input, nil)
//	// rm deleted banners from count if arg not provided
//	if input.Status == domain.BannerStatusUnspecified {
//		if filter == nil {
//			filter = primitive.M{}
//		}
//
//		filter["status"] = primitive.M{
//			"$ne": domain.BannerStatusDeleted,
//		}
//	}
//
//	return b.col.CountDocuments(ctx, filter)
//}

//func (b *Products) List(
//	ctx context.Context,
//	input *domain.BannersListRequest,
//	fields *bson.M,
//) ([]*domain.Banner, error) {
//	var banners []*domain.Banner
//
//	opts := options.Find()
//	database.FindFillPagination(opts, input.Pagination)
//	filter := mapper.ConvertStructToBSONMap(input, nil)
//
//	if fields != nil && len(*fields) > 0 {
//		opts = opts.SetProjection(*fields)
//	}
//
//	// rm deleted banners from result
//	if input.Status == domain.BannerStatusUnspecified {
//		if filter == nil {
//			filter = primitive.M{}
//		}
//
//		filter["status"] = primitive.M{
//			"$ne": domain.BannerStatusDeleted,
//		}
//	}
//
//	cursor, err := b.col.Find(ctx, filter, opts)
//	if err != nil {
//		return nil, errors.Wrap(err, "BannersRepo List Find")
//	}
//	defer cursor.Close(ctx)
//
//	err = cursor.All(ctx, &banners)
//
//	return banners, errors.Wrap(err, "BannersRepo List All()")
//}
//
//func (b *Products) Delete(
//	ctx context.Context,
//	input int32,
//) error {
//	_, err := b.col.UpdateByID(ctx, input, bson.M{"$set": bson.M{"status": domain.BannerStatusDeleted}})
//
//	return errors.Wrap(err, "BannersRepo Delete UpdateByID")
//}
//
//func (b *Products) Upsert(
//	ctx context.Context,
//	input *domain.BannerUpsertRequest,
//) (*domain.Banner, error) {
//	var res *domain.Banner
//
//	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
//	err := b.col.FindOneAndUpdate(
//		ctx,
//		bson.M{"_id": input.ID},
//		bson.M{"$set": input},
//		opts,
//	).Decode(&res)
//
//	return res, errors.Wrap(err, "BannersRepo Upsert Decode")
//}
