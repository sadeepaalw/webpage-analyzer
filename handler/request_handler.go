package handler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web-analyzer/adapter"
)

func InvokeInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", nil)
}

func InvokeAnalyzer(c *gin.Context) {

	url := c.PostForm("url")
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	//http.Get(url)

	resp, err := adapter.InvokeRequest(url, "GET")
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          url,
			"StatusCode":   http.StatusBadRequest,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Failed to fetch URL: %v", err),
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          url,
			"StatusCode":   resp.StatusCode,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Non-200 status code: %d", resp.StatusCode),
		})
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          url,
			"StatusCode":   resp.StatusCode,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Error parsing HTML: %v", err),
		})
		return
	}

	title := doc.Find("title").Text()
	headings := map[string]int{}
	for i := 1; i <= 6; i++ {
		headings[fmt.Sprintf("h%d", i)] = doc.Find(fmt.Sprintf("h%d", i)).Length()
	}
	hasLogin := doc.Find("input[type='password']").Length() > 0

	internal, external := 0, 0
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, "/") || strings.Contains(href, url) {
			internal++
		} else {
			external++
		}
	})

	c.HTML(http.StatusOK, "result.html", gin.H{
		"URL":      url,
		"Title":    title,
		"Headings": headings,
		"Internal": internal,
		"External": external,
		"HasLogin": hasLogin,
	})
}
