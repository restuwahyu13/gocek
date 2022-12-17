package gocek

import "testing"

func TestToHaveReturnedTimes(action *testing.T) {
	action.Run("Should be ToHaveReturnedTimes - value is call times is to be success", func(t *testing.T) {
		sum := func() int {
			return 2 + 2
		}

		test := NewGocek(t)
		test.Expect([]interface{}{sum, sum}).ToHaveReturnedTimes(2)
	})

	action.Run("Should be ToHaveReturnedTimes - value is call times is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToHaveReturnedTimes(2)
	})
}
