package main

import (
	"fmt"

	"Lab2/order"
)

func main() {
	fee, err := order.DeliveryFeeSantim("Ayat")
	if err != nil {
		fmt.Println("Error caught:", err)
		_ = fee
	}

	myOrder := order.Order{Area: "Bole"}

	myOrder.AddItem(order.Dish{Name: "Doro Wat", PriceSantim: 35000})
	myOrder.AddItem(order.Dish{Name: "Tibs", PriceSantim: 40000})
	myOrder.AddItem(order.Dish{Name: "Kitfo", PriceSantim: 45000})
	myOrder.AddItem(order.Dish{Name: "Injera", PriceSantim: 5000})
	

	subtotal := myOrder.SubtotalSantim()
	serviceFee := subtotal * 5 / 100
	total := myOrder.TotalSantim()

	fmt.Println("Subtotal Santim:", subtotal)
	fmt.Println("Service Fee (5%) Santim:", serviceFee)
	fmt.Println("Total Santim (with delivery):", total)
}
