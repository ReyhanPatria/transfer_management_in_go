package main

import (
	"main/connect"
	"main/mapping"
	"main/repository"
	"main/service"

	"github.com/gin-gonic/gin"
)

var (
	router            *gin.Engine                  = gin.Default()
	db                connect.Database             = connect.NewDB()
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
)

func main() {
	mapping.SetProductMapping(router, db.Cache, productService)
	mapping.SetTransactionMappings(router, db.Cache, productService)

	router.Run()
}
