package handler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
	"web-analyzer/adapter"
	"web-analyzer/modals"
	"web-analyzer/services"
	"web-analyzer/validators"
)

func GetBaseURL(rawUrl string) (*url.URL, error) {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	parsedURL.Path = ""
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""
	return parsedURL, nil
}

func InvokeInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", nil)
}

func InvokeAnalyzer(c *gin.Context) {

	formUrl := c.PostForm("formUrl")

	isValid := validators.IsValidURL(formUrl)
	if !isValid {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusBadRequest,
			"ErrorMessage": "Invalid URL please check the URL and try again",
		})
		return
	}

	baseUrl, err := GetBaseURL(formUrl)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusBadRequest,
			"ErrorMessage": "Unable to parse URL please check the URL and try again",
		})
	}

	if !strings.HasPrefix(formUrl, "http") {
		formUrl = "https://" + formUrl
	}

	body, status, err := adapter.InvokeRequest(formUrl, "GET")
	if err != nil {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusNotFound,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Failed to fetch URL: %v", err),
		})
		return
	}

	if status != http.StatusOK {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   status,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Non-200 status code: %d", status),
		})
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   status,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Error parsing HTML: %v", err),
		})
		return
	}

	ctx := services.AnalyzerContext{
		Document: doc,
		Manager:  modals.NewPageInfoModalManager(),
		BaseURL:  baseUrl,
	}

	services.TitleAnalyzer().Analyze(ctx)
	services.NewLoginFormAnalyzer().Analyze(ctx)
	services.NewHeadingAnalyzer().Analyze(ctx)

	internal, external := 0, 0
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if strings.HasPrefix(href, "/") || strings.Contains(href, formUrl) {
			internal++
		} else {
			external++
		}
	})

	fmt.Println("internal:", ctx.Manager.GetPageInfoModal())

	c.HTML(http.StatusOK, "result.html", gin.H{
		"URL":      formUrl,
		"Title":    ctx.Manager.GetPageInfoModal().Title,
		"Headings": ctx.Manager.GetPageInfoModal().HeadingProperties,
		"Internal": internal,
		"External": external,
		"HasLogin": ctx.Manager.GetPageInfoModal().HasLogin,
	})
}
