package domain

type Shop struct {
	ID        int     `json:"id"`
	UID       string  `json:"uid"`
	Name      string  `json:"name"`
	Direction string  `json:"dir"`
	Commune   string  `json:"com"`
	latitude  float64 `json:"lat"`
	longitude float64 `json:"lng"`
}

type Shops []Shop
