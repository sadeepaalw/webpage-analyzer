package services

import (
	"github.com/PuerkitoBio/goquery"
	"web-analyzer/modals"
)

type loginFormAnalyzer struct{}

func (l loginFormAnalyzer) Analyze(doc *goquery.Document, manager *modals.PageInfoModalManager) {

	hasLogin := doc.Find("input[type='password']").Length() > 0
	manager.SetHasLogin(hasLogin)
}

func NewLoginFormAnalyzer() Analyzer {
	return loginFormAnalyzer{}
}
