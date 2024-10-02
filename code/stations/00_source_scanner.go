package stations

import (
	"io/fs"
	"path/filepath"

	"tobloggan/code/contracts"
)

type SourceScanner struct {
	fs fs.FS
}

func NewSourceScanner(fs fs.FS) contracts.Station {
	return &SourceScanner{fs: fs}
}

func (this *SourceScanner) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceDirectory:
		err := fs.WalkDir(this.fs, string(input), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return contracts.Error(err)
			}
			if d.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".md" {
				return nil
			}
			output(contracts.SourceFilePath(path))
			return nil
		})
		if err != nil {
			output(err)
		}
	default:
		output(input)
	}
}
