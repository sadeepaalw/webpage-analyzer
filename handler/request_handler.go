package handler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"web-analyzer/adapter"
	"web-analyzer/modals"
	"web-analyzer/services"
	"web-analyzer/utils"
	"web-analyzer/validators"
)

func LoadInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", nil)
}

func InvokeAnalyzer(c *gin.Context) {

	formUrl := c.PostForm("url")

	isValid := validators.IsValidURL(formUrl)
	if !isValid {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusBadRequest,
			"ErrorMessage": "Invalid URL please check the URL and try again",
		})
		return
	}

	baseUrl, err := utils.GetBaseURL(formUrl)
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
	services.LoginFormAnalyzer().Analyze(ctx)
	services.HeadingAnalyzer().Analyze(ctx)
	services.HtmlVersionAnalyzer().Analyze(ctx)
	services.LinkAnalyzer().Analyze(ctx)

	//internal, external := 0, 0
	//doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
	//	href, _ := s.Attr("href")
	//	if strings.HasPrefix(href, "/") || strings.Contains(href, formUrl) {
	//		internal++
	//	} else {
	//		external++
	//	}
	//})

	pageInfoModal := ctx.Manager.GetPageInfoModal()
	fmt.Println("internal:", pageInfoModal)

	c.HTML(http.StatusOK, "result.html", gin.H{
		"URL":      formUrl,
		"Title":    pageInfoModal.Title,
		"Version":  pageInfoModal.HtmlVersion,
		"Headings": pageInfoModal.HeadingProperties,
		"Internal": pageInfoModal.NoOfInternalLinks,
		"External": pageInfoModal.NoOfExternalLinks,
		"HasLogin": pageInfoModal.HasLogin,
	})
}
