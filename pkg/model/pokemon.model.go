package model

type Pokemon struct {
	Abilities      []string `bson:"abilities"`
	Weight         float64  `bson:"weight"`
	Weakness       []string `bson:"weakness"`
	Height         int      `bson:"height"`
	Slug           string   `bson:"slug"`
	Name           string   `bson:"name"`
	ThumbnailImage string   `bson:"ThumbnailImage"`
	ID             int      `bson:"_id,omitempty"`
	Type           []string `bson:"type"`
}
