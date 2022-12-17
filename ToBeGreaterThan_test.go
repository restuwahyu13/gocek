package gocek

import "testing"

func TestToBeGreaterThan(action *testing.T) {
	action.Run("Should be ToBeGreaterThan - value int greaterThan is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(10).ToBeGreaterThan(5)
	})

	action.Run("Should be ToBeGreaterThan - value int greaterThan is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeGreaterThan(20)
	})
}
