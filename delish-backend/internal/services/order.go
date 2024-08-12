package services

import (
	"delish-backend/internal/models"
	"delish-backend/internal/utils"
	"errors"
)

// Data type for DishOrder, containing the ordered dish's ID and quantity
type DishOrder struct {
	DishID string
	Qty    uint64
}

// Parameters for CREATE operations
type CreateOrderParams struct {
	CustomerName string
	DishOrders   []DishOrder
}

// Create a new Order record
func CreateOrder(params CreateOrderParams) error {
	// Get the database instance
	db, err := utils.GetDbInstance()
	if err != nil {
		return err
	}

	// Capitalize the first letter of the customer name
	formattedName := utils.ToTitleCase(params.CustomerName)

	// Create a new Order record
	order := models.Order{
		ID:           utils.GenerateUUID(),
		CustomerName: formattedName,
		TotalPrice:   0, // Initialize the total price to 0 for now
	}

	// Save the new Order record
	if err := db.Create(&order).Error; err != nil {
		// Return the error if any
		return err
	}

	// Initialize total price
	var totalPrice uint64 = 0

	// Iterate over the DishOrders
	for _, orderedDish := range params.DishOrders {
		// Validate the ordered dish's quantity
		if orderedDish.Qty < 1 {
			return errors.New("Quantity must be at least 1")
		}

		// Get the Order record by its ID listed in the DishOrder slice
		currentDish, err := GetDishById(orderedDish.DishID)

		// Handdle errors if any
		if err != nil {
			return err
		}
		if currentDish == nil {
			return errors.New("Dish not found")
		}

		// Calculate the price for this dish order
		orderPrice := currentDish.Price * orderedDish.Qty

		// Add to the total price
		totalPrice += orderPrice

		// Create a new DishOrder
		orderDetail := models.DishOrder{
			ID:      utils.GenerateUUID(),
			OrderID: order.ID,
			DishID:  currentDish.ID,
		}

		// Save the new DishOrder record
		if err := db.Create(&orderDetail).Error; err != nil {
			// Return the error if any
			return err
		}
	}

	// Update the Order's total price
	order.TotalPrice = totalPrice
	if err := db.Save(&order).Error; err != nil {
		// Return the error if any
		return err
	}

	// Return nil if no error occurred
	return nil
}

// Get all Order records
func GetAllOrders() ([]*models.Order, error) {
	// Get the database instance
	db, err := utils.GetDbInstance()
	if err != nil {
		return nil, err
	}

	// Prepare a slice to store the Order records
	var orders []*models.Order

	// Get all Order records and preload associated DishOrders and Dishes
    if err := db.Preload("DishOrders").Preload("DishOrders.Dish").Find(&orders).Error; err != nil {
        // Return the error if any
        return nil, err
    }

	// Return the Order records if no error occurred
	return orders, nil
}

// Get one Order record by its ID
func GetOrderById(id string) (*models.Order, error) {
	// Get the database instance
	db, err := utils.GetDbInstance()
	if err != nil {
		return nil, err
	}

	// Prepare a variable to store the Order record
	order := &models.Order{}

	// Get the Order record by its ID
	if err := db.Preload("DishOrders").Preload("DishOrders.Dish").Where("id = ?", id).First(order).Error; err != nil {
		// Return the error if any
		return nil, err
	}

	// Return the Order record if no error occurred
	return order, nil
}

// Parameters for UPDATE operations
type UpdateOrderParams struct {
	CustomerName string
}

// Update a Order record by its ID
func UpdateOrder(id string, params UpdateOrderParams) error {
	// Get the database instance
	db, err := utils.GetDbInstance()
	if err != nil {
		return err
	}

	// Get the Order record
	order, err := GetOrderById(id)
	if err != nil {
		return err
	}

	// Capitalize the first letter of the customer name
	customerName := utils.ToTitleCase(params.CustomerName)

	// Update the Order record with the new values
	order.CustomerName = customerName

	// Save the updated Order record
	if err := db.Save(order).Error; err != nil {
		return err
	}

	// Return nil if no error occurred
	return nil
}

// Delete an Order record by its ID
func DeleteOrderById(id string) error {
	// Get the database instance
	db, err := utils.GetDbInstance()
	if err != nil {
		return err
	}

	// Get the Order record
	order, err := GetOrderById(id)
	if err != nil {
		return err
	}

	// Delete the Order record
	if err := db.Delete(order).Error; err != nil {
		return err
	}

	// Return nil if no error occurred
	return nil
}
