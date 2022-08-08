package controllers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func GetRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		shortUrl := c.Param("shortUrl")
		originalUrl, err := findOriginalLink(shortUrl)
		if err != nil {
			c.Redirect(301, "/404.html")
		} else {
			c.Redirect(301, originalUrl)
		}
	}
}

type GetUrlRequest struct {
	Url string `json:"url"`
}

func GetShortUrl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUrlRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		matched, err := regexp.MatchString(`.+\..+`, req.Url)
		if (err != nil) || !matched {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Введите корректный url"})
			return
		}
		shortUrl, err := createShortLink(req.Url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"shortUrl": shortUrl})
	}
}
