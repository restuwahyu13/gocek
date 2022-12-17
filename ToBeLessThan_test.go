package gocek

import "testing"

func TestToBeLessThan(action *testing.T) {
	action.Run("Should be ToBeLessThan - value int lessThan is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(5).ToBeLessThan(10)
	})

	action.Run("Should be ToBeLessThan - value int lessThan is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeLessThan(10)
	})
}
