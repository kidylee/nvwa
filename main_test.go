package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestPlaceNewOrders(t *testing.T) {
	ob := &OrderBook{
		asks: make([]*Order, 0),
		bids: make([]*Order, 0),
	}

	buy := &Order{
		amount: decimal.NewFromFloat(1.3),
		buy:    true,
		price:  decimal.NewFromFloat(30.6661),
		time:   time.Now(),
	}

	placed := ob.placeNewOrder(buy)
	if len(placed) != 0 {
		t.Errorf("Should not placed any %v", placed)
	}

	sell := &Order{
		amount: decimal.NewFromFloat(1.3),
		buy:    false,
		price:  decimal.NewFromFloat(30.6661),
		time:   time.Now(),
	}

	placed = ob.placeNewOrder(sell)
	if len(placed) != 2 {
		t.Errorf("Should have 2 placed %v ", placed)
	}
	if len(ob.asks) != 0 || len(ob.bids) != 0 {
		t.Errorf("The list should empty.")
	}

	buy = &Order{
		id:     "1",
		amount: decimal.NewFromFloat(1.3),
		buy:    true,
		price:  decimal.NewFromFloat(30.6661),
		time:   time.Now(),
	}
	ob.placeNewOrder(buy)
	buy = &Order{
		id:     "2",
		amount: decimal.NewFromFloat(1.8),
		buy:    true,
		price:  decimal.NewFromFloat(30.6661),
		time:   time.Now(),
	}

	ob.placeNewOrder(buy)

	sell = &Order{
		amount: decimal.NewFromFloat(1.7),
		buy:    false,
		price:  decimal.NewFromFloat(30.666),
		time:   time.Now(),
	}

	placed = ob.placeNewOrder(sell)

	if len(placed) != 4 {
		t.Errorf("Should have 4 placed %v ", placed)
	}

	amount := decimal.Zero
	for _, po := range placed {

		amount = amount.Add(po.amount)
	}

	result, _ := decimal.NewFromString("3.4")
	if !amount.Equal(result) {
		t.Errorf("Placed invalid amount %v ", amount.String())
	}

	result, _ = decimal.NewFromString("1.4")
	if len(ob.asks) != 1 || !ob.asks[0].amount.Equal(result) || ob.asks[0].id != "2" {
		t.Errorf("Placed invalid order left %v ", ob.asks[0])
	}

	sell = &Order{
		amount: decimal.NewFromFloat(1.7),
		buy:    false,
		price:  decimal.NewFromFloat(30.8),
		time:   time.Now(),
	}

	placed = ob.placeNewOrder(sell)
	sell = &Order{
		amount: decimal.NewFromFloat(1.7),
		buy:    false,
		price:  decimal.NewFromFloat(30.7),
		time:   time.Now(),
	}

	placed = ob.placeNewOrder(sell)
	sell = &Order{
		amount: decimal.NewFromFloat(1.7),
		buy:    false,
		price:  decimal.NewFromFloat(30.71),
		time:   time.Now(),
	}

	placed = ob.placeNewOrder(sell)

	buy = &Order{
		id:     "2",
		amount: decimal.NewFromFloat(1.8),
		buy:    true,
		price:  decimal.NewFromFloat(30.9),
		time:   time.Now(),
	}

	placed = ob.placeNewOrder(buy)
	amount = decimal.Zero
	for _, po := range placed {

		amount = amount.Add(po.amount)
	}

	result, _ = decimal.NewFromString("3.6")
	if !amount.Equal(result) {
		t.Errorf("Placed invalid amount %v ", amount.String())
	}

	result, _ = decimal.NewFromString("1.4")
	if len(ob.bids) != 1 || !ob.bids[0].amount.Equal(result) || ob.bids[0].id != "2" {
		t.Errorf("Placed invalid order left %v ", ob.asks[0])
	}
}

func TestAsksOrder(t *testing.T) {
	ob := &OrderBook{
		asks: make([]*Order, 0),
		bids: make([]*Order, 0),
	}

	for i := 1; i <= 1000; i++ {
		o := &Order{
			buy:   true,
			price: decimal.NewFromFloat(rand.Float64()),
			time:  time.Now(),
		}

		ob.addToBook(o)
	}

	for i := range ob.asks {

		if i == 0 || i == len(ob.asks) {
			continue
		}
		if ob.asks[i-1].price.GreaterThan(ob.asks[i].price) {
			t.Errorf("Sorting of %f > %f failed.", ob.asks[i].price, ob.asks[i+1].price)
		}

	}

}

func TestEqualPrice(t *testing.T) {
	o1 := &Order{
		id:    "1",
		buy:   true,
		price: decimal.NewFromFloat(1),
		time:  time.Now(),
	}
	o2 := &Order{
		id:    "2",
		buy:   true,
		price: decimal.NewFromFloat(1),
		time:  time.Now(),
	}

	os := []*Order{o1, o2}
	sort.Sort(byPriceTimeDesc(os))

	if os[0].id != "1" {
		t.Errorf("The ID should be 1 but actual is %v", os[0].id)
	}

	sort.Sort(byPriceTimeAsc(os))

	if os[1].id != "1" {
		t.Errorf("The ID should be 1 but actual is %v", os[0].id)
	}

}

func TestPrice(t *testing.T) {
	o1 := &Order{
		id:    "1",
		buy:   true,
		price: decimal.NewFromFloat(1),
		time:  time.Now(),
	}
	o2 := &Order{
		id:    "2",
		buy:   true,
		price: decimal.NewFromFloat(2),
		time:  time.Now(),
	}

	os := []*Order{o1, o2}
	sort.Sort(byPriceTimeDesc(os))

	if os[0].id != "1" {
		t.Errorf("The ID should be 1 but actual is %v", os[0].id)
	}

	sort.Sort(byPriceTimeAsc(os))

	if os[1].id != "1" {
		t.Errorf("The ID should be 1 but actual is %v", os[0].id)
	}

}

func TestBidsOrder(t *testing.T) {
	ob := &OrderBook{
		asks: make([]*Order, 0),
		bids: make([]*Order, 0),
	}

	for i := 1; i <= 1000; i++ {
		o := &Order{
			buy:   false,
			price: decimal.NewFromFloat(rand.Float64()),
			time:  time.Now(),
		}

		ob.addToBook(o)
	}

	for i := range ob.asks {

		if i == 0 || i == len(ob.asks) {
			continue
		}
		if ob.asks[i-1].price.LessThan(ob.asks[i].price) {
			t.Errorf("Sorting of %f > %f failed.", ob.asks[i].price, ob.asks[i+1].price)
		}

	}

}

func BenchmarkAsksOrder(t *testing.B) {
	ob := &OrderBook{
		asks: make([]*Order, 0),
		bids: make([]*Order, 0),
	}

	for i := 0; i <= t.N; i++ {

		o := &Order{
			buy:   true,
			price: decimal.NewFromFloat(rand.Float64()),
			time:  time.Now(),
		}

		ob.addToBook(o)

	}

}
