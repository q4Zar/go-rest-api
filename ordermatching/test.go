package main

import (
	"testing"
	"time"
)

func TestOrderMatchingService(t *testing.T) {
	orderChannel := make(chan *Order, 100)
	matchedOrders := make(chan *Order, 100)

	go orderMatchingService(orderChannel, matchedOrders)

	// Test matching orders
	order1 := &Order{
		ID:        1,
		Side:      BUY,
		AssetPair: EURUSD,
		Amount:    1000.0,
		Price:     1.2,
		UserID:    "User1",
		Status:    Pending,
	}
	order2 := &Order{
		ID:        2,
		Side:      SELL,
		AssetPair: EURUSD,
		Amount:    1000.0,
		Price:     1.2,
		UserID:    "User2",
		Status:    Pending,
	}

	orderChannel <- order1
	orderChannel <- order2

	select {
	case matchedOrder := <-matchedOrders:
		if matchedOrder.ID != 1 && matchedOrder.ID != 2 {
			t.Errorf("Expected order ID 1 or 2, got %d", matchedOrder.ID)
		}
	case <-time.After(time.Second):
		t.Error("Expected matching orders, but timed out")
	}

	select {
	case matchedOrder := <-matchedOrders:
		if matchedOrder.ID != 1 && matchedOrder.ID != 2 {
			t.Errorf("Expected order ID 1 or 2, got %d", matchedOrder.ID)
		}
	case <-time.After(time.Second):
		t.Error("Expected matching orders, but timed out")
	}

	// Test no matching orders
	order3 := &Order{
		ID:        3,
		Side:      BUY,
		AssetPair: EURUSD,
		Amount:    1000.0,
		Price:     1.3,
		UserID:    "User3",
		Status:    Pending,
	}
	orderChannel <- order3

	select {
	case matchedOrder := <-matchedOrders:
		t.Errorf("Did not expect any matching orders, but got order ID %d", matchedOrder.ID)
	case <-time.After(time.Second):
		// Expected timeout
	}

	// Test multiple pending orders
	order4 := &Order{
		ID:        4,
		Side:      BUY,
		AssetPair: EURUSD,
		Amount:    500.0,
		Price:     1.2,
		UserID:    "User4",
		Status:    Pending,
	}
	order5 := &Order{
		ID:        5,
		Side:      SELL,
		AssetPair: EURUSD,
		Amount:    500.0,
		Price:     1.2,
		UserID:    "User5",
		Status:    Pending,
	}

	orderChannel <- order4
	orderChannel <- order5

	select {
	case matchedOrder := <-matchedOrders:
		if matchedOrder.ID != 4 && matchedOrder.ID != 5 {
			t.Errorf("Expected order ID 4 or 5, got %d", matchedOrder.ID)
		}
	case <-time.After(time.Second):
		t.Error("Expected matching orders, but timed out")
	}

	select {
	case matchedOrder := <-matchedOrders:
		if matchedOrder.ID != 4 && matchedOrder.ID != 5 {
			t.Errorf("Expected order ID 4 or 5, got %d", matchedOrder.ID)
		}
	case <-time.After(time.Second):
		t.Error("Expected matching orders, but timed out")
	}
}
