package services

import (
	"github.com/PuerkitoBio/goquery"
	"web-analyzer/modals"
)

type Analyzer interface {
	Analyze(doc *goquery.Document, manager *modals.PageInfoModalManager)
}
