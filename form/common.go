package form

import "time"

type BaseModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type BaseListResp struct {
	Total         int64       `json:"total"`
	BaseModelList []BaseModel `json:"items"`
}
