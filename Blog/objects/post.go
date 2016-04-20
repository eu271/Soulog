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

//This package contains the definitions and functions to contruct and work with posts.
package soul

import (
	"bytes"
	"encoding/json"
	"regexp"
	"time"
)

//Enum for posts states.
const (
	PUBLISH = iota
	DRAFT
	IDEA
)

//Constants define for JSON and other formats transformations.
const (
	ID_FIELD              = "Id"
	PERMALINK_FIELD       = "Permalink"
	TITLE_FIELD           = "Title"
	SLUG_FIELD            = "Slug"
	CONTENT_FIELD         = "Content"
	STATE_FIELD           = "State"
	PUBLICATIONDATE_FIELD = "PublicationDate"
)

//Constants define for JSON and other formats transformations.
const (
	ID_FIELD_TYPE              = "string"
	PERMALINK_FIELD_TYPE       = "string"
	TITLE_FIELD_TYPE           = "string"
	SLUG_FIELD_TYPE            = "string"
	CONTENT_FIELD_TYPE         = "string"
	STATE_FIELD_TYPE           = "string"
	PUBLICATIONDATE_FIELD_TYPE = "time.Time"
)

//All possible post states
var postState = []string{
	"publish",
	"draft",
	"idea",
}

var validIdRegexpTest = regexp.MustCompile("[a-zA-Z0-9]")

//Post structure.
type Post struct {
	Id              string    `json:"Id"`
	Permalink       string    `json:"Permalink"`
	Title           string    `json:"Title"`
	Slug            string    `json:"Slug"`
	Content         string    `json:"Content"`
	State           string    `json:"State"`
	PublicationDate time.Time `json:"PublicationDate"`
}

//Some util Post functions.
type PostUtil interface {
	//Transforms a post to JSON format.
	Json(Post) (string, error)
	//Returns all possible post states.
	PostStates() []string
}

