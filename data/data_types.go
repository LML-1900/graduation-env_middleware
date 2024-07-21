package data

type LonLatPosition struct {
	Longitude float64 `bson:"longitude" json:"longitude"`
	Latitude  float64 `bson:"latitude" json:"latitude"`
}

type Crater struct {
	Position LonLatPosition
	CraterID string  `bson:"crater_id" json:"crater_id"`
	Width    float64 `bson:"width" json:"width"`
	Depth    float64 `bson:"depth" json:"depth"`
}
