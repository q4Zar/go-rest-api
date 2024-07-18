package main

import (
	"fmt"
	"sync"
	"time"
)

type OrderSide string
type AssetPair string
type OrderStatus string

const (
	BUY  OrderSide = "BUY"
	SELL OrderSide = "SELL"

	EURUSD AssetPair = "EUR-USD"
	USDEUR AssetPair = "USD-EUR"

	Pending OrderStatus = "Pending"
	Filled  OrderStatus = "Filled"
)

type Order struct {
	ID        int
	Side      OrderSide
	AssetPair AssetPair
	Amount    float64
	Price     float64
	UserID    string
	Status    OrderStatus
}

var (
	orderID       = 0
	orderChannel  = make(chan *Order, 100)
	matchedOrders = make(chan *Order, 100)
	wg            sync.WaitGroup
)

func main() {
	wg.Add(1)
	go orderMatchingService(orderChannel, matchedOrders)

	// Simulate order creation
	createOrder(BUY, EURUSD, 1000.0, 1.2, "User1")
	createOrder(SELL, EURUSD, 1000.0, 1.2, "User2")

	// Simulate a delay to allow processing
	time.Sleep(1 * time.Second)

	// Close the orderChannel and wait for goroutines to finish
	close(orderChannel)
	wg.Wait()
}

func createOrder(side OrderSide, assetPair AssetPair, amount, price float64, userID string) {
	orderID++
	order := &Order{
		ID:        orderID,
		Side:      side,
		AssetPair: assetPair,
		Amount:    amount,
		Price:     price,
		UserID:    userID,
		Status:    Pending,
	}
	orderChannel <- order
}

func orderMatchingService(orderChannel, matchedOrders chan *Order) {
	defer wg.Done()

	pendingOrders := make(map[AssetPair][]*Order)

	for order := range orderChannel {
		matchOrder(order, pendingOrders, matchedOrders)
	}

	// After processing all orders, close the matchedOrders channel
	close(matchedOrders)
}

func matchOrder(order *Order, pendingOrders map[AssetPair][]*Order, matchedOrders chan *Order) {
	key := order.AssetPair

	if oppositeOrders, exists := pendingOrders[key]; exists {
		for i, oppositeOrder := range oppositeOrders {
			if oppositeOrder.Side != order.Side && oppositeOrder.Price == order.Price && oppositeOrder.Amount == order.Amount {
				fmt.Printf("Matching Orders: ID %d (side: %s) and ID %d (side: %s)\n", order.ID, order.Side, oppositeOrder.ID, oppositeOrder.Side)
				matchedOrders <- order
				matchedOrders <- oppositeOrder
				// Remove the matched order from pending orders
				pendingOrders[key] = append(oppositeOrders[:i], oppositeOrders[i+1:]...)
				return
			}
		}
	}
	// If no match found, add the order to pending orders
	pendingOrders[key] = append(pendingOrders[key], order)
}
