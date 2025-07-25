package dto

// Response DTOs
type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type UserBeansResponse struct {
	User  UserResponse   `json:"user"`
	Beans []BeanResponse `json:"beans"`
	Count uint           `json:"count"`
}
