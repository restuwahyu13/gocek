package gocek

import "testing"

func TestToBeFalsy(action *testing.T) {
	action.Run("Should be ToBeFalsy - value i to be false", func(t *testing.T) {
		test := NewAssertion(t)
		test.Expect(false).ToBeFalsy()
	})

	action.Run("Should be ToBeFalsy - value is not to be false", func(t *testing.T) {
		test := NewAssertion(t)
		test.Expect(true).Not().ToBeFalsy()
	})
}
