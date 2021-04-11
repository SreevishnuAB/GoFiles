package models

type File struct {
	Name      string      `json:"name"`
	Content   interface{} `json:"content,omitempty"`
	CreatedOn string      `json:"createdOn"`
}
