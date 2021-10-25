package model

import "blog/pkg/app"

// Tag database model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// TagSwagger swagger model
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
