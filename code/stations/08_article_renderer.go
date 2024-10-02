package stations

import (
	"strings"

	"tobloggan/code/contracts"
)

type ArticleRenderer struct {
	template string
}

func NewArticleRenderer(template string) contracts.Station {
	return &ArticleRenderer{template: template}
}
func (this *ArticleRenderer) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		replacer := strings.NewReplacer(
			"{{Title}}", input.Title,
			"{{Slug}}", input.Slug,
			"{{Date}}", input.Date.Format("January 2, 2006"),
			"{{Body}}", input.Body,
		)
		output(contracts.Page{
			Path:    input.Slug,
			Content: replacer.Replace(this.template),
		})
	default:
		output(input)
	}
}
