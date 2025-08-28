package repository

import (
	"context"

	"apis_service/domain"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type AutoIncrement struct {
	col *mongo.Collection
	log *zap.SugaredLogger
}

type RepoAutoIncrement interface {
	GetIncrement(ctx context.Context, field string) (int32, error)
	GetCurrent(ctx context.Context, field string) (int32, error)
	Upsert(ctx context.Context, field string, value int32) error
}

func NewRepoAutoIncrement(
	dbConnection *mongo.Database,
	log *zap.SugaredLogger,
	collectionName string,
) *AutoIncrement {
	collection := dbConnection.Collection(collectionName)

	repo := &AutoIncrement{
		col: collection,
		log: log,
	}

	return repo
}

func (a *AutoIncrement) GetIncrement(ctx context.Context, field string) (int32, error) {
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	var doc *domain.AutoIncrement
	err := a.col.FindOneAndUpdate(
		ctx,
		bson.M{"_id": field},
		bson.M{"$inc": bson.M{"seq": 1}},
		opts,
	).Decode(&doc)
	if err != nil {
		return 0, errors.Wrap(err, "GetIncrement FindOneAndUpdate Decode")
	}

	return doc.SEQ, nil
}

func (a *AutoIncrement) Upsert(ctx context.Context, field string, value int32) error {
	a.log.Warn("Upsert increment called with field: %s, value: %d, USE ONLY in DEVELOPMENT", field, value)

	opts := options.Update().SetUpsert(true)
	_, err := a.col.UpdateOne(ctx, bson.M{"_id": field}, bson.M{"$set": bson.M{"seq": value}}, opts)
	if err != nil {
		return errors.Wrap(err, "UpdateOne")
	}

	return nil
}

func (a *AutoIncrement) GetCurrent(ctx context.Context, field string) (int32, error) {
	var doc *domain.AutoIncrement
	err := a.col.FindOne(
		ctx,
		bson.M{"_id": field},
	).Decode(&doc)
	if err != nil {
		return 0, errors.Wrap(err, "GetCurrent FindOne Decode")
	}

	return doc.SEQ, nil
}
