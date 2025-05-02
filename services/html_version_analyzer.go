package services

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type htmlVersionAnalyzer struct{}

func (a htmlVersionAnalyzer) Analyze(ctx AnalyzerContext) {

	var buf bytes.Buffer
	err := html.Render(&buf, ctx.Document.Nodes[0])
	if err != nil {
		return
	}
	htmlBytes := buf.Bytes()

	content := strings.ToLower(string(htmlBytes))
	content = strings.TrimSpace(content)

	version := detectHTMLVersion(content)
	fmt.Println(version)
	ctx.Manager.SetHtmlVersion(version)
}

func detectHTMLVersion(content string) string {
	if strings.HasPrefix(content, "<!doctype html>") {
		return "HTML5"
	}

	if strings.Contains(content, "xhtml 1.0 strict") {
		return "XHTML 1.0 Strict"
	} else if strings.Contains(content, "xhtml 1.0 transitional") {
		return "XHTML 1.0 Transitional"
	} else if strings.Contains(content, "xhtml 1.0 frameset") {
		return "XHTML 1.0 Frameset"
	} else if strings.Contains(content, "html 4.01 strict") {
		return "HTML 4.01 Strict"
	} else if strings.Contains(content, "html 4.01 transitional") {
		return "HTML 4.01 Transitional"
	} else if strings.Contains(content, "html 4.01 frameset") {
		return "HTML 4.01 Frameset"
	}

	return "Unknown or Missing DOCTYPE"
}

func HtmlVersionAnalyzer() Analyzer {
	return &htmlVersionAnalyzer{}
}
