package stations

import (
	"path/filepath"

	"github.com/mdwhatcott/tobloggan/code/contracts"
)

type FileSystemWriter interface {
	contracts.MkdirAll
	contracts.WriteFile
}
type PageWriter struct {
	targetDirectory string
	fs              FileSystemWriter
}

func NewPageWriter(targetDirectory string, fs FileSystemWriter) *PageWriter {
	return &PageWriter{
		targetDirectory: targetDirectory,
		fs:              fs,
	}
}
func (this *PageWriter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Page:
		path := filepath.Join(this.targetDirectory, input.Path, "index.html")
		err := this.fs.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			output(contracts.Error(err))
			return
		}
		err = this.fs.WriteFile(path, []byte(input.Content), 0644)
		if err != nil {
			output(contracts.Error(err))
			return
		}
		output(input)
	default:
		output(input)
	}
}