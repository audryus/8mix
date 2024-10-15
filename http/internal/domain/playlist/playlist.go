package playlist

import "go.mongodb.org/mongo-driver/bson/primitive"

type Playlist struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	User   string             `json:"user" bson:"user"`
	Tracks []string           `json:"tracks" bson:"tracks"`
	Status string             `json:"status" bson:"status"`
}
