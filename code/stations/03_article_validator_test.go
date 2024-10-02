package stations

import (
	"strings"
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleValidatorFixture(t *testing.T) {
	gunit.Run(new(ArticleValidatorFixture), t)
}

type ArticleValidatorFixture struct {
	StationFixture
}

func (this *ArticleValidatorFixture) Setup() {
	this.station = NewArticleValidator()
}
func (test *ArticleValidatorFixture) validArticle() contracts.Article {
	return contracts.Article{
		Draft: false,
		Slug:  "the/slug",
		Title: "The Title",
		Date:  time.Now(),
		Body:  "The body.",
	}
}
func (this *ArticleValidatorFixture) assertInvalidSlug(invalidSlug string) {
	this.outputs = nil
	input := this.validArticle()
	input.Slug = invalidSlug
	this.do(input)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errInvalidContent)
		this.So(this.outputs[0].(error).Error(), should.ContainSubstring, "slug")
	}
}
func (this *ArticleValidatorFixture) assertInvalidTitle(invalidTitle string) {
	input := this.validArticle()
	input.Title = invalidTitle
	this.outputs = nil
	this.do(input)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errInvalidContent)
		this.So(this.outputs[0].(error).Error(), should.ContainSubstring, "title")
	}
}
func (this *ArticleValidatorFixture) TestValidArticle() {
	input := this.validArticle()
	this.do(input)
	this.assertOutputs(input)
}
func (this *ArticleValidatorFixture) TestInvalidSlugs() {
	this.assertInvalidSlug("")
	this.assertInvalidSlug("slug with spaces")
	this.assertInvalidSlug("SLUG/WITH/ALL-CAPS")
	this.assertInvalidSlug("consecutive//slashes")
	this.assertInvalidSlug(strings.Repeat("a", 129))
}
func (this *ArticleValidatorFixture) TestInvalidTitles() {
	this.assertInvalidTitle("")                       // empty
	this.assertInvalidTitle(strings.Repeat("a", 257)) // too long
}
func (this *ArticleValidatorFixture) TestSlugsMustBeUnique() {
	input := this.validArticle()
	this.do(input)
	this.do(input)
	if this.So(this.outputs, should.HaveLength, 2) {
		this.So(this.outputs[0], should.Equal, input)
		this.So(this.outputs[1], should.Wrap, errDuplicateSlug)
	}
}
