package stations

import (
	"os"
	"testing"
	"testing/fstest"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleWriterFixture(t *testing.T) {
	gunit.Run(new(ArticleWriterFixture), t)
}

type ArticleWriterFixture struct {
	StationFixture
	fs           fstest.MapFS
	writeFileErr error
	mkdirAllErr  error
}

func (this *ArticleWriterFixture) MkdirAll(path string, perm os.FileMode) error {
	this.fs[path] = &fstest.MapFile{Mode: perm}
	return this.mkdirAllErr
}
func (this *ArticleWriterFixture) WriteFile(filename string, data []byte, perm os.FileMode) error {
	this.fs[filename] = &fstest.MapFile{Data: data, Mode: perm}
	return this.writeFileErr
}
func (this *ArticleWriterFixture) Setup() {
	this.fs = fstest.MapFS{}
	this.station = NewPageWriter("target/directory", this)
}
func (this *ArticleWriterFixture) TestMkdirErr() {
	this.mkdirAllErr = boink
	input := contracts.Page{Path: "the/slug", Content: "b"}
	this.do(input)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, boink)
	}
	this.So(this.fs, should.NotContainKey, "target/directory/the/slug/index.html")
}
func (this *ArticleWriterFixture) TestWriteFileErr() {
	this.writeFileErr = boink
	input := contracts.Page{Path: "the/slug", Content: "b"}
	this.do(input)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, boink)
	}
}
func (this *ArticleWriterFixture) TestHappyPath() {
	input := contracts.Page{Path: "/the/slug", Content: "The Body"}

	this.do(input)

	this.So(this.outputs, should.Equal, []any{input})
	content, err := this.fs.ReadFile("target/directory/the/slug/index.html")
	this.So(err, should.BeNil)
	this.So(string(content), should.Equal, "The Body")
}
