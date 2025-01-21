package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	stock int
}

var Inventory []Product

func AddProduct(id int, name string, price float64, stock int) error {

	for _, product := range Inventory {
		if product.ID == id {
			return errors.New("Product with same ID already exists")
		}
	}
	newProduct := Product{ID: id, Name: name, Price: price, stock: stock}
	Inventory = append(Inventory, newProduct)
	return nil
}

func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("Stock cannot be negative")
	}
	for i, product := range Inventory {
		if product.ID == id {
			Inventory[i].stock = newStock
			return nil
		}
	}
	return errors.New("Product not found")
}

func SearchProduct(productQuery any) (*Product, error) {
	switch productQuery.(type) {
	case int:
		for _, product := range Inventory {
			if product.ID == productQuery {
				return &product, nil
			}
		}
		return nil, errors.New("Product not found")
	case string:
		for _, product := range Inventory {
			if product.Name == productQuery {
				return &product, nil
			}
		}
		return nil, errors.New("Product not found")
	default:
		return nil, errors.New("Invalid search query")
	}
}

func DisplayInverntory() {
	fmt.Println("\nID\tName\t\tPrice\tStock")
	fmt.Println("----------------------------")
	for _, product := range Inventory {
		fmt.Printf("%d\t%-15s\t%.2f\t%d\n", product.ID, product.Name, product.Price, product.stock)
	}
}

func SortByPrice() {
	sort.Slice(Inventory, func(i, j int) bool {
		return Inventory[i].Price < Inventory[j].Price
	})
}

func SortByStock() {
	sort.Slice(Inventory, func(i, j int) bool {
		return Inventory[i].stock < Inventory[j].stock
	})
}

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	err := AddProduct(1, "Product 1", 100.0, 10)
// 	if err != nil {
// 		fmt.Println("Error adding product: ", err)
// 	}

// 	err = AddProduct(2, "Product 2", 200.0, 20)
// 	if err != nil {
// 		fmt.Println("Error adding product: ", err)
// 	}

// 	err = AddProduct(3, "Product 3", 300.0, 30)
// 	if err != nil {
// 		fmt.Println("Error adding product: ", err)
// 	}

// 	for {
// 		fmt.Println("====================================")
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Add Product")
// 		fmt.Println("2. Update Stock")
// 		fmt.Println("3. Search Product")
// 		fmt.Println("4. Display Inventory")
// 		fmt.Println("5. Sort by Price")
// 		fmt.Println("6. Sort by Stock")
// 		fmt.Println("7. Exit")
// 		fmt.Print("Enter your choice: ")
// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			var id int
// 			var name string
// 			var price float64
// 			var stock int
// 			fmt.Print("\nEnter Product ID: ")
// 			fmt.Scanln(&id)
// 			fmt.Print("Enter Product Name: ")
// 			nameInput, _ := reader.ReadString('\n')
// 			name = strings.TrimSpace(nameInput)

// 			fmt.Print("Enter Product Price: ")
// 			priceInput, _ := reader.ReadString('\n')
// 			priceInput = strings.TrimSpace(priceInput)
// 			price, err = strconv.ParseFloat(priceInput, 64)
// 			if err != nil {
// 				fmt.Println("Invalid price")
// 				continue
// 			}
// 			fmt.Print("Enter Product Stock: ")
// 			fmt.Scanln(&stock)
// 			err = AddProduct(id, name, price, stock)
// 			if err != nil {
// 				fmt.Println("Error adding product: ", err)
// 			} else {
// 				fmt.Println("Product added successfully")
// 			}
// 		case 2:
// 			var id int
// 			var newStock int
// 			fmt.Print("\nEnter Product ID: ")
// 			fmt.Scanln(&id)
// 			fmt.Print("Enter New Stock: ")
// 			fmt.Scanln(&newStock)
// 			err := UpdateStock(id, newStock)
// 			if err != nil {
// 				fmt.Println("Error updating stock: ", err)
// 			} else {
// 				fmt.Println("Stock updated successfully")
// 			}
// 		case 3:
// 			fmt.Println("\nEnter Product ID or Name: ")
// 			var searchQuery string
// 			queryInput, _ := reader.ReadString('\n')
// 			searchQuery = strings.TrimSpace(queryInput)
// 			var id int
// 			if _, err := fmt.Sscanf(searchQuery, "%d", &id); err == nil {
// 				product, err := SearchProduct(id)
// 				if err != nil {
// 					fmt.Println("Error: ", err)
// 				} else {
// 					fmt.Println("Product found:")
// 					fmt.Println("ID: ", product.ID)
// 					fmt.Println("Name: ", product.Name)
// 					fmt.Println("Price: ", product.Price)
// 					fmt.Println("Stock: ", product.stock)
// 				}
// 			} else {
// 				product, err := SearchProduct(searchQuery)
// 				if err != nil {
// 					fmt.Println("Error: ", err)
// 				} else {
// 					fmt.Println("Product found:")
// 					fmt.Println("ID: ", product.ID)
// 					fmt.Println("Name: ", product.Name)
// 					fmt.Println("Price: ", product.Price)
// 					fmt.Println("Stock: ", product.stock)
// 				}
// 			}
// 		case 4:
// 			DisplayInverntory()

// 		case 5:
// 			SortByPrice()
// 			fmt.Println("Inventory sorted by price")
// 			DisplayInverntory()
// 		case 6:
// 			SortByStock()
// 			fmt.Println("Inventory sorted by stock")
// 			DisplayInverntory()
// 		case 7:
// 			fmt.Println("Exiting...")
// 			return

// 		default:
// 			fmt.Println("Invalid choice")
// 		}
// 		if choice == 7 {
// 			break
// 		}
// 	}
// }
