package memory

import (
	"errors"
	"testing"

	"github.com/ZahidSayil/Go-DDD/aggregate"
	"github.com/ZahidSayil/Go-DDD/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	cust, err := aggregate.NewCustomer("zahid")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{id: cust},
	}

	testCases := []testCase{
		{
			name:        "Not customer found",
			id:          uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "Customer by Id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v , got %v ", tc.expectedErr, err)
			}
		})
	}

}
