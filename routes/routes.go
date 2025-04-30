package routes

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRoutes(r *gin.Engine) {
	//r.SetFuncMap(template.FuncMap{"safe": func(s string) template.HTML { return template.HTML(s) }})
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "input.html", nil)
	})

	r.POST("/analyze", func(c *gin.Context) {
		url := c.PostForm("url")
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to fetch URL: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.String(resp.StatusCode, "Non-200 status code: %d", resp.StatusCode)
			return
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing HTML: %v", err)
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
	})
}
