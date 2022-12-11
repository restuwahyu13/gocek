package gocek

import "testing"

func TestToBeGreaterThanOrEqual(action *testing.T) {
	action.Run("Should be ToBeGreaterThanOrEqual - value int is greaterEqual not tobe success", func(t *testing.T) {
		test := NewAssertion(t)
		test.Expect(10).ToBeGreaterThanOrEqual(10)
	})

	action.Run("Should be ToBeGreaterThanOrEqual - value int is greaterEqual not tobe failed", func(t *testing.T) {
		test := NewAssertion(t)
		test.Expect(10).Not().ToBeGreaterThanOrEqual(20)
	})
}
