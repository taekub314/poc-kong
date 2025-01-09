package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Product 1", Description: "Description of Product 1", Price: 10.99},
	{ID: 2, Name: "Product 2", Description: "Description of Product 2", Price: 20.49},
}

func main() {
	r := gin.Default()

	productRoutes := r.Group("/product")
	{
		productRoutes.GET("list", GetAllProducts)
		productRoutes.GET("/:id", GetProductByID)
		productRoutes.POST("create", CreateProduct)
		productRoutes.PUT("update/:id", UpdateProduct)
		productRoutes.DELETE("delete/:id", DeleteProduct)
	}

	r.Run(":8081")
}

func GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	for _, product := range products {
		if product.ID == productID {
			c.JSON(http.StatusOK, product)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set ID for the new product
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}
