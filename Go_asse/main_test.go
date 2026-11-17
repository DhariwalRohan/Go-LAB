package main

import "testing"

// Test 1: Add item
func TestAddItem(t *testing.T) {
	inv := &Inventory{}

	item := Item{
		ID:       1,
		Name:     "Laptop",
		Quantity: 5,
		Price:    50000,
	}

	err := inv.Add(item)

	if err != nil {
		t.Fatal(err)
	}

	if len(inv.items) != 1 {
		t.Fatal("Item was not added")
	}
}

// Test 2: Update item
func TestUpdateItem(t *testing.T) {
	inv := &Inventory{}

	inv.Add(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 50000})

	err := inv.Update(1, 10, 60000)

	if err != nil {
		t.Fatal(err)
	}

	if inv.items[0].Quantity != 10 {
		t.Fatal("Quantity was not updated")
	}
}

// Test 3: Remove item
func TestRemoveItem(t *testing.T) {
	inv := &Inventory{}

	inv.Add(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 50000})

	err := inv.Remove(1)

	if err != nil {
		t.Fatal(err)
	}

	if len(inv.items) != 0 {
		t.Fatal("Item was not removed")
	}
}

// Test 4: Invalid item
func TestInvalidItem(t *testing.T) {
	inv := &Inventory{}

	err := inv.Add(Item{ID: 0, Name: "Laptop", Quantity: 5, Price: 50000})

	if err == nil {
		t.Fatal("Expected validation error")
	}
}
