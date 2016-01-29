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
package soul

import (
	"bytes"
	"encoding/json"
	"time"
)

const (
	PUBLISH = iota
	DRAFT
	IDEA
)

const (
	ID_FIELD              = "Id"
	PERMALINK_FIELD       = "Permalink"
	TITLE_FIELD           = "Title"
	SLUG_FIELD            = "Slug"
	CONTENT_FIELD         = "Content"
	STATE_FIELD           = "State"
	PUBLICATIONDATE_FIELD = "PublicationDate"
)

var postState = [...]string{
	"publish",
	"draft",
	"idea",
}

type Post struct {
	Id              string    `json:"Id"`
	Permalink       string    `json:"Permalink"`
	Title           string    `json:"Title"`
	Slug            string    `json:"Slug"`
	Content         string    `json:"Content"`
	State           string    `json:"State"`
	PublicationDate time.Time `json:"PublicationDate"`
}

type PostUtil interface {
	Json(Post) (string, error)
}

type PostBuilder interface {
	Id(string) PostBuilder
	Permalink(string) PostBuilder
	Title(string) PostBuilder
	Slug(string) PostBuilder
	Content(string) PostBuilder
	State(string) PostBuilder
	PublicationDate(time.Time) PostBuilder

	Json(string) (Post, error)
	Build() (Post, error)
}

type postBuilder struct {
	id              string
	permalink       string
	title           string
	slug            string
	content         string
	state           string
	publicationDate time.Time
}

func (post *postBuilder) Id(id string) PostBuilder {
	post.id = id
	return post
}
func isValidId(id string) (string, error) {
	return id, nil
}

func (post *postBuilder) Permalink(permalink string) PostBuilder {
	post.permalink = permalink
	return post
}
func isValidPermalink(permalink string) (string, error) {
	return permalink, nil
}
func (post *postBuilder) Title(title string) PostBuilder {
	post.title = title
	return post
}
func isValidTitle(title string) (string, error) {
	return title, nil
}
func (post *postBuilder) Slug(slug string) PostBuilder {
	post.slug = slug
	return post
}
func isValidSlug(slug string) (string, error) {
	return slug, nil
}
func (post *postBuilder) Content(content string) PostBuilder {
	post.content = content
	return post
}
func isValidContent(content string) (string, error) {
	return content, nil
}
func (post *postBuilder) State(state string) PostBuilder {
	post.state = state
	return post
}
func isValidState(state string) (string, error) {
	return state, nil
}
func (post *postBuilder) PublicationDate(publicationDate time.Time) PostBuilder {
	post.publicationDate = publicationDate
	return post
}
func isValidPublicationDate(publicationDate time.Time) (time.Time, error) {
	return publicationDate, nil
}

func (post *postBuilder) Json(jsonData string) (Post, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		return Post{}, err
	}
	return post.Build()
}

func (post *postBuilder) Build() (Post, error) {
	var p Post
	var err error

	p.Id, err = isValidId(post.id)
	p.Permalink, err = isValidPermalink(post.permalink)
	p.Title, err = isValidTitle(post.title)
	p.Slug, err = isValidSlug(post.slug)
	p.Content, err = isValidContent(post.content)
	p.State, err = isValidState(post.state)
	p.PublicationDate, err = isValidPublicationDate(post.publicationDate)

	return p, err
}

func NewPostBuilder() PostBuilder {
	return &postBuilder{}
}

func NewPost(permalink, title, content, state string, publicationDate time.Time) Post {

	post, _ := NewPostBuilder().
		Id("").
		Permalink(permalink).
		Title(title).
		Slug(createSlugFromTitle(title)).
		Content(content).
		State(state).
		PublicationDate(publicationDate).
		Build()

	if post.Permalink == "" {
		post.Permalink = post.Slug
	} else {
		post.Slug = post.Permalink
	}

	return post
}

func (post *Post) PostToJson() (string, error) {
	var compactPost *bytes.Buffer
	postJson, err := json.Marshal(post)
	if err != nil {
		// TODO should panic ???!!!
	}
	json.Compact(compactPost, postJson)
	return compactPost.String(), err
}
