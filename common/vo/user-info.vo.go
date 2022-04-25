package vo

import (
	"time"
)

type UserInfo struct {
	ID int `json:"id,omitempty"`

	CreatedAt time.Time `json:"createdAt"`

	UpdatedAt time.Time `json:"updatedAt"`

	DeletedAt *time.Time `json:"deletedAt"`

	Username string `json:"username"`

	Avatar string `json:"avatar"`

	Phone *string `json:"phone"`
	
	Email *string `json:"email"`
}
