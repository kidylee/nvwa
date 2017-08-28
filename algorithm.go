package main

// func binaryLeftSearch(os []*Order, order *Order, left, right int) (position int) {
// 	if os[left].price == os[right].price {
// 		return left
// 	}

// 	mid := (left + right) / 2
// 	if os[mid].price < order.price {
// 		return binaryLeftSearch(os, order, mid+1, right)
// 	}
// 	return binaryLeftSearch(os, order, left, mid)

// }

// func binaryRightSearch(os []*Order, order *Order, left, right int) (position int) {
// 	if len(os) == 0 {
// 		return 0
// 	}

// 	if len(os) == 1 {
// 		if order.price-os[0].price < 0 {
// 			return 0
// 		}
// 		return 1

// 	}

// 	if os[left].price == os[right].price {

// 		return right
// 	}

// 	mid := (left + right + 1) / 2
// 	if os[mid].price > order.price {
// 		return binaryRightSearch(os, order, left, mid-1)
// 	}
// 	return binaryRightSearch(os, order, mid, right)
// }

// func expotentialSearch(os []*Order, len int, order *Order) {

// }
