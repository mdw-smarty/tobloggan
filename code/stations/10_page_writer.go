package stations

import (
	"path/filepath"

	"tobloggan/code/contracts"
)

type PageWriter struct {
	targetDirectory string
	fs              contracts.FSWriter
}

func NewPageWriter(targetDirectory string, fs contracts.FSWriter) contracts.Station {
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
