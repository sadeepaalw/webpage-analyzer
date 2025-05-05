package services

import (
	"sync"
	"time"
	"web-analyzer/utils"
)

type AnalyzerWrapper struct {
	Name    string
	Analyze func(ctx AnalyzerContext)
}

func InvokeAnalyzers(ctx AnalyzerContext) {

	var wg sync.WaitGroup

	analyzers := []AnalyzerWrapper{
		{Name: "TitleAnalyzer", Analyze: TitleAnalyzer().Analyze},
		{Name: "LoginFormAnalyzer", Analyze: LoginFormAnalyzer().Analyze},
		{Name: "HeadingAnalyzer", Analyze: HeadingAnalyzer().Analyze},
		{Name: "HtmlVersionAnalyzer", Analyze: HtmlVersionAnalyzer().Analyze},
		{Name: "LinkAnalyzer", Analyze: LinkAnalyzer().Analyze},
	}

	for _, analyzer := range analyzers {
		wg.Add(1)

		go func(a AnalyzerWrapper) {
			defer wg.Done()
			start := time.Now()
			utils.Log.Infof("Starting execution of %s", a.Name)
			a.Analyze(ctx)
			utils.Log.Infof("Completed execution of %s, Elapsed Time: %ds", a.Name, int(time.Since(start).Seconds()))
		}(analyzer)
	}
	wg.Wait()

}
