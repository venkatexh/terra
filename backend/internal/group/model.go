package group

import "time"

type Group struct {
	ID string `json:"id"`
	Name string `json:"name"`
	UserId string `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

