package objectsTestUtil

import (
	"github.com/eu271/Soulog/Blog/objects"
	"testing"
)

type TestPostBuilder interface {
	soul.PostBuilder
	WithParagraphs(int) TestPostBuilder
	WithRandomTitle() TestPostBuilder
	Before(*soul.Post) TestPostBuilder
	After(*soul.Post) TestPostBuilder
	FillWithRandom() TestPostBuilder
}

func NewTestPostBuilder() *TestPostBuilder {
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

	return &TestPostBuilder{}

}

func AssertEquals(post, post1 *soul.Post, t *testing.T) {

	if post.Id != post1.Id {
		t.Error("Posts ids not equals: " + post.Id + " " + post1.Id)
	}

	if post.Permalink != post1.Permalink {
		t.Error("Posts permalins not equals: " + post.Permalink + " " + post1.Permalink)
	}

}
