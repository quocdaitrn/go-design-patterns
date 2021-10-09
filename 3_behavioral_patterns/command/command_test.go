package command

import "testing"

func TestCommand(t *testing.T) {
	stock := &Stock{Name: "VNM", Count: 100}
	buyOrder := &BuyStock{Stock: stock}
	sellOrder := &SellStock{Stock: stock}

	broker := &Broker{}
	broker.TakeOrder(buyOrder)
	broker.TakeOrder(buyOrder)
	broker.TakeOrder(sellOrder)

	broker.PlaceOrders()

	if stock.Count != 101 {
		t.Errorf("number of stock having code %s should be 101, but found %d", stock.Name, stock.Count)
	}
}