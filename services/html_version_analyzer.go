package services

import (
	"golang.org/x/net/html"
)

type HtmlVersionAnalyzer struct{}

func (a HtmlVersionAnalyzer) Analyze(ctx AnalyzerContext) {

	//version := detectHTMLVersion(&ctx.Document.)
	//ctx.Manager.SetHtmlVersion(version)
}

func detectHTMLVersion(doc *html.Node) string {
	for node := doc; node != nil; node = node.NextSibling {
		if node.Type == html.DoctypeNode {
			switch node.Data {
			case "html":
				return "HTML5"
			case "HTML 4.01":
				return "HTML 4.01"
			case "XHTML 1.0":
				return "XHTML 1.0"
			default:
				return node.Data
			}
		}
	}
	return "Unknown"
}

func NewHtmlVersionAnalyzer() Analyzer {
	return &HtmlVersionAnalyzer{}
}
