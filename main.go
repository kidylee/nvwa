package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("Hello")
}

//OrderBook is the matching engine core
type OrderBook struct {
	asks        []*Order //desc
	bids        []*Order //asc
	NewOrder    chan<- *Order
	PlacedOrder <-chan *PlacedOrder
}

//Order is a transaction
type Order struct {
	id     string
	buy    bool
	amount decimal.Decimal
	price  decimal.Decimal
	time   time.Time
}

func (ob OrderBook) New() *OrderBook {

	o := &OrderBook{
		asks:        make([]*Order, 0),
		bids:        make([]*Order, 0),
		NewOrder:    make(chan *Order),
		PlacedOrder: make(chan *PlacedOrder),
	}

	return o

}

func (ob *OrderBook) Run() {

}

func (ob *OrderBook) addToBook(o *Order) {
	if o.buy {
		ob.asks = append(ob.asks, o)
		sort.Sort(byPriceTimeDesc(ob.asks))

	} else {
		ob.bids = append(ob.bids, o)
		sort.Sort(byPriceTimeAsc(ob.bids))
	}
}

//PlacedOrder is matched transaction
type PlacedOrder struct {
	id      string
	orderID string
	amount  decimal.Decimal
	price   decimal.Decimal
	partial bool
}

type byPriceTimeDesc []*Order

func (a byPriceTimeDesc) Len() int { return len(a) }

func (a byPriceTimeDesc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a byPriceTimeDesc) Less(i, j int) bool {

	if a[i].price.Equal(a[j].price) {
		return a[i].time.Before(a[j].time)
	}

	return a[i].price.LessThan(a[j].price)

}

type byPriceTimeAsc []*Order

func (a byPriceTimeAsc) Len() int { return len(a) }

func (a byPriceTimeAsc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a byPriceTimeAsc) Less(i, j int) bool {

	if a[i].price.Equal(a[j].price) {
		return a[i].time.After(a[j].time)
	}

	return a[i].price.GreaterThan(a[j].price)

}

func (ob *OrderBook) placeNewOrder(o *Order) (placed []*PlacedOrder) {
	placed = make([]*PlacedOrder, 0)
	if ask := o; ask.buy {

		if len(ob.bids) == 0 {
			ob.addToBook(o)
			return
		}

		for _, bid := range ob.bids {
			if ask.amount.Equal(decimal.Zero) {
				return
			}

			if ask.price.LessThan(bid.price) {
				ob.addToBook(o)
				return
			}

			m := decimal.Min(ask.amount, bid.amount)

			pask := &PlacedOrder{
				orderID: ask.id,
				amount:  m,
				price:   bid.price,
				partial: m != ask.amount,
			}
			pbid := &PlacedOrder{
				orderID: bid.id,
				amount:  m,
				price:   bid.price,
				partial: m != bid.amount,
			}
			placed = append(placed, pask)
			placed = append(placed, pbid)

			ask.amount = ask.amount.Sub(m)
			bid.amount = bid.amount.Sub(m)

			if bid.amount.Equal(decimal.Zero) {
				ob.bids = ob.bids[1:]
				continue
			}

		}

	}
	if bid := o; !bid.buy {

		if len(ob.asks) == 0 {
			ob.addToBook(o)
			return
		}

		for _, ask := range ob.asks {

			if bid.amount.Equal(decimal.Zero) {
				return
			}
			if bid.price.GreaterThan(ask.price) {
				ob.addToBook(o)
				return
			}

			m := decimal.Min(bid.amount, ask.amount)

			pask := &PlacedOrder{
				orderID: bid.id,
				amount:  m,
				price:   ask.price,
			}
			pbid := &PlacedOrder{
				orderID: ask.id,
				amount:  m,
				price:   ask.price,
			}
			placed = append(placed, pask)
			placed = append(placed, pbid)

			bid.amount = bid.amount.Sub(m)
			ask.amount = ask.amount.Sub(m)

			if ask.amount.Equal(decimal.Zero) {
				ob.asks = ob.asks[1:]
				continue
			}

		}
	}

	return placed
}
