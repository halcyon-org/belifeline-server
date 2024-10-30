package domain

import "github.com/halcyon-org/kizuna/ent/schema/pulid"

type Point struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type Value struct {
	Point    Point             `json:"point"`
	Value    float32           `json:"value"`
	Optional map[string]string `json:"optional"`
}

type Content struct {
	ID     pulid.ID `json:"id"`
	Values []Value  `json:"values"`
}
