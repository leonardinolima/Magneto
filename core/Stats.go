package core

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stats struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Dna string `json:"_dna,omitempty" bson:"_dna,omitempty"`
	IsMutant bool `json:"_isMutant,omitempty" bson:"_isMutant,omitempty"`
}
