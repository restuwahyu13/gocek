package gocek

import (
	"testing"
)

func TestToBe(action *testing.T) {
	action.Run("Should be ToBe - value is to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(12).ToBe(12)
	})

	action.Run("Should be ToBe - value not to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(12).Not().ToBe(13)
	})

	action.Run("Should be ToBe - value is to be null", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).ToBe(nil)
	})

	action.Run("Should be ToBe - value not to be null", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBe(12)
	})
}
