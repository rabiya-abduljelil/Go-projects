// function name that start with small case are private / for the emodule
// function with name that startt with capital are exported
// pass by reference and value
// default = pass by value
// y= &x              is called pass by reference
// func main(){
//     var x =100
//     var y = &y
//     *y = 200
//     fmt.println(x)
// }


// if number of variable is unknown we use "...int"
//cap = return capacity
//len = length
//slice is a reference type used to store a sequence of values of the same type
//map is a reference type used to store a collection of key-value pairs
// struct is a user-defined type that groups together values of different types

//no class on go, no inheritance too... instead we use structure.


// value receiver: reads, does not modify
// func (o Order) SubtotalSantim() int {
// 	total := 0
// 	for _, d := range o.Items { total += d.PriceSantim }return total}

// pointer receiver: modifies the order
// func (o *Order) AddItem(d Dish) {
// 	o.Items = append(o.Items, d)}

// order.AddItem(shiro)
// fmt.Print1n(order.SubtotalSantim())

// Go takes the address for you



//Lab2

package order

import "errors"

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

var order1 = Order{
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
	order1.AddItem(tibs)
}

func (o Order) SubtotalSantim() int {
	total := 0
	for _, d := range o.Items {
		total += d.PriceSantim
	}
	return total
}

func deliveryFeeSantim(area string) (int, error) {
	switch area {
		case "Bole":
			return 3000, nil
		case "Kazanchis":
			return 3500, nil}
	return 0, fnt.Errorf("unserved area Xq", area)
}
fee, err := deliveryFeeSantim("Ayat")
if err != nil {
// handle it here, immediately
}

func (o Order) TotalSantim() int {
	subtotal := o.SubtotalSantim()
	serviceFee := subtotal * 5 / 100
	deliveryFee, _ := DeliveryFeeSantim(o.Area)
	return subtotal + serviceFee + deliveryFee
}


package main

func main() {
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


//composition

//panic and recover means handling runtime errors in Go
    //defer -- recover, code flow work normally
	//panic code flow stops and program crashes

