package types

type Product struct {
	Name       string  `json:"name,omitempty"`
	Price      float32 `json:"price,omitempty"`
	UserRating float32 `json:"userRating,omitempty"`
}

type GroqAIResp struct {
	Name   string
	Price  float32
	Rating float32
}
