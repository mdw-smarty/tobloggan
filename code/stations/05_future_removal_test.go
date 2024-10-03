package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestFutureRemovalFixture(t *testing.T) {
	gunit.Run(new(FutureRemovalFixture), t)
}

type FutureRemovalFixture struct {
	StationFixture
	now time.Time
}

func (this *FutureRemovalFixture) Setup() {
	this.now = time.Now()
	this.station = NewFutureRemoval(this.now)
}

func (this *FutureRemovalFixture) TestPastArticleKept() {
	input := contracts.Article{Date: this.now.Add(-time.Nanosecond)}
	this.do(input)
	this.assertOutputs(input)
}
func (this *FutureRemovalFixture) TestCurrentArticleKept() {
	input := contracts.Article{Date: this.now}
	this.do(input)
	this.assertOutputs(input)
}
func (this *FutureRemovalFixture) TestFutureArticleDropped() {
	input := contracts.Article{Date: this.now.Add(time.Nanosecond)}
	this.do(input)
	this.So(this.outputs, should.BeEmpty)
}
