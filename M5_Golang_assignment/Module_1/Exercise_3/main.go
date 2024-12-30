package main

import (
	"errors"
	"fmt"
	"sort"
)

// Product struct to represent a product in the inventory
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory []Product // Slice to store the inventory

// AddProduct adds a new product to the inventory
func AddProduct(id int, name string, price float64, stock int) error {
	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID must be unique")
		}
	}

	if price < 0 || stock < 0 {
		return errors.New("price and stock cannot be negative")
	}

	inventory = append(inventory, Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	})
	return nil
}

// UpdateStock updates the stock of a product by ID
func UpdateStock(id int, newStock int) error {
	for i, product := range inventory {
		if product.ID == id {
			if newStock < 0 {
				return errors.New("stock cannot be negative")
			}
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

// SearchProduct searches for a product by ID or name
func SearchProduct(query string) (*Product, error) {
	for _, product := range inventory {
		if fmt.Sprintf("%d", product.ID) == query || product.Name == query {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

// DisplayInventory displays all products in the inventory
func DisplayInventory() {
	fmt.Println("\n--- Inventory ---")
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	for _, product := range inventory {
		fmt.Printf("%-10d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// SortInventory sorts the inventory by price or stock in ascending order
func SortInventory(by string) {
	switch by {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	default:
		fmt.Println("Invalid sort parameter. Use 'price' or 'stock'.")
	}
}

func main() {
	// Prepopulate some products
	AddProduct(1, "Laptop", 75000.00, 10)
	AddProduct(2, "Smartphone", 25000.00, 25)
	AddProduct(3, "Headphones", 1500.00, 50)

	DisplayInventory()

	// Update stock
	if err := UpdateStock(2, 30); err != nil {
		fmt.Println("Error:", err)
	}

	// Search for a product
	if product, err := SearchProduct("2"); err == nil {
		fmt.Println("\nProduct Found:", *product)
	} else {
		fmt.Println("Error:", err)
	}

	// Sort inventory by price
	SortInventory("price")
	fmt.Println("\nSorted by Price:")
	DisplayInventory()

	// Sort inventory by stock
	SortInventory("stock")
	fmt.Println("\nSorted by Stock:")
	DisplayInventory()
}
