package services

import (
	"bytes"
	"net/url"
	"os"
	"testing"
	"web-analyzer/modals"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
)

func TestLinkAnalyzer_Analyze(t *testing.T) {

	html, err := os.ReadFile("../resources/HTML5.html")
	require.NoError(t, err)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	require.NoError(t, err)

	modalManager := modals.NewPageInfoModalManager()

	ctx := AnalyzerContext{
		Document: doc,
		Manager:  modalManager,
		BaseURL:  &url.URL{},
	}

	analyzer := LinkAnalyzer()
	analyzer.Analyze(ctx)

	assert.Equal(t, 1, modalManager.GetPageInfoModal().NoOfInternalLinks)
	assert.Equal(t, 1, modalManager.GetPageInfoModal().NoOfExternalLinks)
	assert.Equal(t, 1, modalManager.GetPageInfoModal().NoOfInaccessibleLinks)
}
