package gocek

import "testing"

func TestToBeLessThan(action *testing.T) {
	action.Run("Should be ToBeLessThan - value int is lessThan to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(5).ToBeLessThan(10)
	})

	action.Run("Should be ToBeLessThan - value int is lessThan not to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(20).Not().ToBeLessThan(10)
	})
}
