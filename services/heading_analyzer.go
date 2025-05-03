package services

import (
	"fmt"
)

type headingAnalyzer struct{}

func (h *headingAnalyzer) Analyze(ctx AnalyzerContext) {
	for i := 1; i <= 6; i++ {
		ctx.Manager.SetHeadingProperties(fmt.Sprintf("h%d", i), ctx.Document.Find(fmt.Sprintf("h%d", i)).Length())
	}
}

func HeadingAnalyzer() Analyzer {
	return &headingAnalyzer{}
}
