package mapping

import (
	"context"
	"fmt"
	"main/entity"
	"main/service"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
)

func SetTransactionMappings(router *gin.Engine, cache_ *cache.Cache, productService service.ProductService) {
	// POST Mapping
	router.POST("/transaction", func(c *gin.Context) {
		var transaction entity.Transaction
		c.ShouldBindJSON(&transaction)

		bcaId := strings.ToUpper(transaction.BcaId)
		type_ := strings.ToUpper(transaction.Type)
		key := strings.ToUpper(transaction.Key)
		amount := transaction.Amount

		// Should validate bcaId exists
		/*
		 *
		 *
		 */

		minLimit, errMinLimit := productService.FindMinLimitByTypeAndKey(type_, key)
		maxLimit, errMaxLimit := productService.FindMaxLimitByTypeAndKey(type_, key)

		// Check if Product exists
		if errMinLimit != nil || errMaxLimit != nil {
			c.AbortWithStatusJSON(http.StatusOK, "Transaction not valid")
			return
		}

		if amount < minLimit {
			// Raise error
			c.AbortWithStatusJSON(http.StatusOK, "Transaction is lower that the minimum limit")
			return
		}

		// Get cached amount
		var cacheKey strings.Builder
		cacheKey.Grow(64)
		fmt.Fprintf(&cacheKey, "/%s/%s/%s/AMOUNT", bcaId, type_, key)

		cachedAmount, errGetCache := getCachedAmount(cache_, cacheKey.String())
		if errGetCache != nil {
			cachedAmount = 0
		}

		// Calculate new amount
		newAmount := cachedAmount + amount

		if newAmount > maxLimit {
			// Raise error
			c.AbortWithStatusJSON(http.StatusOK, "Total transactions goes over the maximum limit")
			return
		}

		c.JSON(http.StatusOK, updateCachedAmount(cache_, cacheKey.String(), newAmount))
	})

	// GET Mapping
	router.GET("/:bcaid/:type/:key/amount", func(c *gin.Context) {
		bcaId := strings.ToUpper(c.Param("bcaid"))
		type_ := strings.ToUpper(c.Param("type"))
		key := strings.ToUpper(c.Param("key"))

		var cacheKey strings.Builder
		cacheKey.Grow(64)
		fmt.Fprintf(&cacheKey, "/%s/%s/%s/AMOUNT", bcaId, type_, key)

		cachedAmount, err := getCachedAmount(cache_, cacheKey.String())

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
			return
		}

		c.JSON(http.StatusOK, cachedAmount)
	})
}

func getCachedAmount(cache_ *cache.Cache, key string) (int, error) {
	var value int
	err := cache_.Get(context.TODO(), key, &value)
	return value, err
}

func updateCachedAmount(cache_ *cache.Cache, key string, amount int) int {
	cache_.Set(&cache.Item{
		Key:   key,
		Value: amount,
		TTL:   24 * time.Hour,
	})
	return amount
}
