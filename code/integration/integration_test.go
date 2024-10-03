package integration

import (
	"bytes"
	"io/fs"
	"log"
	"os"
	"sync"
	"testing"
	"testing/fstest"
	"time"

	"tobloggan/code/html"
	"tobloggan/code/markdown"

	"github.com/smarty/assertions/should"
)

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in -short mode")
	}

	var logBuffer bytes.Buffer
	fileSystem := NewFileSystem()
	_ = fileSystem.WriteFile("article-1.md", []byte(article1Content), 0644)
	_ = fileSystem.WriteFile("article-2.md", []byte(article2Content), 0644)
	config := Config{
		Clock:             time.Now,
		Logger:            log.New(&logBuffer, "[TEST] ", 0),
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  fileSystem,
		FileSystemWriter:  fileSystem,
		TargetDirectory:   "output",
		ArticleTemplate:   html.ArticleTemplate,
		ListingTemplate:   html.ListingTemplate,
		BaseURL:           "file://",
	}
	ok := GenerateBlog(config)
	if !ok {
		t.Error("failed to generate blog")
	}

	t.Log("\n" + logBuffer.String())

	listing, _ := fs.ReadFile(fileSystem.MapFS, "output/index.html")
	article1, _ := fs.ReadFile(fileSystem.MapFS, "output/article/1/index.html")
	article2, _ := fs.ReadFile(fileSystem.MapFS, "output/article/2/index.html")

	should.So(t, string(listing), should.ContainSubstring, `<li><a href="file:///article/1">Article 1</a></li>`)
	should.So(t, string(listing), should.ContainSubstring, `<li><a href="file:///article/2">Article 2</a></li>`)
	should.So(t, string(article1), should.ContainSubstring, `<p>The contents of article 1.</p>`)
	should.So(t, string(article2), should.ContainSubstring, `<p>The contents of article 2.</p>`)
}

type FileSystem struct {
	lock *sync.Mutex
	fstest.MapFS
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		lock:  new(sync.Mutex),
		MapFS: make(fstest.MapFS),
	}
}
func (this FileSystem) ReadFile(name string) ([]byte, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.MapFS.ReadFile(name)
}
func (this FileSystem) MkdirAll(path string, perm os.FileMode) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.MapFS[path] = &fstest.MapFile{Mode: perm}
	return nil
}
func (this FileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.MapFS[filename] = &fstest.MapFile{Data: data, Mode: perm}
	return nil
}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`

const article2Content = `{
	"date": "2024-09-05T00:00:00Z",
	"slug": "/article/2",
	"title": "Article 2"
}

+++

The contents of article 2.`
