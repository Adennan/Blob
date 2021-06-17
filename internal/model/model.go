package model

import (
	"time"
)

// public model
type Model struct {
	ID        uint32    `gorm:"primary_key" json:"id"`
	CreateAt  time.Time `json:"create_at"`
	CreateBy  string    `json:"create_by"`
	UpdateAt  time.Time `json:"update_at"`
	UpdateBy  string    `json:"update_by"`
	DeleteAt  time.Time `json:"delete_at"`
	DelStatus int       `json:"del_status"`
}

type Tag struct {
	Model
	Name   string `json:"name"`
	Status int    `json:"status"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type Article struct {
	Model
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	ImageUrl string `json:"image_url"`
	Status   int    `json:"status"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleTag struct {
	Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (at ArticleTag) TableName() string {
	return "article_tag"
}
