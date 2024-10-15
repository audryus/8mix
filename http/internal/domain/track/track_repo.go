package track

import (
	"context"

	db "github.com/audryus/8mix/http/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type TrackRepo struct {
	mongo *db.Mongo
}

func NewTrackRepo(mongo *db.Mongo) *TrackRepo {
	return &TrackRepo{
		mongo: mongo,
	}
}

func (r *TrackRepo) Create(ctx context.Context, track *Track) (*Track, error) {
	col := r.mongo.Collection("track")

	if _, err := col.InsertOne(ctx, track); err != nil {
		return nil, err
	}

	return track, nil
}

func (r *TrackRepo) Find(ctx context.Context, track *Track) (*Track, error) {
	col := r.mongo.Collection("track")

	filter := bson.D{{Key: "_id", Value: track.Url}}

	result := new(Track)

	if err := col.FindOne(ctx, filter).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
