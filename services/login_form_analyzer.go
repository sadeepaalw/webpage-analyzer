package services

type loginFormAnalyzer struct{}

func (l *loginFormAnalyzer) Analyze(ctx AnalyzerContext) {
	hasLogin := ctx.Document.Find("input[type='password']").Length() > 0
	ctx.Manager.SetHasLogin(hasLogin)
}

func LoginFormAnalyzer() Analyzer {
	return &loginFormAnalyzer{}
}
