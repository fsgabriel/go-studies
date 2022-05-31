package main

import (
	"encoding/json"
	"io"

	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	Id      int     `json:"id"`
	Nome    string  `json:"nome"`
	Cor     string  `json:"cor"`
	Preco   float64 `json:"preco"`
	Estoque bool    `json:"estoque"`
	Codigo  string  `json:"codigo"`
}

func getAllProducts(c *gin.Context) {
	content, err := os.ReadFile("./products.json")
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var p Products

	if err := json.Unmarshal(content, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": &p,
	})
}

func getProductByID(c *gin.Context) {
	idStr, _ := c.Params.Get("id")

	id, _ := strconv.Atoi(idStr)

	content, err := os.ReadFile("./products.json")
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	var p Products

	if err := json.Unmarshal(content, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result Product

	for _, produto := range p.Products {
		if produto.Id == id {
			result = produto
		}
	}

	c.JSON(200, &result)
}

func createProduct(c *gin.Context) {
	var p Products
	var newProduct Product

	content, err := os.ReadFile("./products.json")
	if err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if err := json.Unmarshal(content, &p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	body, err := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(body, &newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	p.Products = append(p.Products, newProduct)

	c.JSON(http.StatusCreated, &p.Products)
}

func main() {
	var p Products

	content, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal(content, &p); err != nil {
		panic(err)
	}

	router := gin.Default()

	v1 := router.Group("/products")

	v1.GET("/", func(c *gin.Context) {
		token := c.Request.Header.Get("Bearer")

		if token != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token invalido"})
			return
		}

		c.JSON(200, gin.H{
			"data": &p,
		})
	})

	v1.GET("/:id", func(c *gin.Context) {
		token := c.Request.Header.Get("Bearer")

		if token != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token invalido"})
			return
		}

		idStr, _ := c.Params.Get("id")

		id, _ := strconv.Atoi(idStr)

		var result Product

		for _, produto := range p.Products {
			if produto.Id == id {
				result = produto
			}
		}

		c.JSON(200, &result)
	})

	v1.POST("/", func(c *gin.Context) {
		token := c.Request.Header.Get("Bearer")

		if token != "123456" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "token invalido"})
			return
		}

		var newProduct Product

		body, _ := io.ReadAll(c.Request.Body)

		if err := json.Unmarshal(body, &newProduct); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if newProduct.Cor == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "cor do produto é obrigatorio"})
			return
		}

		if newProduct.Nome == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "nome do produto é obrigatorio"})
			return
		}

		if newProduct.Preco == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "preco do produto é obrigatorio"})
			return
		}

		if newProduct.Codigo == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "codigo do produto é obrigatorio"})
			return
		}

		p.Products = append(p.Products, newProduct)

		c.JSON(http.StatusCreated, &p.Products)
	})

	router.Run()
}
