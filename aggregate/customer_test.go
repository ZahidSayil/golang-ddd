package aggregate_test

import (
	"errors"
	"testing"

	"github.com/ZahidSayil/Go-DDD/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "empty name validaiton",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "valid name",
			name:        "some-name",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected err %v , got %v", tc.expectedErr, err)
			}
		})
	}
}
