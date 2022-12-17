package gocek

import "testing"

func TestToBeZero(action *testing.T) {
	action.Run("Should be ToBeZero - value zero is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(0).ToBeZero()
	})

	action.Run("Should be ToBeZero - value zero is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeZero()
	})
}
