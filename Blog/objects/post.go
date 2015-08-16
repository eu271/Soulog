/*
	Copyright (c) 2015 Eugenio Ochoa

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/
package soulObjects

import (
	"time"
)

type Post struct {
	Id              string    `json:"Id"`
	Permalink       string    `json:"Permalink"`
	Title           string    `json:"Title"`
	Slug            string    `json:"Slug"`
	Content         string    `json:"Content"`
	PublicationDate time.Time `json:"PublicationDate"`
}

type PostBuilder interface {
	Id(string) PostBuilder
	Permalink(string) PostBuilder
	Title(string) PostBuilder
	Slug(string) PostBuilder
	Content(string) PostBuilder
	PublicationDate(time.Time) PostBuilder

	Build() (Post, error)
}

type postBuilder struct {
	id              string
	permalink       string
	title           string
	slug            string
	content         string
	publicationDate time.Time
}

func (post *postBuilder) Id(id string) PostBuilder {
	post.id = id
	return post
}
func (post *postBuilder) Permalink(permalink string) PostBuilder {
	post.permalink = permalink
	return post
}
func (post *postBuilder) Title(title string) PostBuilder {
	post.title = title
	return post
}
func (post *postBuilder) Slug(slug string) PostBuilder {
	post.slug = slug
	return post
}
func (post *postBuilder) Content(content string) PostBuilder {
	post.content = content
	return post
}
func (post *postBuilder) PublicationDate(publicationDate time.Time) PostBuilder {
	post.publicationDate = publicationDate
	return post
}

func (post *postBuilder) Build() (Post, error) {
	p := Post{
		Id:              post.id,
		Permalink:       post.permalink,
		Title:           post.title,
		Slug:            createSlugFromTitle(post.title),
		Content:         post.content,
		PublicationDate: post.publicationDate,
	}
	return p, nil
}

func NewPostBuilder() PostBuilder {
	return &postBuilder{}
}

func NewPost(permalink, title, content string, publicationDate time.Time) Post {

	post, _ := NewPostBuilder().Id("").
		Permalink(permalink).
		Title(title).
		Slug(createSlugFromTitle(title)).
		Content(content).
		PublicationDate(publicationDate).
		Build()

	if post.Permalink == "" {
		post.Permalink = post.Slug
	} else {
		post.Slug = post.Permalink
	}

	return post
}
