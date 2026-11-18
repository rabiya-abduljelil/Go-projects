package order

import "errors"

func LineTotal(priceSantim, qty int) int {
	return priceSantim * qty
}

func ServiceFeeSantim(subtotal int) int {
	return subtotal / 10
}

func DeliveryFeeSantim(area string) (int, error) {
	switch area {
	case "Bole", "Kazanchis", "Piazza", "CMC":
		return 10000, nil
	default:
		return 0, errors.New("area outside served zones")
	}
}