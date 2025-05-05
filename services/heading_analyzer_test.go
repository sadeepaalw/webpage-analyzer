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

func TestHeadingAnalyzer_Analyze(t *testing.T) {

	html, err := os.ReadFile("../resources/HTML5.html")
	require.NoError(t, err)

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	require.NoError(t, err)

	modalManager := modals.NewPageInfoModalManager()

	expectedHeadingProperties := []modals.Property{
		{
			PropertyName:        "h1",
			NumberOfOccurrences: 1,
		}, {
			PropertyName:        "h2",
			NumberOfOccurrences: 2,
		},
		{
			PropertyName:        "h3",
			NumberOfOccurrences: 1,
		},
		{
			PropertyName:        "h4",
			NumberOfOccurrences: 1,
		},
		{
			PropertyName:        "h5",
			NumberOfOccurrences: 1,
		},
		{
			PropertyName:        "h6",
			NumberOfOccurrences: 1,
		},
	}

	ctx := AnalyzerContext{
		Document: doc,
		Manager:  modalManager,
		BaseURL:  &url.URL{},
	}

	analyzer := HeadingAnalyzer()
	analyzer.Analyze(ctx)

	assert.Equal(t, expectedHeadingProperties, modalManager.GetPageInfoModal().HeadingProperties)
}
