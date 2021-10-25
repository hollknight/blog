package model

import "blog/pkg/app"

// Article database model
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

// ArticleSwagger swagger model
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
