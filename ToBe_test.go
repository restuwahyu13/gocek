package gocek

import (
	"testing"
)

func TestToBe(action *testing.T) {
	action.Run("Should be ToBe - value int is to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect("gocek").ToBe("gocek")
	})

	action.Run("Should be ToBe - value int not to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect("gocek").Not().ToBe(13)
	})

	action.Run("Should be ToBe - value int is to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(12).ToBe(12)
	})

	action.Run("Should be ToBe - value int not to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(12).Not().ToBe(13)
	})

	action.Run("Should be ToBe - value float is to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(2.5).ToBe(2.5)
	})

	action.Run("Should be ToBe - value float not to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(2.5).Not().ToBe(13)
	})

	action.Run("Should be ToBe - value boolean is to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(true).ToBe(true)
	})

	action.Run("Should be ToBe - value boolean not to be match", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(true).Not().ToBe(13)
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
