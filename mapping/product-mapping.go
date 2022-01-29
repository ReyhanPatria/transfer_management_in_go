package mapping

import (
	"main/entity"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
)

func SetProductMapping(router *gin.Engine, cache_ *cache.Cache, productService service.ProductService) {
	router.GET("/products", func(c *gin.Context) {
		cacheKey := c.Request.RequestURI
		var value []entity.Product

		cache_.Once(&cache.Item{
			Key:   cacheKey,
			Value: &value,
			Do: func(c *cache.Item) (interface{}, error) {
				return productService.FindAll(), nil
			},
		})

		c.AbortWithStatusJSON(http.StatusOK, value)
	})

	router.GET("/product/:type/:key", func(c *gin.Context) {
		type_ := c.Param("type")
		key := c.Param("key")

		cacheKey := c.Request.RequestURI
		var value entity.Product

		err := cache_.Once(&cache.Item{
			Key:   cacheKey,
			Value: &value,
			Do: func(i *cache.Item) (interface{}, error) {
				return productService.FindByTypeAndKey(type_, key)
			},
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
			return
		}

		c.JSON(http.StatusOK, value)
	})

	router.GET("/product/:type/:key/min-limit", func(c *gin.Context) {
		type_ := c.Param("type")
		key := c.Param("key")

		cacheKey := c.Request.RequestURI
		var value int

		err := cache_.Once(&cache.Item{
			Key:   cacheKey,
			Value: &value,
			Do: func(i *cache.Item) (interface{}, error) {
				return productService.FindMinLimitByTypeAndKey(type_, key)
			},
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
			return
		}

		c.JSON(http.StatusOK, value)
	})

	router.GET("/product/:type/:key/max-limit", func(c *gin.Context) {
		type_ := c.Param("type")
		key := c.Param("key")

		cacheKey := c.Request.RequestURI
		var value int

		err := cache_.Once(&cache.Item{
			Key:   cacheKey,
			Value: &value,
			Do: func(i *cache.Item) (interface{}, error) {
				return productService.FindMaxLimitByTypeAndKey(type_, key)
			},
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
			return
		}

		c.JSON(http.StatusOK, value)
	})
}
