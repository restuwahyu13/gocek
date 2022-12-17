package gocek

import "testing"

func TestToBeLessThanOrEqual(action *testing.T) {
	action.Run("Should be ToBeLessThanOrEqual - value int lessThanEqual is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(10).ToBeLessThanOrEqual(10)
	})

	action.Run("Should be ToBeLessThanOrEqual - value int lessThanEqual is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeLessThanOrEqual(10)
	})
}
