package entity

import "errors"

type Order struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) { //*Order vai modificar nessa instancia
	order := &Order{ //endereco de memoria
		Id:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.Validate()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) Validate() error {
	if o.Id == "" {
		return errors.New("Id is required.")
	}
	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if o.Tax <= 0 {
		return errors.New("Invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil
}