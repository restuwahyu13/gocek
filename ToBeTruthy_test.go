package gocek

import "testing"

func TestToBeTruthy(action *testing.T) {
	action.Run("Should be ToBeTruthy - value true is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(true).ToBeTruthy()
	})

	action.Run("Should be ToBeTruthy - value true is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeTruthy()
	})
}
