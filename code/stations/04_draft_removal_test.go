package stations

import (
	"testing"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestDraftRemovalFixture(t *testing.T) {
	gunit.Run(new(DraftRemovalFixture), t)
}

type DraftRemovalFixture struct {
	StationFixture
}

func (this *DraftRemovalFixture) Setup() {
	this.station = NewDraftRemoval()
}

func (this *DraftRemovalFixture) TestDraftDropped() {
	input := contracts.Article{Draft: true}
	this.do(input)
	this.So(this.outputs, should.BeEmpty)
}
func (this *DraftRemovalFixture) TestNonDraftRetained() {
	input := contracts.Article{Draft: false}
	this.do(input)
	this.assertOutputs(input)
}
