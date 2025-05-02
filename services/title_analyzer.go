package services

import (
	"github.com/PuerkitoBio/goquery"
	"web-analyzer/modals"
)

type titleAnalyzer struct{}

func (t titleAnalyzer) Analyze(doc *goquery.Document, manager *modals.PageInfoModalManager) {
	manager.SetTitle(doc.Find("title").Text())
}

func TitleAnalyzer() Analyzer {
	return &titleAnalyzer{}
}
