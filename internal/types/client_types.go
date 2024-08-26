package types

type ErrorResponse struct {
	Error Error `json:"error,omitempty"`
}

type Error struct {
	Message        string `json:"message,omitempty"`
	DisplayMessage string `json:"displayMessage,omitempty"`
	Code           string `json:"code,omitempty"`
}

type ReverseImageGeneratorRequest struct {
	ImageUrl string `json:"imageUrl" binding:"required"`
	Page     int    `json:"page"`
}

type ReverseImageGeneratorResponse struct {
	Products []Product
}
