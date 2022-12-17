package gocek

import "testing"

func TestToBeMinus(action *testing.T) {
	action.Run("Should be ToBeMinus - value int min is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(-1).ToBeMinus()
	})

	action.Run("Should be ToBeMinus - value int min is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeMinus()
	})

	action.Run("Should be ToBeMinus - value float min is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(-2.45).ToBeMinus()
	})

	action.Run("Should be ToBeMinus - value float min is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeMinus()
	})
}
