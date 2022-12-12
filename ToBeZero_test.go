package gocek

import "testing"

func TestToBeZero(action *testing.T) {
	action.Run("Should be ToBeZero - value is to be zero", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(0).ToBeZero()
	})

	action.Run("Should be ToBeZero - value is not to be zero", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeZero()
	})
}
