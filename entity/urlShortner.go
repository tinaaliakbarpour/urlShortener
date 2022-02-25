package entity

//UrlCreationRequest contains request parameters for creating short url
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}
