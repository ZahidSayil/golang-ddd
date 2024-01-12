package customer

import (
	"errors"

	"github.com/ZahidSayil/Go-DDD/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer Not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("failed to update customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
