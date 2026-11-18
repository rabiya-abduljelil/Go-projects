package main

import (
	"fmt"
	"example.com/myproject/internal/order"
)

func main() {
	// Call the three functions
	itemTotal := order.LineTotal(50000, 2)
	serviceFee := order.ServiceFeeSantim(itemTotal)
	deliveryFee, _ := order.DeliveryFeeSantim("Bole")

	// Print receipt
	fmt.Println("--- RECEIPT ---")
	fmt.Println("Item Total:", itemTotal)
	fmt.Println("Service Fee:", serviceFee)
	fmt.Println("Delivery Fee:", deliveryFee)
	fmt.Println("Grand Total:", itemTotal+serviceFee+deliveryFee)

	// Test error path with unserved area
	_, err := order.DeliveryFeeSantim("Adama")
	if err != nil {
		fmt.Println("\nError Test Passed:", err)
	}
}