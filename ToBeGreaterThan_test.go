package gocek

import "testing"

func TestToBeGreaterThan(action *testing.T) {
	action.Run("Should be ToBeGreaterThan - value int is greaterThan to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(10).ToBeGreaterThan(5)
	})

	action.Run("Should be ToBeGreaterThan - value int is greaterThan not to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(10).Not().ToBeGreaterThan(20)
	})
}
