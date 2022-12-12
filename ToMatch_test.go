package gocek

import "testing"

func TestToMatch(action *testing.T) {

	action.Run("Should be ToMatch - value is to be match", func(t *testing.T) {
		str := "hello wordl"

		test := NewGocek(t)
		test.Expect(str).ToMatch(`(wordl)+$`)
	})

	action.Run("Should be ToMatch - value is not to be match", func(t *testing.T) {
		str := "hello wordl"

		test := NewGocek(t)
		test.Expect(str).Not().ToMatch(`(jamal)+$`)
	})
}