//Post builder creates posts and checks for errors.
type PostBuilder interface {
	Id(string) PostBuilder
	Permalink(string) PostBuilder
	Title(string) PostBuilder
	Slug(string) PostBuilder
	Content(string) PostBuilder
	State(string) PostBuilder
	PublicationDate(time.Time) PostBuilder

	Json(string) (*Post, error)
	BuildFromPost(untrustPost Post) (*Post, error)
	Build() (*Post, error)
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
	if id == "" {
		return "", &NilField{object: "Post", field: "Id"}
	}

	if !validIdRegexpTest.MatchString(id) {
		return "", &TypeError{object: "Post", expectedType: " must compile regex: " + validIdRegexpTest.String(), field: "Id"}
	}
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

//Sets the slug for the builder. Don't check for errors.
func (post *postBuilder) Slug(slug string) PostBuilder {
	post.slug = slug
	return post
}

//Test if the slug is valid. Returns an empty string and error if
//any errors are found.
func isValidSlug(slug string) (string, error) {
	return slug, nil
}

//Sets the content for the builder. Don't check for errors in the content.
func (post *postBuilder) Content(content string) PostBuilder {
	post.content = content
	return post
}

//Tests if the content is valid html/commonmark compilant. Returns an empty string and error if
//any errors are found.
func isValidContent(content string) (string, error) {
	// TODO Should test markdown compilant content.
	return content, nil
}

//Sets the state for the builder. Don't check for errors in the state passed.
func (post *postBuilder) State(state string) PostBuilder {
	post.state = state
	return post
}

//Test if a post state is valid. If not valid returs an empty state and a TypeError.
func isValidState(state string) (string, error) {
	if !Contains(state, postState) {
		return "", &TypeError{object: "Post", expectedType: "State", field: "State", value: state}
	}
	return state, nil
}

//Sets the publication date for the builder. Don't check for erros in the publication date.
func (post *postBuilder) PublicationDate(publicationDate time.Time) PostBuilder {
	post.publicationDate = publicationDate
	return post
}

//Test if a publication date is valid. Returns a valid publication date or if the
//publication date is not valid, returns nil and an error message.
func isValidPublicationDate(publicationDate time.Time) (time.Time, error) {
	return publicationDate, nil
}

//Creates a post form an untrusted source. Useful if you d not know if a post
// is bad contructed. All post from API using posts test their obejects here.
func (post *postBuilder) BuildFromPost(untrustPost Post) (*Post, error) {
	post.id = untrustPost.Id
	post.permalink = untrustPost.Permalink
	post.title = untrustPost.Title
	post.slug = untrustPost.Slug
	post.content = untrustPost.Content
	post.state = untrustPost.State
	post.publicationDate = untrustPost.PublicationDate
	return post.Build()
}

//Creates a new post from JSON data.
func (post *postBuilder) Json(jsonData string) (*Post, error) {
	var m map[string]interface{}
	var _temp interface{}

	err := json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		return nil, err
	}

	typeCheck := func(field, expected string, couldNotBeNil bool, data interface{}) (interface{}, error) {
		if data == nil && couldNotBeNil {
			return nil, &NilField{object: "Post", field: field}
		}
		switch data.(type) {
		case string:
			if expected == "string" {
				return data.(string), nil
			}
		case time.Time:
			if expected == "time.Time" {
				return data.(time.Time), nil
			}
		}
		return nil, &TypeError{object: "Post", expectedType: expected, field: field}
	}

	_temp, err = typeCheck(ID_FIELD, ID_FIELD_TYPE, true, m[ID_FIELD])
	if err != nil {
		return nil, err
	}
	post.id = _temp.(string)

	_temp, err = typeCheck(PERMALINK_FIELD, PERMALINK_FIELD_TYPE, true, m[PERMALINK_FIELD])
	if err != nil {
		return nil, err
	}
	post.permalink = _temp.(string)

	_temp, err = typeCheck(TITLE_FIELD, TITLE_FIELD_TYPE, true, m[TITLE_FIELD])
	if err != nil {
		return nil, err
	}
	post.title = _temp.(string)

	_temp, err = typeCheck(SLUG_FIELD, SLUG_FIELD_TYPE, true, m[SLUG_FIELD])
	if err != nil {
		return nil, err
	}
	post.slug = _temp.(string)

	_temp, err = typeCheck(CONTENT_FIELD, CONTENT_FIELD_TYPE, true, m[CONTENT_FIELD])
	if err != nil {
		return nil, err
	}
	post.content = _temp.(string)

	_temp, err = typeCheck(STATE_FIELD, STATE_FIELD_TYPE, true, m[STATE_FIELD])
	if err != nil {
		return nil, err
	}
	post.state = _temp.(string)

	// TODO Rewrite the checking.
	if m[PUBLICATIONDATE_FIELD] != nil {
		_temp, _ = time.Parse(time.RFC3339Nano, _temp.(string))
	}
	_temp, err = typeCheck(PUBLICATIONDATE_FIELD, PUBLICATIONDATE_FIELD_TYPE, false, _temp)
	if err != nil {
		return nil, err
	}
	post.publicationDate = _temp.(time.Time)

	return post.Build()
}

//Creates a post and test for config errors.
func (post *postBuilder) Build() (*Post, error) {
	var p Post
	var err error

	p.Id, err = isValidId(post.id)
	if err != nil {
		return nil, err
	}
	p.Permalink, err = isValidPermalink(post.permalink)
	if err != nil {
		return nil, err
	}
	p.Title, err = isValidTitle(post.title)
	if err != nil {
		return nil, err
	}
	p.Slug, err = isValidSlug(post.slug)
	if err != nil {
		return nil, err
	}
	p.Content, err = isValidContent(post.content)
	if err != nil {
		return nil, err
	}
	p.State, err = isValidState(post.state)
	if err != nil {
		return nil, err
	}
	p.PublicationDate, err = isValidPublicationDate(post.publicationDate)
	if err != nil {
		return nil, err
	}

	return &p, err
}

//Creates a post builder.
func NewPostBuilder() PostBuilder {
	return &postBuilder{}
}

//Creates a post. This call is insecure and do not check for erros in the parameters.
//The builder should be use instead.
func NewPost(permalink, title, content, state string, publicationDate time.Time) *Post {

	post, _ := NewPostBuilder().
		Id("24324").
		Permalink(permalink).
		Title(title).
		Slug(createSlugFromTitle(title)).
		Content(content).
		State(state).
		PublicationDate(publicationDate).
		Build()

	return post
}

//From a post to the equivalent in JSON format.
func (post *Post) PostToJson() (string, error) {
	var compactPost *bytes.Buffer
	postJson, err := json.Marshal(post)
	if err != nil {
		// TODO should panic ???!!!
	}
	json.Compact(compactPost, postJson)
	return compactPost.String(), err
}
