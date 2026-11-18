package order

import (
	"fmt"
)

type Dish struct {
	Name        string
	PriceSantim int
}

type Order struct {
	isPaid bool
	Items  []Dish
	Area   string
}

var Shiro = Dish{
	Name:        "Shiro",
	PriceSantim: 200,
}

var Order1 = Order{
	isPaid: false,
	Items:  []Dish{Shiro},
	Area:   "Bole",
}

func (o *Order) AddItem(d Dish) {
	o.Items = append(o.Items, d)
}

func Mach() {
	tibs := Dish{
		Name:        "Tibs",
		PriceSantim: 400,
	}
	Order1.AddItem(tibs)
}

func (o Order) SubtotalSantim() int {
	total := 0
	for _, d := range o.Items {
		total += d.PriceSantim
	}
	return total
}

func DeliveryFeeSantim(area string) (int, error) {
	switch area {
	case "Bole":
		return 3000, nil
	case "Kazanchis":
		return 3500, nil
	}
	return 0, fmt.Errorf("unserved area ", area)
}

func (o Order) TotalSantim() int {
	subtotal := o.SubtotalSantim()
	serviceFee := subtotal * 5 / 100
	deliveryFee, _ := DeliveryFeeSantim(o.Area)
	return subtotal + serviceFee + deliveryFee
}

