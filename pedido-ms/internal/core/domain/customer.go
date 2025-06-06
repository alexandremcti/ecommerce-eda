package domain

type (
	Address struct {
		street, number, postalCode, city, state string
	}

	Customer struct {
		id                         string
		firstName, lastName, email string
		deliveryAddress            Address
	}

	CustomerParams struct {
		Id                                                                        string
		FirstName, LastName, Email, Street, Streetnumber, PostalCode, City, State string
	}
)

func CreateCustomer(input CustomerParams) *Customer {
	customer := Customer{
		firstName: input.FirstName,
		lastName:  input.LastName,
		email:     input.Email,
		deliveryAddress: Address{
			street:     input.Street,
			number:     input.Streetnumber,
			postalCode: input.PostalCode,
			city:       input.City,
			state:      input.State,
		},
	}

	return &customer
}

func RecoverCustomer(input CustomerParams) *Customer {
	customer := Customer{
		id:        input.Id,
		firstName: input.FirstName,
		lastName:  input.LastName,
		email:     input.Email,
		deliveryAddress: Address{
			street:     input.Street,
			number:     input.Streetnumber,
			postalCode: input.PostalCode,
			city:       input.City,
			state:      input.State,
		},
	}

	return &customer
}

func (c *Customer) Map() *CustomerParams {
	cp := CustomerParams{}
	cp.Id = c.id
	cp.FirstName = c.firstName
	cp.LastName = c.lastName
	cp.Email = c.email
	cp.Street = c.deliveryAddress.street
	cp.Streetnumber = c.deliveryAddress.number
	cp.PostalCode = c.deliveryAddress.postalCode
	cp.City = c.deliveryAddress.city
	cp.State = c.deliveryAddress.state

	return &cp
}
