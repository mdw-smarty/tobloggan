package stations

import (
	"os"
	"testing"
	"testing/fstest"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSourceReaderFixture(t *testing.T) {
	gunit.Run(new(SourceReaderFixture), t)
}

type SourceReaderFixture struct {
	StationFixture
	fs fstest.MapFS
}

func (this *SourceReaderFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.station = NewSourceReader(this.fs)
}

func (this *SourceReaderFixture) TestSourceFileContentReadFromDiskAndEmitted() {
	content := "article 1 content"
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte(content)}
	this.do(contracts.SourceFilePath("src/article-1.md"))
	this.assertOutputs(contracts.SourceFile(content))
}
func (this *SourceReaderFixture) TestIOError() {
	delete(this.fs, "src/article-1.md")
	this.do(contracts.SourceFilePath("src/article-1.md"))
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, os.ErrNotExist)
	}
}
