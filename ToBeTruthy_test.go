package gocek

import "testing"

func TestToBeTruthy(action *testing.T) {
	action.Run("Should be ToBeTruthy - value is to be boolean", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(true).ToBeTruthy()
	})

	action.Run("Should be ToBeTruthy - value is not to be boolean", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(false).Not().ToBeTruthy()
	})
}
