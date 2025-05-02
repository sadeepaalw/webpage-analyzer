package services

type titleAnalyzer struct{}

func (t titleAnalyzer) Analyze(ctx AnalyzerContext) {
	ctx.Manager.SetTitle(ctx.Document.Find("title").Text())
}

func TitleAnalyzer() Analyzer {
	return &titleAnalyzer{}
}
