// createOrder creates an order based on the given order type and parameters.
// It uses the appropriate order factory to create the order.
// The orderType parameter specifies the type of order to create.
// The paymentMethod parameter specifies the payment method for the order.
// The currency parameter specifies the currency for the order.
// The address parameter specifies the shipping address for the order.
// The name parameter specifies the name of the customer placing the order.
// The email parameter specifies the email of the customer placing the order.
// The items parameter specifies the list of items in the order.
// It returns the created order and any error encountered during the creation process.
package order

func createOrder(orderType string, paymentMethod, currency, address, name, email string, items []OrderItem) (Order, error) {
	factory, err := getOrderFactory(orderType)
	if err != nil {
		return Order{}, err
	}

	order, err := factory.CreateOrder(items, paymentMethod, currency, address, name, email)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}
