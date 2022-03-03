package handler

import (
	"crawler/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"net/http"
)

func (h *Handler) getTitlesByUrls(c *gin.Context) {
	var input model.Urls
	var output = NewResponse()

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{"urls (array) parameter required"})
		return
	}

	for _, url := range input.List {
		resp, err := http.Get(url)
		if err != nil {
			output.Result = append(output.Result, Title{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
			continue
		}

		body, err := html.Parse(resp.Body)
		if err != nil {
			output.Result = append(output.Result, Title{
				Status:  http.StatusMethodNotAllowed,
				Message: err.Error(),
			})
			continue
		}

		title := extractPageTitle(body)

		output.Result = append(output.Result, Title{
			Status: http.StatusOK,
			Title:  title,
		})
	}

	c.JSON(http.StatusOK, output)
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
