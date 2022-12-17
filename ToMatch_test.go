package gocek

import (
	"testing"
)

func TestToMatch(action *testing.T) {

	action.Run("Should be ToMatch - value match is to be success", func(t *testing.T) {
		str := "hello wordl"

		test := NewGocek(t)
		test.Expect(str).ToMatch(`(wordl)+$`)
	})

	action.Run("Should be ToMatch - value match is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToMatch(`(jamal)+$`)
	})
}
