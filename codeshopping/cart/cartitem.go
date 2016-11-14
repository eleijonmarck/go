package cart

// CartItem is a item inside of a cart
type CartItem struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Price      float64                `json:"price"`
	Quantity   int                    `json:"quantity"`
	Attributes map[string]interface{} `json:"attrs"`
}
