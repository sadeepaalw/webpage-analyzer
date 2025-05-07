package services

import (
	"net/url"
	"strings"
	"sync"
	"web-analyzer/adapter"
	"web-analyzer/utils"

	"github.com/PuerkitoBio/goquery"
)

type linkAnalyzer struct{}

func (l *linkAnalyzer) Analyze(ctx AnalyzerContext) {

	internal, external, inaccessible := 0, 0, 0

	var inaccessibleLinks []string

	var wg sync.WaitGroup
	var mu sync.Mutex

	channel := make(chan struct{}, 10)

	ctx.Document.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" {
			return
		}

		if strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "tel:") ||
			strings.HasPrefix(href, "javascript:") {
			return
		}

		linkURL, err := url.Parse(href)
		if err != nil {
			return
		}

		resolvedURL := ctx.BaseURL.ResolveReference(linkURL)

		wg.Add(1)
		channel <- struct{}{}

		if resolvedURL.Host == ctx.BaseURL.Host {
			internal++
		} else {
			external++
		}
		go checkInaccessiblity(resolvedURL.String(), &wg, &mu, &inaccessible, &inaccessibleLinks, channel)

	})

	wg.Wait()

	ctx.Manager.SetNoOfInternalLinks(internal)
	ctx.Manager.SetNoOfExternalLinks(external)
	ctx.Manager.SetNoOfInaccessibleLinks(inaccessible)

}

func checkInaccessiblity(url string, wg *sync.WaitGroup, mu *sync.Mutex, inaccessible *int, inaccessibleLinks *[]string, sem chan struct{}) {

	defer wg.Done()
	defer func() { <-sem }()

	isInaccessible := false

	_, statusCode, err := adapter.NewRequestInvoker().InvokeRequest(url, "HEAD")

	if (err != nil || statusCode == 0 || statusCode == 503) && statusCode != 429 {
		utils.Log.Infof("statusCode : %v\n", statusCode)
		utils.Log.Infof("err : %v\n", err)
		utils.Log.Infof("Inaccess HEAD: %v\n", url)
		isInaccessible = true

	} else {
		_, statusCode, err = adapter.NewRequestInvoker().InvokeRequest(url, "GET")

		if (err != nil || statusCode == 0 || statusCode == 503) && statusCode != 429 {

			utils.Log.Infof("statusCode : %v\n", statusCode)
			utils.Log.Infof("err : %v\n", err)
			utils.Log.Infof("Inaccess GET: %v\n", url)
			isInaccessible = true
		}
	}

	if isInaccessible {
		mu.Lock()
		*inaccessible++
		*inaccessibleLinks = append(*inaccessibleLinks, url)
		mu.Unlock()
	}

}

func LinkAnalyzer() Analyzer {
	return &linkAnalyzer{}
}
