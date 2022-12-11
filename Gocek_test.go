package gocek

import (
	"testing"
)

func TestNewGocek(action *testing.T) {
	action.Run("Shoud be TestNewGocek - expect return value success", func(t *testing.T) {
		test := NewGocek(t)
		res := test.Expect(2022)

		if res.v == nil {
			t.FailNow()
		}
	})

	action.Run("Shoud be TestNewGocek - expect return value nil", func(t *testing.T) {
		test := NewGocek(t)
		res := test.Expect(nil)

		if res.v != nil {
			t.FailNow()
		}
	})
}
