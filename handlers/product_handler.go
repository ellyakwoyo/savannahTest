package handlers

import (
	"net/http"
	"savannahTest/services"
	"strconv"

	"github.com/labstack/echo/v4"
	"savannahTest/models"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a product with name, description, and price.
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.Product true "Product data"
// @Success 200 {object} map[string]interface{} "Product created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "invalid request",
		})
	}

	createdProduct, err := h.productService.CreateProduct(product.Name, product.Description, product.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": "Failed to create product",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Product created successfully",
		"data":    createdProduct,
	})
}

// GetAllProducts godoc
// @Summary Retrieve all products
// @Description Fetches all products from the database.
// @Tags Product
// @Produce json
// @Success 200 {object} map[string]interface{} "Products fetched successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c echo.Context) error {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Products fetched successfully",
		"data":    products,
	})
}

// GetProductByID godoc
// @Summary Retrieve a product by ID
// @Description Fetches a single product using its ID.
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Product fetched successfully"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Product fetched successfully",
		"data":    product,
	})
}

// UpdateProduct godoc
// @Summary Update an existing product
// @Description Updates the product details by ID.
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Updated product data"
// @Success 200 {object} map[string]interface{} "Product updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":  http.StatusBadRequest,
			"error": "Invalid request",
		})
	}
	productID := uint(id)
	if err := h.productService.UpdateProduct(productID, product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":  http.StatusInternalServerError,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Product updated successfully",
		"data":    product,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Deletes a product by ID.
// @Tags Product
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Product deleted successfully"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.productService.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":  http.StatusNotFound,
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Product deleted successfully",
	})
}
