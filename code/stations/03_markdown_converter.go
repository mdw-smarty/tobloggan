package stations

import (
	"github.com/mdwhatcott/pipelines"
	"github.com/mdwhatcott/tobloggan/code/contracts"
)

type Markdown interface {
	Convert(content string) (string, error)
}

type MarkdownConverter struct {
	md Markdown
}

func NewMarkdownConverter(md Markdown) pipelines.Station {
	return &MarkdownConverter{md: md}
}

func (this *MarkdownConverter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		converted, err := this.md.Convert(input.Body)
		if err != nil {
			output(contracts.Errorf("%w (%w): %s", errMalformedSource, err, input))
		} else {
			input.Body = converted
			output(input)
		}
	default:
		output(input)
	}
}