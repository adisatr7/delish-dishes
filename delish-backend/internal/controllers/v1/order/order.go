package order

import (
	order_service "delish-backend/internal/services/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a new Order record
func CreateOrder(ctx *gin.Context) {
	// Parse the request body
	var body struct {
		CustomerName *string
		Note         *string
		DishOrders   []order_service.DishOrder
	}
	ctx.BindJSON(&body)

	// Validate the request body
	if body.CustomerName == nil || *body.CustomerName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Customer name is required!",
			"success": false,
		})
		return
	}
	if len(body.DishOrders) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dish orders are required!",
			"success": false,
		})
		return
	}

	// Create a new Order record
	err := order_service.CreateOrder(order_service.CreateOrderParams{
		CustomerName: *body.CustomerName,
		DishOrders:   body.DishOrders,
	})

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 201 Created response
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully!",
		"success": true,
	})
}

// Get all Order records
func GetAllOrders(ctx *gin.Context) {
	// Get all Order records
	orders, err := order_service.GetAllOrders()

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response with the Order records
	ctx.JSON(http.StatusOK, gin.H{
		"data":    orders,
		"success": true,
	})
}

// Get one Order record by its ID
func GetOrderById(ctx *gin.Context) {
	// Get the Order ID from the URL
	orderId := ctx.Param("id")

	// Get the Order record by its ID
	order, err := order_service.GetOrderById(orderId)

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response with the Order record
	ctx.JSON(http.StatusOK, gin.H{
		"data":    order,
		"success": true,
	})
}

// Update an Order record by its ID
func UpdateOrderById(ctx *gin.Context) {
	// Get the Order ID from the URL
	orderId := ctx.Param("id")

	// Parse the request body
	var body struct {
		CustomerName *string
	}
	ctx.BindJSON(&body)

	// Validate the request body
	if body.CustomerName == nil || *body.CustomerName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Customer name is required!",
			"success": false,
		})
		return
	}

	// Update the Order record by its ID. For now, only the customer name can be updated
	err := order_service.UpdateOrder(orderId, order_service.UpdateOrderParams{
		CustomerName: *body.CustomerName,
	})

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order has been updated successfully!",
		"success": true,
	})
}

// Delete an Order record by its ID
func DeleteOrderById(ctx *gin.Context) {
	// Get the Order ID from the URL
	orderId := ctx.Param("id")

	// Delete the Order record by its ID
	err := order_service.DeleteOrderById(orderId)

	// On error, return a 500 Internal Server Error response
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"success": false,
		})
		return
	}

	// On success, return a 200 OK response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Order has been deleted successfully!",
		"success": true,
	})
}
