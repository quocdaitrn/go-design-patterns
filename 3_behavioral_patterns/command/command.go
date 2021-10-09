package command

// Order the command interface.
type Order interface {
	Execute()
}

// Stock the receiver.
type Stock struct {
	Name  string
	Count int
}

// Buy stock's action buy.
func (s *Stock) Buy() {
	s.Count += 1
}

// Sell stock's action sell.
func (s *Stock) Sell() {
	s.Count -= 1
}

type BuyStock struct {
	*Stock
}

func (b *BuyStock) Execute() {
	b.Stock.Buy()
}

type SellStock struct {
	*Stock
}

func (s *SellStock) Execute() {
	s.Stock.Sell()
}

// Broker the invoker.
type Broker struct {
	Orders []Order
}

func (b *Broker) TakeOrder(o Order) {
	b.Orders = append(b.Orders, o)
}

func (b *Broker) PlaceOrders() {
	for _, o := range b.Orders {
		o.Execute()
	}

	b.Orders = make([]Order, 0)
}


