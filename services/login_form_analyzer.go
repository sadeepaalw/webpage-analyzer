package services

import "strings"

type loginFormAnalyzer struct{}

func (l *loginFormAnalyzer) Analyze(ctx AnalyzerContext) {

	if strings.Contains(strings.ToLower(ctx.InputUrl), "login") {

		ctx.Manager.SetHasLogin(true)
		return
	}

	hasLogin := ctx.Document.Find("input[type='password']").Length() > 0
	ctx.Manager.SetHasLogin(hasLogin)
}

func LoginFormAnalyzer() Analyzer {
	return &loginFormAnalyzer{}
}
