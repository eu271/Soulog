package objectsTestUtil

import (
	"github.com/eu271/Soulog/Blog/objects"
)

func NewTestPost() *soul.Post {
	var post *soul.Post
	var err error

	post, err = soul.NewPostBuilder().
		Id("asdasd").
		Permalink("asdasdasd").
		Title("asdasda").
		Slug("asasd").
		Content("asdasdasdasd").
		State("publish").
		Build()

	if err != nil {
		panic(err.Error())
	}
	if post == nil {
		panic("Post es nil")
	}

	return post

}
