package services

import (
	"net/url"
	"web-analyzer/modals"

	"github.com/PuerkitoBio/goquery"
)

type Analyzer interface {
	Analyze(ctx AnalyzerContext)
}

type AnalyzerContext struct {
	Document *goquery.Document
	Manager  *modals.PageInfoModalManager
	BaseURL  *url.URL
	InputUrl string
}
