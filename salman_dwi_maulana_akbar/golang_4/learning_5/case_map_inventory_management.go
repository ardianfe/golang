package main

import "fmt"

// Product struct defines the attributes of a product
type Product struct {
	ID       string
	Name     string
	Category string
	Price    float64
	Quantity int
}

// Inventory struct contains a map of products with product ID as key
type Inventory struct {
	Products map[string]Product
}

// Function to add a product to the inventory
func (inv *Inventory) AddProduct(p Product) {
	inv.Products[p.ID] = p
}

// Function to update a product in the inventory
func (inv *Inventory) UpdateProduct(p Product) {
	if _, exists := inv.Products[p.ID]; exists {
		inv.Products[p.ID] = p
	} else {
		fmt.Println("Product not found!")
	}
}

// Function to delete a product from the inventory
func (inv *Inventory) DeleteProduct(productID string) {
	if _, exists := inv.Products[productID]; exists {
		delete(inv.Products, productID)
	} else {
		fmt.Println("Product not found!")
	}
}

// Function to display all products in the inventory
func (inv Inventory) DisplayProducts() {
	fmt.Println("Inventory Products:")
	for _, product := range inv.Products {
		fmt.Printf("ID: %s, Name: %s, Category: %s, Price: %.2f, Quantity: %d\n", product.ID, product.Name, product.Category, product.Price, product.Quantity)
	}
}

// Function to find products by category
func (inv Inventory) FindProductsByCategory(category string) []Product {
	var products []Product
	for _, product := range inv.Products {
		if product.Category == category {
			products = append(products, product)
		}
	}
	return products
}

// Function to calculate the total value of the inventory
func (inv Inventory) TotalInventoryValue() float64 {
	total := 0.0
	for _, product := range inv.Products {
		total += product.Price * float64(product.Quantity)
	}
	return total
}

func main() {
	// Initializing the inventory
	inventory := Inventory{Products: make(map[string]Product)}

	// Adding products
	inventory.AddProduct(Product{ID: "P001", Name: "Laptop", Category: "Electronics", Price: 1200.00, Quantity: 10})
	inventory.AddProduct(Product{ID: "P002", Name: "Phone", Category: "Electronics", Price: 800.00, Quantity: 15})
	inventory.AddProduct(Product{ID: "P003", Name: "Desk", Category: "Furniture", Price: 150.00, Quantity: 20})

	// Displaying products
	inventory.DisplayProducts()

	// Updating a product
	inventory.UpdateProduct(Product{ID: "P002", Name: "Smartphone", Category: "Electronics", Price: 850.00, Quantity: 15})
	fmt.Println("\nAfter updating product P002:")
	inventory.DisplayProducts()

	// Deleting a product
	inventory.DeleteProduct("P003")
	fmt.Println("\nAfter deleting product P003:")
	inventory.DisplayProducts()

	// Finding products by category
	fmt.Println("\nFinding products in category 'Electronics':")
	electronics := inventory.FindProductsByCategory("Electronics")
	for _, product := range electronics {
		fmt.Printf("ID: %s, Name: %s, Category: %s, Price: %.2f, Quantity: %d\n", product.ID, product.Name, product.Category, product.Price, product.Quantity)
	}

	// Calculating total inventory value
	totalValue := inventory.TotalInventoryValue()
	fmt.Printf("\nTotal Inventory Value: %.2f\n", totalValue)
}
