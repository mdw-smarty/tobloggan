package stations

import (
	"strings"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"

	"tobloggan/code/contracts"
	"tobloggan/code/html"
)

func TestListingRendererFixture(t *testing.T) {
	gunit.Run(new(ListingRendererFixture), t)
}

type ListingRendererFixture struct {
	StationFixture
}

func (this *ListingRendererFixture) Setup() {
	this.station = NewListingRenderer(html.ListingTemplate)
}

func (this *ListingRendererFixture) TestArticlesWrittenToListing() {
	article1 := contracts.Article{Slug: "s1", Title: "t1", Date: date("2024-09-01"), Body: "b1"}
	article2 := contracts.Article{Slug: "s2", Title: "t2", Date: date("2024-09-02"), Body: "b2"}
	article3 := contracts.Article{Slug: "s3", Title: "t3", Date: date("2024-09-03"), Body: "b3"}
	this.do(article1)
	this.do(article2)
	this.do(article3)

	this.finalize()

	this.So(this.outputs[:3], should.Equal, []any{
		article1,
		article2,
		article3,
	})
	page := this.outputs[3].(contracts.Page)
	content := page.Content
	this.So(page.Path, should.Equal, "/")
	this.So(content, should.ContainSubstring, `href="s1"`)
	this.So(content, should.ContainSubstring, `href="s2"`)
	this.So(content, should.ContainSubstring, `href="s3"`)
	this.So(content, should.ContainSubstring, `>t1<`)
	this.So(content, should.ContainSubstring, `>t2<`)
	this.So(content, should.ContainSubstring, `>t3<`)
	d1 := strings.Index(content, ">t1<")
	d2 := strings.Index(content, ">t2<")
	d3 := strings.Index(content, ">t3<")
	this.So(d1, should.BeLessThan, d2)
	this.So(d2, should.BeLessThan, d3)
}
