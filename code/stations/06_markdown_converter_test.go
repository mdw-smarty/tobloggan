package stations

import (
	"testing"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestMarkdownConverterFixture(t *testing.T) {
	gunit.Run(new(MarkdownConverterFixture), t)
}

type MarkdownConverterFixture struct {
	StationFixture
	markdownErr error
}

func (this *MarkdownConverterFixture) Convert(content string) (string, error) {
	return content + " CONVERTED", this.markdownErr
}

func (this *MarkdownConverterFixture) Setup() {
	this.station = NewMarkdownConverter(this)
}

func (this *MarkdownConverterFixture) TestBodyConverted() {
	input := contracts.Article{Body: article1Content}
	this.do(input)
	this.assertOutputs(
		contracts.Article{Body: article1Content + " CONVERTED"},
	)
}
func (this *MarkdownConverterFixture) TestInvalidMarkdown() {
	this.markdownErr = boink
	this.do(contracts.Article{Body: article1Content})
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errMalformedContent)
	}
}
