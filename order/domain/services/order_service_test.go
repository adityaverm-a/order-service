package services

import (
	"demo/oms/order/data/models"
	"demo/oms/order/data/repositories"
	"demo/oms/order/domain/repositories/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetOrderByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockOrderRepository(ctrl)

	t.Run("return no error while fetching order, with ID: 1", func(t *testing.T) {
		orderID := int64(1)
		var want *models.Order

		createOrderForTesting(t, orderID, want)

		mockRepo.EXPECT().GetByID(orderID).Return(want, nil)

		testService := NewOrderService(mockRepo)

		got, err := testService.GetOrderByID(orderID)
		if err != nil {
			t.Fatal("Got error while fetching want by ID", err)
		}

		if got != want {
			t.Errorf("Expected no error Got : %v", got)
		}
	})

	t.Run("return error while fetching order, with ID: 1. As order with the given ID does not exists.", func(t *testing.T) {
		orderID := int64(1)

		mockRepo.EXPECT().GetByID(orderID).Return(nil, repositories.OrderNotFound)

		testService := NewOrderService(mockRepo)

		_, err := testService.GetOrderByID(orderID)

		assertError(t, err, repositories.OrderNotFound)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func createOrderForTesting(t testing.TB, orderID int64, order *models.Order) {
	t.Helper()

	price := float32(14.55)
	var orderItems []models.OrderItem

	orderItem := models.OrderItem{
		ID:          3,
		OrderID:     orderID,
		Description: "Test Product",
		Price:       price,
		Quantity:    1,
	}

	orderItems = append(orderItems, orderItem)

	orderModel := &models.Order{
		ID:           orderID,
		Status:       "COMPLETED",
		Total:        price,
		CurrencyUnit: "USD",
		OrderItem:    orderItems,
	}

	order = orderModel
}
