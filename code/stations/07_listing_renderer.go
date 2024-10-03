package stations

import (
	"fmt"
	"sort"
	"strings"

	"tobloggan/code/contracts"
)

type ListingRenderer struct {
	articles []contracts.Article
	template string
}

func NewListingRenderer(template string) contracts.Station {
	return &ListingRenderer{template: template}
}

func (this *ListingRenderer) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		this.articles = append(this.articles, input)
	}
	output(input)
}
func (this *ListingRenderer) Finalize(output func(any)) {
	sort.SliceStable(this.articles, func(i, j int) bool {
		return this.articles[i].Date.Before(this.articles[j].Date)
	})
	var builder strings.Builder
	for _, article := range this.articles {
		builder.WriteString("\t\t\t")
		_, _ = fmt.Fprintf(&builder, `<li><a href="%s">%s</a></li>`, article.Slug, article.Title)
		builder.WriteString("\n")
	}
	replacer := strings.NewReplacer("{{Listing}}", builder.String())
	output(contracts.Page{Path: "/", Content: replacer.Replace(this.template)})
}
