package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
	"web-analyzer/adapter"
	"web-analyzer/modals"
	"web-analyzer/services"
	"web-analyzer/utils"
	"web-analyzer/validators"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func LoadInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "input.html", nil)
}

func InvokeAnalyzer(c *gin.Context) {

	start := time.Now()

	formUrl := c.PostForm("url")
	utils.Log.Infof("Started analyzing. url: %s", formUrl)

	isValid := validators.IsValidURL(formUrl)
	if !isValid {

		utils.Log.Errorf("Invalid url: %s", formUrl)

		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusBadRequest,
			"ErrorMessage": "Invalid URL please check the URL and try again",
		})
		return
	}

	if !strings.HasPrefix(formUrl, "http") {
		formUrl = "https://" + formUrl
	}

	baseUrl, err := utils.GetBaseURL(formUrl)
	if err != nil {

		utils.Log.Errorf("Unable to parse URL please: %s", formUrl)

		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusBadRequest,
			"ErrorMessage": "Unable to parse URL please check the URL and try again",
		})
	}

	body, status, err := adapter.NewRequestInvoker().InvokeRequest(formUrl, "GET")

	if err != nil {

		utils.Log.Errorf("Failed to fetch URL: %s", formUrl)

		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"URL":          formUrl,
			"StatusCode":   http.StatusNotFound,
			"Error":        err,
			"ErrorMessage": fmt.Sprintf("Failed to fetch URL: %v", err),
		})
		return
	}

	if status != http.StatusOK {

		utils.Log.Errorf("Error in status code: %s", formUrl)

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

		utils.Log.Errorf("HTML parsing error: %s", formUrl)

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
		InputUrl: formUrl,
	}

	services.InvokeAnalyzers(ctx)

	pageInfoModal := ctx.Manager.GetPageInfoModal()

	defer func() {
		elapsed := time.Since(start).Seconds()
		utils.Log.Infof("Completed analysis. url: %s,  elapsed_time: %ds", formUrl, int(elapsed))
	}()

	c.HTML(http.StatusOK, "result.html", gin.H{
		"URL":           formUrl,
		"Title":         pageInfoModal.Title,
		"Version":       pageInfoModal.HtmlVersion,
		"Headings":      pageInfoModal.HeadingProperties,
		"Internal":      pageInfoModal.NoOfInternalLinks,
		"External":      pageInfoModal.NoOfExternalLinks,
		"Inaccessible":  pageInfoModal.NoOfInaccessibleLinks,
		"HasLogin":      pageInfoModal.HasLogin,
		"ExecutionTime": fmt.Sprintf("%2d seconds", int(time.Since(start).Seconds())),
	})
}
