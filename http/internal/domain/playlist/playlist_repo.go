package playlist

import (
	"context"

	db "github.com/audryus/8mix/http/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlaylistRepo struct {
	col *mongo.Collection
}

func NewPlaylistRepo(mongo *db.Mongo) *PlaylistRepo {
	return &PlaylistRepo{
		col: mongo.Collection("playlist"),
	}
}

func (r *PlaylistRepo) Create(ctx context.Context, playlist *Playlist) (*Playlist, error) {
	playlist.ID = primitive.NewObjectID()

	_, err := r.col.InsertOne(ctx, playlist)

	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (r *PlaylistRepo) Find(ctx context.Context, playlist *Playlist) (*Playlist, error) {
	filter := bson.D{{Key: "_id", Value: playlist.ID}}

	result := new(Playlist)

	if err := r.col.FindOne(ctx, filter).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
