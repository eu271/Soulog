package objectsTestUtil

import (
	"github.com/eu271/Soulog/Blog/objects"
)

func NewTestPost() soul.Post {
	var post soul.Post

	post, _ = soul.NewPostBuilder().
		Id("asdasd").
		Permalink("asdasdasd").
		Title("asdasda").
		Slug("asasd").
		Content("asdasdasdasd").
		State("publish").
		Build()
	return post

}
