package soulObjects

import (
	"time"
)

type Post struct {
	Id               string    `json:"Id"`	
	Permalink		 string	   `json:"Permalink"`
	Title           string    `json:"Title"`
	Slug			 string	   `json:"Slug"`
	Content        string    `json:"Content"`
	PublicationDate time.Time `json:"PublicationDate"`
}

func NewPost(permalink, title, content string, pulicationDate time.Time) Post {
	post := Post{
		Id: "",
		Permalink: permalink,
		Title: title,
		Slug: createSlugFromTitle(title),
		Content: content,
		PublicationDate: pulicationDate,
	}
	
	if post.Permalink == "" {
		post.Permalink = post.Slug
	} else {
		post.Slug = post.Permalink
	}
		
	return post
}