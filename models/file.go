package models

type File struct {
	Name      string `json:"name"`
	Content   string `json:"content,omitempty"`
	CreatedOn string `json:"createdOn"`
}
