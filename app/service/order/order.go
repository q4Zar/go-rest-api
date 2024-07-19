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

type AssetRepository interface {
	GetByUserIDAndType(ctx context.Context, userID uint, assetType string) (*model.Asset, error)
	Update(ctx context.Context, asset *model.Asset) (*model.Asset, error)
}

type Service struct {
	Session         session.Session
	Repository      Repository
	AssetRepository AssetRepository
	mu              sync.Mutex
	channels        map[string]OrderChannels
}

func NewService(session session.Session, repository Repository, assetRepository AssetRepository) *Service {
	service := &Service{
		Session:         session,
		Repository:      repository,
		AssetRepository: assetRepository,
		channels:        make(map[string]OrderChannels),
	}
	service.initChannels()
	return service
}

// initChannels initializes the buy and sell channels for order matching
func (s *Service) initChannels() {
	s.channels["EUR-USD"] = OrderChannels{
		Buy:  make(chan *model.Order, 100),
		Sell: make(chan *model.Order, 100),
	}
	s.channels["USD-EUR"] = OrderChannels{
		Buy:  make(chan *model.Order, 100),
		Sell: make(chan *model.Order, 100),
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
	// Determine asset types based on order side and asset pair
	// var assetTypeToCheck string
	// var requiredAmount float64
	// switch createDTO.AssetPair {
	// case "EUR-USD":
	// 	if createDTO.Side == "BUY" {
	// 		assetTypeToCheck = "EUR"
	// 		requiredAmount = createDTO.Amount
	// 	} else if createDTO.Side == "SELL" {
	// 		assetTypeToCheck = "USD"
	// 		requiredAmount = createDTO.Amount
	// 	}
	// case "USD-EUR":
	// 	if createDTO.Side == "BUY" {
	// 		assetTypeToCheck = "USD"
	// 		requiredAmount = createDTO.Amount
	// 	} else if createDTO.Side == "SELL" {
	// 		assetTypeToCheck = "EUR"
	// 		requiredAmount = createDTO.Amount
	// 	}
	// default:
	// 	return errors.New("unsupported asset pair")
	// }

	// // Fetch the user's balance for the determined asset type
	// balance, err := GetBalance(ctx, s.AssetRepository, createDTO.UserID, assetTypeToCheck)
	// if err != nil {
	// 	return errors.New(fmt.Sprintf("could not get balance: %v", err))
	// }

	// // Check if the user has enough balance
	// if createDTO.Side == "BUY" {
	// 	if requiredAmount < balance {
	// 		return errors.New("insufficient balance to complete the purchase")
	// 	}
	// } else if createDTO.Side == "SELL" {
	// 	if requiredAmount < balance {
	// 		return errors.New("insufficient balance to complete the sale")
	// 	}
	// }

	// Proceed to create the order
	order := typeutil.Copy(&model.Order{}, createDTO)
	order, err := s.Repository.Create(ctx, order)
	if err != nil {
		return errors.New(err)
	}

	// Add the order to the matching process
	s.addOrderToMatching(order)
	return nil
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
	s.mu.Lock()
	defer s.mu.Unlock()

	switch order.Side {
	case "BUY":
		s.channels[order.AssetPair].Buy <- order
	case "SELL":
		s.channels[order.AssetPair].Sell <- order
	}
}

func (s *Service) matchOrders(pair string) {
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
	// fmt.Println("processOrder")
	s.mu.Lock()
	defer s.mu.Unlock()

	oppositeOrders := s.channels[order.AssetPair]
	var oppositeChan chan *model.Order
	if oppositeSide == "BUY" {

		oppositeChan = oppositeOrders.Buy
	} else {
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

				// Update user balances
				fmt.Printf("Updating balances for users %d and %d\n", order.UserID, oppositeOrder.UserID)
				if err := UpdateBalance(context.Background(), s.AssetRepository, order.UserID, "EUR", -order.Amount); err != nil {
					fmt.Printf("Failed to update balance for user %d: %v\n", order.UserID, err)
					fmt.Println(8)
				}
				if err := UpdateBalance(context.Background(), s.AssetRepository, oppositeOrder.UserID, "USD", order.Amount); err != nil {
					fmt.Printf("Failed to update balance for user %d: %v\n", oppositeOrder.UserID, err)
					fmt.Println(9)
				}

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

func UpdateBalance(ctx context.Context, repo AssetRepository, userID uint, assetType string, amount float64) error {
	asset, err := repo.GetByUserIDAndType(ctx, userID, assetType)
	if err != nil {
		return fmt.Errorf("could not get asset: %w", err)
	}
	if asset == nil {
		return fmt.Errorf("asset not found for user %d and type %s", userID, assetType)
	}

	asset.Balance += amount
	if _, err := repo.Update(ctx, asset); err != nil {
		return fmt.Errorf("could not update asset: %w", err)
	}
	return nil
}

func GetBalance(ctx context.Context, repo AssetRepository, userID uint, assetType string) (float64, error) {
	asset, err := repo.GetByUserIDAndType(ctx, userID, assetType)

	if err != nil {
		return 0, fmt.Errorf("could not get asset: %w", err)
	}

	if asset == nil {
		return 0, fmt.Errorf("asset not found for user %d and type %s", userID, assetType)
	}

	return asset.Balance, nil
}
