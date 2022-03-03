package handler

import (
	"crawler/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"net/http"
	"sync"
)

func (h *Handler) getTitlesByUrls(c *gin.Context) {
	var input model.Urls
	var output = NewResponse()

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"urls (array) parameter required"})
		return
	}

	var wg sync.WaitGroup
	var rspMap = make(map[string]Title)
	wg.Add(len(input.List))

	work := func(url string) {
		defer wg.Done()
		worker(&rspMap, url)
	}

	for _, url := range input.List {
		go work(url)
	}

	wg.Wait()

	for _, url := range input.List {
		output.Result = append(output.Result, rspMap[url])
	}

	c.JSON(http.StatusOK, output)
}

func worker(m *map[string]Title, url string) {
	resp, err := http.Get(url)
	if err != nil {
		(*m)[url] = Title{
			Status:  http.StatusNotFound,
			Message: err.Error(),
		}
		return
	}

	body, err := html.Parse(resp.Body)
	if err != nil {
		(*m)[url] = Title{
			Status:  http.StatusMethodNotAllowed,
			Message: err.Error(),
		}
		return
	}

	title := extractPageTitle(body)

	(*m)[url] = Title{
		Status: http.StatusOK,
		Title:  title,
	}
}

func extractPageTitle(node *html.Node) string {
	if node.Type == html.ElementNode && node.Data == "title" {
		return node.FirstChild.Data
	}

	var title string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		title = extractPageTitle(child)
		if title != "" {
			return title
		}
	}

	return title
}
