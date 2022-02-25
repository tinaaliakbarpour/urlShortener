package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlShortner/entity"
	"urlShortner/shortener"
	"urlShortner/store"
)

//CreateShortUrl will be called on post request th return value would be a short url
func (handler) CreateShortUrl(c *gin.Context) {
	var creationRequest entity.UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.Shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.Store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

//HandleShortUrlRedirect will be called on get request the return the actual url
func (handler) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.Store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, initialUrl)
}
