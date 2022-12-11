package gocek

import (
	"testing"
)

func TestNewAssertion(action *testing.T) {
	action.Run("Shoud be TestNewAssertion - expect return value success", func(t *testing.T) {
		test := NewAssertion(t)
		res := test.Expect(2022)

		if res.v == nil {
			t.FailNow()
		}
	})

	action.Run("Shoud be TestNewAssertion - expect return value nil", func(t *testing.T) {
		test := NewAssertion(t)
		res := test.Expect(nil)

		if res.v != nil {
			t.FailNow()
		}
	})
}
