package order

import (
	"context"
	"fmt"
	"sync"

	"github.com/q4Zar/go-rest-api/database/model"
	"github.com/q4Zar/go-rest-api/dto"
	"github.com/q4Zar/go-rest-api/service"
	"goyave.dev/filter"
	"goyave.dev/goyave/v5/database"
	"goyave.dev/goyave/v5/util/errors"
	"goyave.dev/goyave/v5/util/session"
	"goyave.dev/goyave/v5/util/typeutil"
)

// func init() {
// }

// OrderChannels holds the channels for buy and sell orders
type OrderChannels struct {
	Buy  chan *model.Order
	Sell chan *model.Order
}

type Repository interface {
	Index(ctx context.Context, request *filter.Request) (*database.Paginator[*model.Order], error)
	GetByID(ctx context.Context, id uint) (*model.Order, error)
	Create(ctx context.Context, order *model.Order) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) (*model.Order, error)
	Delete(ctx context.Context, id uint) error
	IsOwner(ctx context.Context, resourceID, ownerID uint) (bool, error)
}

type Service struct {
	Session    session.Session
	Repository Repository
	mu         sync.Mutex
	channels   map[string]OrderChannels
}

func NewService(session session.Session, repository Repository) *Service {
	service := &Service{
		Session:    session,
		Repository: repository,
		channels:   make(map[string]OrderChannels),
	}
	service.initChannels()
	return service
}

// initChannels initializes the buy and sell channels for order matching
func (s *Service) initChannels() {
	fmt.Println("initChannels")
	s.channels["EUR-USD"] = OrderChannels{
		Buy:  make(chan *model.Order),
		Sell: make(chan *model.Order),
	}
	s.channels["USD-EUR"] = OrderChannels{
		Buy:  make(chan *model.Order),
		Sell: make(chan *model.Order),
	}
	go s.matchOrders("EUR-USD")
	go s.matchOrders("USD-EUR")
}

func (s *Service) Index(ctx context.Context, request *filter.Request) (*database.PaginatorDTO[*dto.Order], error) {
	paginator, err := s.Repository.Index(ctx, request)
	if err != nil {
		return nil, errors.New(err)
	}
	return typeutil.MustConvert[*database.PaginatorDTO[*dto.Order]](paginator), nil
}

func (s *Service) Create(ctx context.Context, createDTO *dto.CreateOrder) error {
	order := typeutil.Copy(&model.Order{}, createDTO)
	order, err := s.Repository.Create(ctx, order)
	// add order into channelSide : Buy|Sell
	s.addOrderToMatching(order)
	return errors.New(err)
}

func (s *Service) Delete(ctx context.Context, id uint) error {
	return s.Repository.Delete(ctx, id)
}

func (s *Service) IsOwner(ctx context.Context, resourceID, ownerID uint) (bool, error) {
	return s.Repository.IsOwner(ctx, resourceID, ownerID)
}

func (s *Service) Name() string {
	return service.Order
}

func (s *Service) addOrderToMatching(order *model.Order) {
	fmt.Println("addOrderToMatching")
	s.mu.Lock()
	defer s.mu.Unlock()

	switch order.Side {
	case "BUY":
		fmt.Println("BUY", order)
		s.channels[order.AssetPair].Buy <- order
	case "SELL":
		fmt.Println("SELL", order)
		s.channels[order.AssetPair].Sell <- order
	}
	fmt.Println("No Case :/")
}

func (s *Service) matchOrders(pair string) {
	fmt.Println("matchOrders")
	for {
		select {
		case buyOrder := <-s.channels[pair].Buy:
			s.processOrder(buyOrder, "SELL")
		case sellOrder := <-s.channels[pair].Sell:
			s.processOrder(sellOrder, "BUY")
		}
	}
}

func (s *Service) processOrder(order *model.Order, oppositeSide string) {
	fmt.Println("processOrder")
	s.mu.Lock()
	defer s.mu.Unlock()

	oppositeOrders := s.channels[order.AssetPair]
	var oppositeChan chan *model.Order
	if oppositeSide == "BUY" {
		fmt.Println("oppositeSide", "BUY")
		oppositeChan = oppositeOrders.Buy
	} else {
		fmt.Println("oppositeSide", "SELL")
		oppositeChan = oppositeOrders.Sell
	}

	for {
		select {
		case oppositeOrder := <-oppositeChan:
			if oppositeOrder.Price == order.Price && oppositeOrder.Amount == order.Amount {
				fmt.Printf("Matching Orders: ID %d (side: %s) and ID %d (side: %s)\n", order.ID, order.Side, oppositeOrder.ID, oppositeOrder.Side)

				// Update orders in the database
				order.Status = "Filled"
				oppositeOrder.Status = "Filled"
				if _, err := s.Repository.Update(context.Background(), order); err != nil {
					fmt.Printf("Failed to update order status: %v\n", err)
				}
				if _, err := s.Repository.Update(context.Background(), oppositeOrder); err != nil {
					fmt.Printf("Failed to update order status: %v\n", err)
				}

				// Simulate balance update (you will need to implement this)
				fmt.Printf("Updating balances for users %d and %d\n", order.UserID, oppositeOrder.UserID)

				return
			} else {
				// Put the order back to the channel if not matched
				oppositeChan <- oppositeOrder
			}
		default:
			// No matching order found, put the order back to the channel
			if order.Side == "BUY" {
				s.channels[order.AssetPair].Buy <- order
			} else {
				s.channels[order.AssetPair].Sell <- order
			}
			return
		}
	}
}
