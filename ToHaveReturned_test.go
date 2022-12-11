package gocek

import "testing"

func TestToHaveReturned(action *testing.T) {
	action.Run("Should be TestToHaveReturned - value is returned success", func(t *testing.T) {
		sum := func() int {
			return 2 + 2
		}

		test := NewGocek(t)
		test.Expect(sum).ToHaveReturned()
	})

	action.Run("Should be TestToHaveReturned - value is returned failed", func(t *testing.T) {
		sum := func() {}

		test := NewGocek(t)
		test.Expect(sum).Not().ToHaveReturned()
	})
}
