package dto

// Output DTOs
type UserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type UserBeansOutput struct {
	User       UserOutput   `json:"user"`
	Beans      []BeanOutput `json:"beans"`
	Count      uint         `json:"count"`
	NextCursor *uint        `json:"next_cursor,omitempty"`
}
