package memory

import (
	"errors"
	"fvm/impl-ddd/aggregate"
	"fvm/impl-ddd/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func TestCustomer_GetCustomer(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := memoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			test:        "no customer by id",
			id:          uuid.MustParse("b245b6de-aa56-11ed-afa1-0242ac120002"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			test:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
