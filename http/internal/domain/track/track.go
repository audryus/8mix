package track

type Track struct {
	Url     string `json:"url" bson:"_id"`
	Status  string `json:"status" bson:"status"`
	Storage string `json:"storage" bson:"storage"`
}
