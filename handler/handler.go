package Handler

import (
	"net/http"

	//"github.com/ManueStoessel/redirecter/shortener"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome!")
}

// func CreateURL(c *gin.Context) {

// 	var json LongURL
// 	if err := c.ShouldBindJSON(&json); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	switch urlType := c.Param("type"); urlType {
// 	case "name":
// 		c.String(http.StatusOK, shortener.GetRandomName(0))
// 	case "hash":
// 		c.String(http.StatusOK, shortener.GetURLHash(json.URL))
// 	default:
// 		c.String(http.StatusBadRequest, "Only /url/name or /url/hash are accepted paths!")
// 	}
// }

// func ListURLs(c *gin.Context) {
// 	c.String(http.StatusOK, "Welcome!")
// }

// func GetURL(c *gin.Context) {
// 	c.String(http.StatusOK, "Welcome!")
//}

func (r *RedirecterHandler) Redirecter(c *gin.Context) {
	shorturl := c.Param("shorthand")

	longurl, err := r.Store.GetLongURL(shorturl)
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not find this short url!")
	}

	c.Redirect(http.StatusMovedPermanently, longurl)
}
