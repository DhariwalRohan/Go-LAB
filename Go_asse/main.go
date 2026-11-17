package main

import (
	"errors"
	"fmt"
	"strings"
)

// Item structure
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory stores items using a slice
type Inventory struct {
	items []Item
}

// Add item
func (inv *Inventory) Add(item Item) error {
	if item.ID <= 0 {
		return errors.New("invalid item ID")
	}

	if strings.TrimSpace(item.Name) == "" {
		return errors.New("item name cannot be empty")
	}

	if item.Quantity <= 0 || item.Price < 0 {
		return errors.New("invalid quantity or price")
	}

	for _, existing := range inv.items {
		if existing.ID == item.ID {
			return errors.New("item already exists")
		}
	}

	inv.items = append(inv.items, item)
	return nil
}

// Update item
func (inv *Inventory) Update(id, quantity int, price float64) error {
	if quantity <= 0 || price < 0 {
		return errors.New("invalid quantity or price")
	}

	for i := range inv.items {
		if inv.items[i].ID == id {
			inv.items[i].Quantity = quantity
			inv.items[i].Price = price
			return nil
		}
	}

	return errors.New("item not found")
}

// Remove item
func (inv *Inventory) Remove(id int) error {
	for i, item := range inv.items {
		if item.ID == id {
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

// Summary statistics
type Summary struct {
	TotalValue float64
	ItemCount  int
}

// Calculate summary using goroutine and channel
func (inv *Inventory) GetSummary() Summary {
	result := make(chan Summary)

	go func() {
		var summary Summary

		summary.ItemCount = len(inv.items)

		for _, item := range inv.items {
			summary.TotalValue += float64(item.Quantity) * item.Price
		}

		result <- summary
	}()

	return <-result
}

// Main function
func main() {
	inventory := &Inventory{}

	fmt.Println("INVENTORY MANAGEMENT SYSTEM")
	fmt.Println("---------------------------")

	// ADD
	err := inventory.Add(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 50000})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Laptop added successfully")
	}

	err = inventory.Add(Item{ID: 2, Name: "Mouse", Quantity: 10, Price: 500})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mouse added successfully")
	}

	// UPDATE
	err = inventory.Update(1, 8, 50000)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Laptop updated successfully")
	}

	// REMOVE
	err = inventory.Remove(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mouse removed successfully")
	}

	// DISPLAY ITEMS
	fmt.Println("\nRemaining Items:")
	for _, item := range inventory.items {
		fmt.Printf("ID: %d | Name: %s | Quantity: %d | Price: %.2f\n",
			item.ID, item.Name, item.Quantity, item.Price)
	}

	// CONCURRENT SUMMARY
	summary := inventory.GetSummary()

	fmt.Println("\nInventory Summary:")
	fmt.Printf("Total Items: %d\n", summary.ItemCount)
	fmt.Printf("Total Inventory Value: Rs. %.2f\n", summary.TotalValue)
}
