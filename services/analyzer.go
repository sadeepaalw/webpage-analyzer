package services

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"web-analyzer/modals"
)

type Analyzer interface {
	Analyze(ctx AnalyzerContext)
}

type AnalyzerContext struct {
	Document *goquery.Document
	Manager  *modals.PageInfoModalManager
	BaseURL  *url.URL
}
