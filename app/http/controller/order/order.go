package controllers

import (
	"sync"

	"goyave.dev/goyave/v5"
)

type Order struct {
	ID        int     `json:"id"`
	UserID    int     `json:"user_id"`
	Side      string  `json:"side"`
	AssetPair string  `json:"asset_pair"`
	Amount    float64 `json:"amount"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
}

type User struct {
	ID      int
	Balance map[string]float64
}

var (
	orders         = []Order{}
	pendingOrders  = []Order{}
	ordersMutex    sync.Mutex
	users          = make(map[int]*User)
	orderIDCounter = 1
	orderChan      = make(chan Order, 100)
	wg             sync.WaitGroup
)

func init() {
	// Start the order matching service
	wg.Add(1)
	go matchOrders()
}

// CreateOrder handles creating a new order
func CreateOrder(response *goyave.Response, request *goyave.Request) {
	order := Order{
		UserID:    request.UserID(),
		Side:      request.String("side"),
		AssetPair: request.String("asset_pair"),
		Amount:    request.Float("amount"),
		Price:     request.Float("price"),
		Status:    "pending",
	}

	ordersMutex.Lock()
	order.ID = orderIDCounter
	orderIDCounter++
	orders = append(orders, order)
	pendingOrders = append(pendingOrders, order)
	ordersMutex.Unlock()

	// Send the order to the matching service
	orderChan <- order

	response.JSON(201, order)
}

// GetOrders handles fetching all orders
func GetOrders(response *goyave.Response, request *goyave.Request) {
	ordersMutex.Lock()
	defer ordersMutex.Unlock()
	response.JSON(200, orders)
}

// GetAssets handles fetching user assets
func GetAssets(response *goyave.Response, request *goyave.Request) {
	userID := request.UserID() // Assuming user ID is retrieved from the request context
	user, exists := users[userID]
	if !exists {
		response.Error(404, "User not found")
		return
	}
	response.JSON(200, user.Balance)
}

func matchOrders() {
	defer wg.Done()
	for order := range orderChan {
		ordersMutex.Lock()
		for i, pendingOrder := range pendingOrders {
			if pendingOrder.AssetPair == order.AssetPair && pendingOrder.Price == order.Price && pendingOrder.Amount == order.Amount && pendingOrder.Side != order.Side {
				// Match found
				pendingOrders = append(pendingOrders[:i], pendingOrders[i+1:]...)
				updateOrderStatus(order.ID, "filled")
				updateOrderStatus(pendingOrder.ID, "filled")
				updateBalances(order, pendingOrder)
				break
			}
		}
		ordersMutex.Unlock()
	}
}

func updateOrderStatus(orderID int, status string) {
	for i, order := range orders {
		if order.ID == orderID {
			orders[i].Status = status
			break
		}
	}
}

func updateBalances(order1, order2 Order) {
	user1 := getUser(order1.UserID)
	user2 := getUser(order2.UserID)

	if order1.Side == "BUY" {
		user1.Balance["EUR"] += order1.Amount / order1.Price
		user1.Balance["USD"] -= order1.Amount
		user2.Balance["EUR"] -= order2.Amount / order2.Price
		user2.Balance["USD"] += order2.Amount
	} else {
		user1.Balance["EUR"] -= order1.Amount / order1.Price
		user1.Balance["USD"] += order1.Amount
		user2.Balance["EUR"] += order2.Amount / order2.Price
		user2.Balance["USD"] -= order2.Amount
	}
}

func getUser(userID int) *User {
	user, exists := users[userID]
	if !exists {
		user = &User{
			ID:      userID,
			Balance: make(map[string]float64),
		}
		users[userID] = user
	}
	return user
}
