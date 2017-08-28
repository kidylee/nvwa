package main

// import (
// 	"testing"
// 	"time"
// )

// func TestBinarySearch(t *testing.T) {
// 	os := make([]*Order, 0)
// 	o1 := &Order{
// 		buy:   true,
// 		price: 1,
// 		time:  time.Now(),
// 	}
// 	o2 := &Order{
// 		buy:   true,
// 		price: 3,
// 		time:  time.Now(),
// 	}
// 	o3 := &Order{
// 		buy:   true,
// 		price: 4,
// 		time:  time.Now(),
// 	}

// 	o4 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	os = append(os, o1, o2, o3, o3)
// 	if p := binaryLeftSearch(os, o4, 0, len(os)-1); p != 1 {
// 		t.Errorf("Binary search of 1 Failed, actual is %d.", p)

// 	}

// 	o4 = &Order{
// 		buy:   true,
// 		price: 5,
// 		time:  time.Now(),
// 	}

// 	if p := binaryLeftSearch(os, o4, 0, len(os)-1); p != 2 {
// 		t.Errorf("Binary search of 4 Failed, actual is %d.", p)

// 	}

// }

// func TestRightBinarySearchZero(t *testing.T) {
// 	os := make([]*Order, 0)
// 	o4 := &Order{
// 		buy:   true,
// 		price: 5,
// 		time:  time.Now(),
// 	}
// 	if p := binaryRightSearch(os, o4, 0, len(os)-1); p != 0 {
// 		t.Errorf("Binary search of 1 Failed, actual is %d.", p)

// 	}

// 	if p := binaryRightSearch(os, o4, 0, len(os)-1); p != 0 {
// 		t.Errorf("Binary search of 4 Failed, actual is %d.", p)

// 	}

// }

// func TestRightBinarySearchOne(t *testing.T) {
// 	os := make([]*Order, 0)
// 	o1 := &Order{
// 		buy:   true,
// 		price: 7,
// 		time:  time.Now(),
// 	}

// 	p := binaryRightSearch(os, o1, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o1}, os[p:]...)...)
// 	if os[0].price != 7 {
// 		t.Errorf("Order is Wrong %f", os[0].price)
// 	}

// 	o2 := &Order{
// 		buy:   true,
// 		price: 6,
// 		time:  time.Now(),
// 	}
// 	p = binaryRightSearch(os, o2, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o2}, os[p:]...)...)
// 	if os[0].price != 6 || os[1].price != 7 {
// 		t.Errorf("Order is Wrong %f %f", os[0].price, os[1].price)
// 	}

// 	o3 := &Order{
// 		id:    "test",
// 		buy:   true,
// 		price: 1,
// 		time:  time.Now(),
// 	}
// 	p = binaryRightSearch(os, o3, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o3}, os[p:]...)...)
// 	if os[0].price != 1 || os[1].price != 6 || os[2].price != 7 {
// 		t.Errorf("Order is Wrong %f %f", os[0].price, os[1].price)
// 	}
// 	p = binaryRightSearch(os, o3, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o3}, os[p:]...)...)
// 	p = binaryRightSearch(os, o3, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o3}, os[p:]...)...)
// 	p = binaryRightSearch(os, o3, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o3}, os[p:]...)...)

// 	p = binaryRightSearch(os, o3, 0, len(os)-1)
// 	os = append(os[:p], append([]*Order{o3}, os[p:]...)...)
// 	if p != 6 {
// 		t.Errorf("Order is Wrong %d", p)
// 	}

// }

// func TestRightBinarySearch(t *testing.T) {
// 	os := make([]*Order, 0)
// 	o1 := &Order{
// 		buy:   true,
// 		price: 1,
// 		time:  time.Now(),
// 	}
// 	o2 := &Order{
// 		buy:   true,
// 		price: 3,
// 		time:  time.Now(),
// 	}
// 	o3 := &Order{
// 		buy:   true,
// 		price: 4,
// 		time:  time.Now(),
// 	}

// 	o4 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	os = append(os, o1, o2, o3, o3)
// 	if p := binaryRightSearch(os, o4, 0, len(os)-1); p != 1 {
// 		t.Errorf("Binary search of 1 Failed, actual is %d.", p)

// 	}

// 	o4 = &Order{
// 		buy:   true,
// 		price: 5,
// 		time:  time.Now(),
// 	}

// 	if p := binaryRightSearch(os, o4, 0, len(os)-1); p != 4 {
// 		t.Errorf("Binary search of 4 Failed, actual is %d.", p)

// 	}

// }
// func TestSameLeftBinarySearch(t *testing.T) {
// 	os := make([]*Order, 0)
// 	o1 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	o2 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	o3 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}

// 	o4 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	os = append(os, o1, o1, o1, o1, o2, o3)
// 	if p := binaryLeftSearch(os, o4, 0, len(os)-1); p != 0 {
// 		t.Errorf("Binary search of 0 Failed, actual is %d.", p)

// 	}
// }

// func TestSameRightBinarySearch(t *testing.T) {
// 	os := make([]*Order, 0)
// 	o1 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	o2 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	o3 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}

// 	o4 := &Order{
// 		buy:   true,
// 		price: 2,
// 		time:  time.Now(),
// 	}
// 	os = append(os, o1, o1, o1, o1, o2, o3)
// 	if p := binaryRightSearch(os, o4, 0, len(os)-1); p != 6 {
// 		t.Errorf("Binary search of 6 Failed, actual is %d.", p)

// 	}
// }
