package services

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
)

type linkAnalyzer struct{}

func (l linkAnalyzer) Analyze(ctx AnalyzerContext) {

	internal, external, inaccessible := 0, 0, 0
	ctx.Document.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" {
			return
		}

		if strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "tel:") ||
			strings.HasPrefix(href, "javascript:") {
			return
		}

		linkURL, err := url.Parse(href)
		if err != nil {
			return
		}

		resolvedURL := ctx.BaseURL.ResolveReference(linkURL)

		if resolvedURL.Host == ctx.BaseURL.Host {
			internal++
		} else {
			external++
		}
		
	})

	ctx.Manager.SetNoOfInternalLinks(internal)
	ctx.Manager.SetNoOfExternalLinks(external)
	ctx.Manager.SetNoOfInaccessibleLinks(inaccessible)

}

func LinkAnalyzer() Analyzer {
	return &linkAnalyzer{}
}
