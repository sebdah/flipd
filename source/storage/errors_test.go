package storage

import "testing"

func TestErrors(t *testing.T) {
	tests := []struct {
		err             error
		expectedMessage string
	}{
		{
			err:             ErrDuplicate,
			expectedMessage: errDuplicate,
		},
		{
			err:             ErrNotFound,
			expectedMessage: errNotFound,
		},
	}

	for _, test := range tests {
		if test.err.Error() != test.expectedMessage {
			t.Errorf("expected %s, got %s", test.expectedMessage, test.err.Error())
		}
	}
}

func TestIsErrorCheckers(t *testing.T) {
	tests := []struct {
		err            error
		checker        func(error) bool
		expectedResult bool
	}{
		{
			err:            ErrDuplicate,
			checker:        IsErrDuplicate,
			expectedResult: true,
		},
		{
			err:            ErrNotFound,
			checker:        IsErrDuplicate,
			expectedResult: false,
		},
		{
			err:            ErrNotFound,
			checker:        IsErrNotFound,
			expectedResult: true,
		},
		{
			err:            ErrDuplicate,
			checker:        IsErrNotFound,
			expectedResult: false,
		},
	}

	for _, test := range tests {
		if result := test.checker(test.err); result != test.expectedResult {
			t.Errorf("expected %t, got %t", test.expectedResult, result)
		}
	}
}
