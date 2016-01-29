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
	"time"
)

type User struct {
	Id        string    `json: "Id"`
	Name      string    `json: "Name"`
	Email     string    `json: "Email"`
	Slug      string    `json: "Slug"`
	Image     string    `json: "Image"`
	Bio       string    `json: "Bio"`
	LastLogin time.Time `json: "LastLogin"`
	CreatedAt time.Time `json: "Created"`
	UpdatedAt time.Time `json: "Update"`
}

type UserBuilder interface {
	Id(string) UserBuilder
	Name(string) UserBuilder
	Email(string) UserBuilder
	Slug(string) UserBuilder
	Image(string) UserBuilder
	Bio(string) UserBuilder
	LastLogin(time.Time) UserBuilder
	CreatedAt(time.Time) UserBuilder
	UpdatedAt(time.Time) UserBuilder

	Build() (User, error)
}

type userBuilder struct {
	id        string
	name      string
	email     string
	slug      string
	image     string
	bio       string
	lastLogin time.Time
	createdAt time.Time
	updatedAt time.Time
}

func (user *userBuilder) Id(id string) UserBuilder {
	user.id = id
	return user
}
func (user *userBuilder) Name(name string) UserBuilder {
	user.name = name
	return user
}
func (user *userBuilder) Email(email string) UserBuilder {
	user.email = email
	return user
}
func (user *userBuilder) Slug(slug string) UserBuilder {
	user.slug = slug
	return user
}
func (user *userBuilder) Image(image string) UserBuilder {
	user.image = image
	return user
}
func (user *userBuilder) Bio(bio string) UserBuilder {
	user.bio = bio
	return user
}
func (user *userBuilder) LastLogin(lastLogin time.Time) UserBuilder {
	user.lastLogin = lastLogin
	return user
}
func (user *userBuilder) CreatedAt(createdAt time.Time) UserBuilder {
	user.createdAt = createdAt
	return user
}
func (user *userBuilder) UpdatedAt(updatedAt time.Time) UserBuilder {
	user.updatedAt = updatedAt
	return user
}

func (user *userBuilder) Build() (User, error) {
	u := User{
		Id:        user.id,
		Name:      user.name,
		Email:     user.email,
		Slug:      user.slug,
		Image:     user.image,
		Bio:       user.bio,
		LastLogin: user.lastLogin,
		CreatedAt: user.createdAt,
		UpdatedAt: user.updatedAt,
	}
	return u, nil
}

func NewUserBuilder() UserBuilder {
	return &userBuilder{}
}
