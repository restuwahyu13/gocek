package gocek

import "testing"

func TestToBeMinus(action *testing.T) {
	action.Run("Should be ToBeMinus - value int is to be minus", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(-1).ToBeMinus()
	})

	action.Run("Should be ToBeMinus - value float is to be minus", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(-2.45).ToBeMinus()
	})

	action.Run("Should be ToBeMinus - value int is not to be minus", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(1).Not().ToBeMinus()
	})

	action.Run("Should be ToBeMinus - value float is not to be minus", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(2.45).Not().ToBeMinus()
	})
}
