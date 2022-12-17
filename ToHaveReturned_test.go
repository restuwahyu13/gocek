package gocek

import "testing"

func TestToHaveReturned(action *testing.T) {
	action.Run("Should be ToHaveReturned - value returned response is to be success", func(t *testing.T) {
		sum := func() int {
			return 2 + 2
		}

		test := NewGocek(t)
		test.Expect(sum).ToHaveReturned()
	})

	action.Run("Should be ToHaveReturned - value returned response is to be failed", func(t *testing.T) {
		sum := func() {}

		test := NewGocek(t)
		test.Expect(sum).Not().ToHaveReturned()
	})
}
