package model

type Page struct {
	Current int64 `json:"current,omitempty"`
	Total int64 `json:"total,omitempty"`
}
